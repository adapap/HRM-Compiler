package main

// Credit to @perimosocordiae for reverse engineering
// http://perimosocordiae.github.io/articles/pyhrm.html

/* Test Cases

hello, world:
eJwrY2Bg+CH4XDdcSK9fVPjgHFHhxM1AIQZR4XvNp8V2ZlTKTk8vVJDNUVf5XK2usnMtSC7NgLcvzvh1
T4d1wISlNiu6WGzN6gwtmvL3GpemqRiWpuXpJuRN1hatNdGePMlTX3CFo9GtddUmshtPWcluBOk/78Kh
v8pNMmiVW+YiEH+Sf5PW+oAq69zAkFUgvnxSSOqJuJ0Z+jHeZReif0w9Eee+8Gpq1+K8jK7FrzLdFz7K
uTblf+6lKvbspvwdKSGps5NDUkH6kisyF7mV22/yK7m0H8Q37tqZsb3DuX17h+R0467Hc4273Beu7z63
9EyP4IrUfqOVTRPeL/s3kWX+nklbJgtPe7+Me6bgCqY5agvk526olJ/7J4tpzvR0kDm3Vv3Jclzxufro
stc9R5c9nquyPHOR44pzS0XWlq72W1+6OnuD0UqWze4LG7bEzFTcMrvx+cbnRX7r7XNBer33lKZl7LXP
Pb73c3Xrvoh65/11s1r35ZY779cqXbQ/IW/WodK0M8ch7n93abH2u0s7Mwwu6fX3X3q/DCT2/dnzIsfH
f7I4Ht7KfH3/T5bAvQ2VG+/q9effk5zO8XD+7AWP588Oez55UvGLk92XXmTGfn+2wtXy2VEdhlEwCogE
AKxj1UM;

shaded corner
eJyzYmBgaGC24mtgZgCDKWByIwuEFy9iwbaB7RkXhPdQGES+0oDw9gl7s0JY7Nwg8pQShGeuvoAPwvKW
BKvzgvDiwp8qQ1g6QLpfdDpQfVAsRIQ33V76mqKOFoTHqsnAcFAqAqhqaj6I7wZ223lOiOxZ6VbGs9JH
GC/JHWE0s2Rg2GHFAPaDWjQDQ2omRM3byg4G77IOhtQoEC/ymf1JhlEwCkYBVgAASVEj4g;

*/

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"
	"os"
	"strings"
	
	"github.com/fogleman/gg"
)

type Coords []Point
type Point [2]float64
var EMPTY_POINT = [2]float64{0, 0}
const HRM_MAX = 65535
const IMG_WIDTH = 420
const IMG_HEIGHT = 140

/* Scales a coordinate into from one range into another. */
func scale(a, maxA, maxB int) float64 {
	return (float64(a) / float64(maxA)) * float64(maxB) + 1
}

/* Check if a pixel is black at a certain point */
func pixelIsBlack(img image.Image, x, y int) bool {
	point := img.At(x, y)
	r, g, b, a := point.RGBA()
	return r == 0 && g == 0 && b == 0 && a == HRM_MAX
}

/* Encodes a PNG image into a HRM comment. */
func encodePNG(path string, checkNeighbors bool) (string, error) {
	return "", fmt.Errorf("PNG encoding not supported.")
	img, err := gg.LoadPNG(path)
	if err != nil {
		return "", err
	}
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	coords := make(Coords, 0)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			if pixelIsBlack(img, x, y) {
				// Check if there exists a pixel a distance of radius 
				// in all four cardinal directions (must be a source point)
				if checkNeighbors {
					n := pixelIsBlack(img, x, int(math.Max(float64(y) - 4, 0)))
					s := pixelIsBlack(img, x, int(math.Min(float64(y) + 4, float64(height))))
					e := pixelIsBlack(img, int(math.Min(float64(x) + 4, float64(width))), y)
					w := pixelIsBlack(img, int(math.Max(float64(x) - 4, 0)), y)
					if !(n && e && s && w) {
						continue
					}
				}
				cx := math.Round(scale(x, width, HRM_MAX))
				cy := math.Round(scale(y, height, HRM_MAX))
				coords = append(coords, [2]float64{cx, cy})
			}
		}
	}
	return encodeComment(coords)
}

/* Encodes ASCII text into a HRM comment. */
func encodeText(text string) (string, error) {
	if len(text) > 26 {
		return "", fmt.Errorf("Maximum text encoding length is 25 characters.")
	}
	coords := make(Coords, 0)
	fontSize := 16.0
	kerning := 16.0
	lineHeight := 1.25
	for n, c := range text {
		ax := fontSize * float64((n % 13) + 1) + (kerning * float64(n % 13))
		var ay float64
		if len(text) <= 13 {
			ay = IMG_HEIGHT / 2.0
		} else {
			ay = IMG_HEIGHT / 3.0 + IMG_HEIGHT / 3.0 * float64(n / 13) * lineHeight
		}
		segments, err := CharacterSegments(c)
		for i := 0; i < len(segments); i += 1 {
			if segments[i] == DIVIDER {
				coords = append(coords, EMPTY_POINT)
				continue
			}
			x := math.Round(scale(int(ax + segments[i].x * fontSize), IMG_WIDTH, HRM_MAX))
			y := math.Round(scale(int(ay - segments[i].y * fontSize), IMG_HEIGHT, HRM_MAX))
			coords = append(coords, [2]float64{x, y})
		}
		if err != nil {
			return "", err
		}
		if coords[len(coords) - 1] != EMPTY_POINT {
			coords = append(coords, EMPTY_POINT)
		}
	}
	return encodeComment(coords)
}

/* Encodes a sequence of coordinates into HRM comment format. */
func encodeComment(coords Coords) (string, error) {
	// To-do: HRM supports a maximum of 256 unique coordinates
	// For text encoding, construct each character as sequence of segments
	binaryData := encodeCoords(coords)
	var zlibData bytes.Buffer
	z := zlib.NewWriter(&zlibData)
	_, err := z.Write(binaryData)
	z.Close()
	if err != nil {
		return "", err
	}
	b64String := base64.RawStdEncoding.EncodeToString(zlibData.Bytes()) + ";"
	return b64String, nil
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

/* Encodes a slice of (x, y) tuples of a comment into binary data. */
func encodeCoords(coords Coords) []byte {
	size := int(math.Min(float64(len(coords)), 255))
	data := make([]byte, 1028)
	binary.LittleEndian.PutUint32(data, uint32(size))
	for i := 0; i < size; i += 1 {
		x := uint16(coords[i][0])
		y := uint16(coords[i][1])
		binary.LittleEndian.PutUint16(data[4 + 4 * i:], x)
		binary.LittleEndian.PutUint16(data[6 + 4 * i:], y)
	}
	return data
}

/* Decodes the binary data of a comment into a slice of (x, y) tuples. */
func decodeCoords(data []byte) Coords {
	header := binary.LittleEndian.Uint32(data[:4])
	data = data[4:4 * header]
	coords := make(Coords, header)
	for i := 0; i < int(4 * header); i += 4 {
		x := binary.LittleEndian.Uint16(data[i:i + 2])
		y := binary.LittleEndian.Uint16(data[i + 2:i + 4])
		coords[i / 4][0] = float64(x)
		coords[i / 4][1] = float64(y)
	}
	return coords
}

/* Displays coordinates onto an image. */
func displayCoords(coords Coords) {
	ctx := gg.NewContext(IMG_WIDTH, IMG_HEIGHT)
	ctx.SetColor(color.White)
	ctx.DrawRectangle(0, 0, IMG_WIDTH, IMG_HEIGHT)
	ctx.Fill()
	ctx.SetColor(color.Black)
	ctx.SetLineWidth(10)
	segments := make([]Coords, 0)
	var segment Coords = nil
	for i := 0; i < len(coords); i += 1 {
		if coords[i] == [2]float64{0, 0} {
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
			x := scale(int(segments[i][0][0]), HRM_MAX, IMG_WIDTH)
			y := scale(int(segments[i][0][1]), HRM_MAX, IMG_HEIGHT)
			ctx.DrawPoint(x, y, 1)
			ctx.Stroke()
		} else {
			for j := 1; j < len(segments[i]); j += 1 {
				point := segments[i][j]
				x := scale(int(point[0]), HRM_MAX, IMG_WIDTH)
				y := scale(int(point[1]), HRM_MAX, IMG_HEIGHT)
				prevX := scale(int(segments[i][j - 1][0]), HRM_MAX, IMG_WIDTH)
				prevY := scale(int(segments[i][j - 1][1]), HRM_MAX, IMG_HEIGHT)
				ctx.DrawLine(prevX, prevY, x, y)
				ctx.StrokePreserve()
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
			encoded, err = encodePNG(encode, true)
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
		b64String = string(bytes)[:len(bytes) - 1]
	} else {
		b64String = flag.Arg(0)
	}
	coords, err := decodeComment(b64String)
	if err != nil {
		panic(err)
	}
	displayCoords(coords)
}
