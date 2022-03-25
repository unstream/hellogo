FROM node:16-alpine3.11 as build-node
#RUN apk --no-cache --virtual build-dependencies add \
#        python \
#        make \
#        g++

WORKDIR /workdir
COPY web/ .
#RUN npm install
#RUN npm run build

FROM golang:1.17rc2-alpine3.14 as build-go

ENV GOPATH ""
RUN go env -w GOPROXY=direct
RUN apk add git

ADD go.mod go.sum ./
#RUN go mod download
ADD . .
COPY --from=build-node /workdir/build ./web/build
RUN go build -o /main main.go

FROM alpine:3.13
COPY --from=build-go /main /main
ENTRYPOINT [ "/main" ]