//
// Go Tools for Laser Range Finder
// arrayifier.go
//
// Cole Smith - css@nyu.edu
// Eric Lin   - eric.lin@nyu.edu
// LICENSE: GPL 3.0
//

package arrayifier

// import image // You will likely need this. Mainly image/color

// A pixel for an image defined in the
// HSV colorspace
type Pixel struct {
	hue float64
	sat float64
	val float64
}

// Creates a reference to a new Pixel struct
// using the provided HSV values
func NewPixel(hue, sat, val float64) *Pixel {
	return &Pixel{hue, sat, val}
}

// TODO
// Reads a JPG or PNG file and returns the two dimensional
// array of pixels of that image in HSV colorspace
func Arrayify(filePath string) [][]Pixel {
	// ...
	return make([][]Pixel, 10) // Placeholder
}
