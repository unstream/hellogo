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

http://localhost:8080/api/v1/mandelbrot
http://localhost:8080/api/v1/mandelbrot?c0=-0.8665014418863999&c0i=-0.24407616104639998&c1=-0.8640678719007999&c1i=-0.24164259106080002&max_iterations=10000&w=1000&h=1000
http://localhost:8080/api/v1/mandelbrot?c0=-0.5307152255936957&c0i=-0.6706386289814761&c1=-0.5307152253486942&c1i=-0.6706386287364745&max_iterations=1000000&w=2000&h=2000