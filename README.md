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
http://localhost:8080/mandelbrot?c0=-0.5%2B0.25i&c1=0.5%2B1i&max_integrations=1000&w=2000&h=1000