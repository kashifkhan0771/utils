package image

import (
	"bytes"
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

	return LoadFromReader(format, f)
}

func LoadFromURL(url string) (*Image, error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
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

	return LoadFromReader(format, resp.Body)
}

func LoadFromBytes(format ImageFormat, data []byte) (*Image, error) {
	buf := bytes.NewBuffer(data)

	return LoadFromReader(format, buf)
}

func LoadFromReader(format ImageFormat, r io.Reader) (*Image, error) {
	img, err := decodeImage(format, r)
	if err != nil {
		return nil, err
	}

	return new(format, img), nil
}

func (img *Image) SaveToFile(path string) error {
	path = filepath.Clean(path)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	return img.encodeImage(f)
}

func (img *Image) SaveToWriter(writer io.Writer) error {
	return img.encodeImage(writer)
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

func (img *Image) encodeImage(w io.Writer) error {
	switch img.Format {
	case FormatJPG, FormatJPEG, FormatJFIF, FormatJPE:
		return jpeg.Encode(w, img.Image, nil)
	case FormatPNG:
		return png.Encode(w, img.Image)
	default:
		return fmt.Errorf("invalid or unsupported image format: %s", img.Format)
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
