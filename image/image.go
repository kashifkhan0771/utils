package image

import (
	"bytes"
	"encoding/base64"
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

	"github.com/nfnt/resize"
)

type Image struct {
	Image  image.Image
	Format ImageFormat
	Width  uint
	Height uint

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
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
	img, err := decodeTo(format, r)
	if err != nil {
		return nil, err
	}

	return newImage(format, img), nil
}

func (img *Image) SaveToFile(path string) error {
	path = filepath.Clean(path)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	return img.encodeTo(f)
}

func (img *Image) SaveToWriter(writer io.Writer) error {
	return img.encodeTo(writer)
}

func (img *Image) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	err := img.encodeTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (img *Image) ToBase64() (string, error) {
	b, err := img.ToBytes()
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func (img *Image) Resize(width, height uint, interp resize.InterpolationFunction) *Image {
	resized := resize.Resize(width, height, img.Image, interp)

	return newImage(img.Format, resized)
}

func (img *Image) ResizeToWidth(width uint, interp resize.InterpolationFunction) *Image {
	return img.Resize(width, img.Height, interp)
}

func (img *Image) ResizeToHeight(height uint, interp resize.InterpolationFunction) *Image {
	return img.Resize(img.Width, height, interp)
}

func (img *Image) ResizeSelf(width, height uint, interp resize.InterpolationFunction) {
	img.Image = resize.Resize(width, height, img.Image, interp)
	img.Width = width
	img.Height = height
	img.ColorModel = img.Image.ColorModel()
}

func decodeTo(format ImageFormat, r io.Reader) (image.Image, error) {
	switch format {
	case FormatJPG, FormatJPEG, FormatJFIF, FormatJPE:
		return jpeg.Decode(r)
	case FormatPNG:
		return png.Decode(r)
	default:
		return nil, fmt.Errorf("invalid or unsupported image format: %s", format)
	}
}

func (img *Image) encodeTo(w io.Writer) error {
	switch img.Format {
	case FormatJPG, FormatJPEG, FormatJFIF, FormatJPE:
		return jpeg.Encode(w, img.Image, nil)
	case FormatPNG:
		return png.Encode(w, img.Image)
	default:
		return fmt.Errorf("invalid or unsupported image format: %s", img.Format)
	}
}

func newImage(format ImageFormat, img image.Image) *Image {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	if width < 0 || height < 0 {
		panic("invalid image bounds: negative width or height")
	}

	return &Image{
		Image:      img,
		Format:     format,
		Width:      uint(width),
		Height:     uint(height),
		ColorModel: img.ColorModel(),
	}
}
