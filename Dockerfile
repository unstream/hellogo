FROM node:16-alpine3.11 as build-node
#RUN apk --no-cache --virtual build-dependencies add \
#        python \
#        make \
#        g++

WORKDIR /workdir
COPY web/ .
RUN npm install
RUN npm run build

FROM golang:1.17-alpine as build-go
WORKDIR /app

ADD go.mod go.sum ./
RUN go mod download
ADD . .
COPY --from=build-node /workdir/build ./web/build

ARG GOOS=linux
ARG GOARCH=amd64
ARG CGO_ENABLED=0
RUN go build -o /main main.go

FROM alpine:3.15
COPY --from=build-go /main /main
ENTRYPOINT [ "/main" ]