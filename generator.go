package generator

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Generate(height, width int, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	c := color.RGBA{R: 128, G: 128, B: 128, A: 255}
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
