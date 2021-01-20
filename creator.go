package go_overlay_multiple_image

import (
	"image"
	"image/draw"
)

type SourceImage struct {
	Img     image.Image
	OffsetX int
	OffsetY int
}

func CreateImage(sourceImages []*SourceImage, size image.Point) *image.RGBA {
	rgba := image.NewRGBA(image.Rectangle{image.Point{0, 0}, size})

	for _, sourceImage := range sourceImages {
		draw.Draw(
			rgba,
			image.Rectangle{image.Point{0, 0}, size},
			sourceImage.Img,
			image.Point{-(size.X / 2) + (sourceImage.Img.Bounds().Size().X / 2), -(size.Y / 2) + (sourceImage.Img.Bounds().Size().Y / 2)},
			draw.Over,
		)
	}

	return rgba
}
