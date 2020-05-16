package main

// Credit to @perimosocordiae for reverse engineering
// http://perimosocordiae.github.io/articles/pyhrm.html

/* Test Cases
eJxjYWBg6M7k7QNSDLuqzyYwjIJRMApGFAAAwIMD9g;
*/

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	
	"github.com/fogleman/gg"
)

type Coords [][2]uint16

/* Scales a coordinate into the range [1-65536]. */
func scale(n, max int) uint16 {
	return uint16((float64(n) / float64(max)) * 65535 + 1)
}

/* Encodes a PNG image into a HRM comment. */
func encodePNG(path string) (string, error) {
	// to-do: take
	img, err := gg.LoadPNG(path)
	if err != nil {
		return "", err
	}
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	coords := make([][2]uint16, 0)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			point := img.At(x, y)
			r, g, b, a := point.RGBA()
			if r == 0 && g == 0 && b == 0 && a == 65535 {
				cx := scale(x, width)
				cy := scale(y, height)
				coords = append(coords, [2]uint16{cx, cy})
			}
		}
	}
	return encodeComment(coords)
}

/* Encodes ASCII text into a HRM comment. */
func encodeText(text string) (string, error) {
	// to-do, use gg.DrawText or variant
	return "to-do", nil
}

/* Encodes a sequence of coordinates into HRM comment format. */
func encodeComment(coords Coords) (string, error) {
	// to-do
	fmt.Println(coords)
	return "todo", nil
}

/* Decodes the base64, zlib-compressed comment into its coordinate representation. */
func decodeComment(b64String string) (Coords, error) {
	zlibData, err := base64.RawStdEncoding.DecodeString(b64String[:len(b64String) - 1])
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(zlibData)
	z, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(z)
	return decodeCoords(data), nil
}

/* Decodes the binary data of a comment into a slice of (x, y) tuples. */
func decodeCoords(data []byte) Coords {
	header := binary.LittleEndian.Uint32(data[:4])
	data = data[4:4 * header]
	coords := make([][2]uint16, header)
	for i := 0; i < int(4 * header); i += 4 {
		x := binary.LittleEndian.Uint16(data[i:i + 2])
		y := binary.LittleEndian.Uint16(data[i + 2:i + 4])
		coords[i / 4][0] = x
		coords[i / 4][1] = y
	}
	return coords
}

/* Displays coordinates onto an image. */
func displayCoords(coords Coords) {
	const width = 65536 / 100
	const height = 65536 / 100
	ctx := gg.NewContext(width, height)
	ctx.SetLineWidth(10)
	segments := make([]Coords, 0)
	var segment Coords = nil
	for i := 0; i < len(coords); i += 1 {
		if coords[i] == [2]uint16{0, 0} {
			if len(segment) > 0 {
				segments = append(segments, segment)
				segment = nil
			}
		} else {
			segment = append(segment, coords[i])
		}
	}
	for i := 0; i < len(segments); i += 1 {
		if len(segments[i]) == 1 {
			x := float64(segments[i][0][0] / 100)
			y := float64(segments[i][0][1] / 100)
			ctx.DrawPoint(x, y, 1)
		} else {
			prevX := 0.0
			prevY := 0.0
			for _, point := range segments[i] {
				x := float64(point[0] / 100)
				y := float64(point[1] / 100)
				if !(prevX == 0 && prevY == 0) {
					ctx.DrawLine(prevX, prevY, x, y)
					ctx.Stroke()
				}
				prevX, prevY = x, y
			}
		}
	}
	ctx.SavePNG("out.png")
}

func main() {
	// Parse input as text or file path
	var decode string
	var encode string
	flag.StringVar(&decode, "decode", "", "Path to input to decode.")
	flag.StringVar(&encode, "encode", "", "Path to PNG image or text to encode.")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Usage: comments [text] [--decode path] [--encode path | text")
		os.Exit(1)
	}
	if encode != "" {
		var encoded string
		var err error
		if strings.HasSuffix(encode, ".png") {
			encoded, err = encodePNG(encode)
		} else {
			encoded, err = encodeText(encode)
		}
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(encoded)
		os.Exit(0)
	}
	var b64String string
	if decode != "" {
		bytes, err := ioutil.ReadFile(decode)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		b64String = string(bytes)
	} else {
		b64String = flag.Arg(0)
	}
	coords, err := decodeComment(b64String)
	if err != nil {
		panic(err)
	}
	displayCoords(coords)
}
