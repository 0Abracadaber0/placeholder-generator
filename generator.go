package generator

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

func HexToRGBA(hexColor string) (c color.RGBA, err error) {
	hexColor = hexColor[1:]

	intColor, err := strconv.ParseInt(hexColor, 16, 32)
	if err != nil {
		return c, err
	}

	r := uint8(intColor >> 16)
	g := uint8((intColor >> 8) & 0xFF)
	b := uint8(intColor & 0xFF)

	return color.RGBA{R: r, G: g, B: b, A: 0xff}, nil
}

func Generate(height, width int, c color.RGBA, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, c)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Запишите изображение в файл
	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
