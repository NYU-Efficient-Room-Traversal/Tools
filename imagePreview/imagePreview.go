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

const VAL_THRESHOLD = 0.00
const HUE_THRESHOLD = 20.0
const HUE_TARGET = 332.0

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Image Requested")

	img := <-ch

	rgba, ok := img.(*image.RGBA)
	if !ok {
		fmt.Println("Type Assertion for RGBA image failed")
		return
	}

	// Convert to Matrix
	imageMatrix := rangefinder.NewImageMatrix(rgba)

	// Filter by hue
	hueFiltered := imageMatrix.ConvertToMonoImageMatrixFromHue(HUE_TARGET, HUE_THRESHOLD)

	// Filter by value
	valFiltered := imageMatrix.ConvertToMonoImageMatrixFromValue(VAL_THRESHOLD)

	// Intersect
	intersectFiltered, err := rangefinder.GetMonoIntersectMatrix(hueFiltered, valFiltered)
	if err != nil {
		fmt.Println(err)
	}

	//mat := rangefinder.NewMonoImageMatrix(rgba, THRESHOLD)

	res.Header().Set("Content-Type", "image/png")
	err = png.Encode(res, matToImage(intersectFiltered))
	//err := png.Encode(res, img)
	if err != nil {
		fmt.Println("PNG ENCODE ERROR: %v", err)
		return
	}
}

func matToImage(mat *rangefinder.MonoImageMatrix) image.Image {
	img := image.NewGray(image.Rectangle{Max: image.Point{X: mat.Height, Y: mat.Width}})
	for x := 0; x < mat.Height; x++ {
		for y := 0; y < mat.Width; y++ {
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
	ch = make(chan image.Image)
	go cameraStreamer.Open(ch)

	fmt.Println("Server Starting at: localhost:8086")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8086", nil)
}
