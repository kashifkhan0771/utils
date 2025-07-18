package image

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"net/http"
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
	FormatJPG  ImageFormat = "jpg"
	FormatJFIF ImageFormat = "jfif"
	FormatJPE  ImageFormat = "jpe"
	FormatJPEG ImageFormat = "jpeg"
	FormatPNG  ImageFormat = "png"
)

func LoadFromFile(path string) (*Image, error) {
	path = filepath.Clean(path)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	format := ImageFormat(filepath.Ext(path)[1:])
	img, err := decodeImage(format, f)
	if err != nil {
		return nil, err
	}

	return new(format, img), nil
}

func LoadFromURL(url string) (*Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	extensions, err := mime.ExtensionsByType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	if len(extensions) == 0 {
		return nil, fmt.Errorf("unable to determine image format: no file extensions found for Content-Type '%s'", resp.Header.Get("Content-Type"))
	}

	format := ImageFormat(extensions[0][1:])
	img, err := decodeImage(format, resp.Body)
	if err != nil {
		return nil, err
	}

	return new(format, img), nil
}

func decodeImage(format ImageFormat, r io.Reader) (image.Image, error) {
	switch format {
	case FormatJPG, FormatJPEG, FormatJFIF, FormatJPE:
		return jpeg.Decode(r)
	case FormatPNG:
		return png.Decode(r)
	default:
		return nil, fmt.Errorf("invalid or unsupported image format: %s", format)
	}
}

func new(format ImageFormat, img image.Image) *Image {
	return &Image{
		Image:      img,
		Format:     format,
		Width:      img.Bounds().Max.X - img.Bounds().Min.X,
		Height:     img.Bounds().Max.Y - img.Bounds().Min.Y,
		ColorModel: img.ColorModel(),
	}
}
