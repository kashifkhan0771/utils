package image

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Image struct {
	Image  image.Image
	Format ImageFormat
	Width  int
	Height int

	// Metadata
	ColorModel color.Model
}

type ImageFormat string

const (
	FormatJPG ImageFormat = "jpg"
	FormatPNG ImageFormat = "png"
)

func LoadFromFile(path string) (*Image, error) {
	path = filepath.Clean(path)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var img image.Image
	format := ImageFormat(filepath.Ext(path)[1:])

	switch format {
	case FormatJPG:
		img, err = jpeg.Decode(f)
	case FormatPNG:
		img, err = png.Decode(f)
	default:
		return nil, fmt.Errorf("invalid or unsupported image format: %s", format)
	}

	if err != nil {
		return nil, err
	}

	return &Image{
		Image:      img,
		Format:     format,
		Width:      img.Bounds().Max.X - img.Bounds().Min.X,
		Height:     img.Bounds().Max.Y - img.Bounds().Min.Y,
		ColorModel: img.ColorModel(),
	}, nil
}
