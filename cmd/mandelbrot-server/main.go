package main

import (
	"bytes"
	"github.com/unstream/hellogo/mandelbrot"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/mandelbrot", MandelbrotImage)
	log.Print("Server started at port 8080.")
	err := http.ListenAndServe("0.0.0.0:8080", handler)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func MandelbrotImage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var img image.Image
	img = mandelbrot.Image()
	writeImage(w, &img)
	log.Print("writing image took ", time.Since(start))
}

// writeImage encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
