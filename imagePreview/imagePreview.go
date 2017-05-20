package main

import (
	"fmt"
	"github.com/NYU-Efficient-Room-Traversal/Rangefinder"
	"github.com/NYU-Efficient-Room-Traversal/Tools/cameraStreamer"
	"image"
	"image/color"
	"image/png"
	"net/http"
)

var ch chan image.Image

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Image Requested")

	img := <-ch
	rgba := img.(*image.RGBA)
	mat := rangefinder.NewMonoImageMatrix(rgba, 100)

	res.Header().Set("Content-Type", "image/png")
	err := png.Encode(res, matToImage(mat))
	if err != nil {
		fmt.Println(err)
	}
}

func matToImage(mat *rangefinder.MonoImageMatrix) image.Image {
	img := image.NewGray(image.Rectangle{Max: image.Point{X: mat.Width, Y: mat.Height}})
	for x := 0; x < mat.Width; x++ {
		for y := 0; y < mat.Height; y++ {
			if mat.Image[x][y] {
				img.SetGray(x, y, color.Gray{Y: 255})
			} else {
				img.SetGray(x, y, color.Gray{Y: 0})
			}
		}
	}
	return img
}

func main() {
	ch = make(chan image.Image, 10)
	go cameraStreamer.Open(ch)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8086", nil)
	fmt.Println("vim-go")
}
