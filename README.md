# Hello Go

Service to create mandelbrot images implemented in Go.

## Usage

```
docker build . -t hellogo
docker run -it -p 8080:8080 hellogo
```

##APIs

### Health API
http://localhost:8080/health

### Mandelbrot image API
http://localhost:8080/mandelbrot