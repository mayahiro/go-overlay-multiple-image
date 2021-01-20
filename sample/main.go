package main

import (
	go_overlay_multiple_image "go-overlay-multiple-image"
	"image"
	"image/png"
	"os"
)

func main() {
	var sourceImages []*go_overlay_multiple_image.SourceImage

	f1, err := os.Open("1.png")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	i1, _, err := image.Decode(f1)
	if err != nil {
		panic(err)
	}

	sourceImages = append(sourceImages, &go_overlay_multiple_image.SourceImage{
		Img:     i1,
		OffsetX: 0,
		OffsetY: 0,
	})

	f2, err := os.Open("2.png")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	i2, _, err := image.Decode(f2)
	if err != nil {
		panic(err)
	}

	sourceImages = append(sourceImages, &go_overlay_multiple_image.SourceImage{
		Img:     i2,
		OffsetX: 0,
		OffsetY: 0,
	})

	destImage := go_overlay_multiple_image.CreateImage(sourceImages, i1.Bounds().Size())

	fso, err := os.Create("dest.png")
	if err != nil {
		panic(err)
	}
	defer fso.Close()

	err = png.Encode(fso, destImage)
	if err != nil {
		panic(err)
	}
}
