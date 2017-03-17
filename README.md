# Tools
Go Tools to Help Us Out

## Installing

`go get github.com/NYU-Efficient-Room-Traversal/Tools/...`

## Arrayifier Library
### Functions
`func Arrayify(filepath string) [][]Pixel`

Converts a JPG or PNG image to a two dimensional array of `Pixels`, which
represent hues of a pixel from the HSV colorspace. For more info on `image` and the
HSV colorspace, look [here](https://golang.org/pkg/image/).

Each value `[x][y]` will hold a `float64` array of size three, **hue**, **saturation**, and **value**

Example: `[0][0][0]` will return the hue of the upper left pixel.

### Other Useful Tools

**Finding a Raspberry Pi on the network**

`# arp -na | grep -i b8:27:eb`
