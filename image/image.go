package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"golang.org/x/image/webp"
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
	FormatJPG     ImageFormat = "jpg"
	FormatJFIF    ImageFormat = "jfif"
	FormatJPE     ImageFormat = "jpe"
	FormatJPEG    ImageFormat = "jpeg"
	FormatPNG     ImageFormat = "png"
	FormatGIF     ImageFormat = "gif"
	FormatTIF     ImageFormat = "tif"
	FormatTIFF    ImageFormat = "tiff"
	FormatTIFF_FX ImageFormat = "tiff-fx"
	FormatWEBP    ImageFormat = "webp"
	FormatBMP     ImageFormat = "bmp"
	FormatDIB     ImageFormat = "dib"
	FormatXBMP    ImageFormat = "x-bmp"
)

// LoadFromFile loads an image from the given file path.
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

// LoadFromURL downloads and loads an image from the given URL.
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

// LoadFromBytes loads an image from a raw byte slice with the specified format.
func LoadFromBytes(format ImageFormat, data []byte) (*Image, error) {
	buf := bytes.NewBuffer(data)

	return LoadFromReader(format, buf)
}

// LoadFromReader loads an image from an io.Reader with the specified format.
func LoadFromReader(format ImageFormat, r io.Reader) (*Image, error) {
	img, err := decodeTo(format, r)
	if err != nil {
		return nil, err
	}

	return newImage(format, img), nil
}

// SaveToFile saves the image to a file on disk in its current format.
func (img *Image) SaveToFile(path string) error {
	path = filepath.Clean(path)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()

	return img.encodeTo(f)
}

// SaveToWriter writes the image to any io.Writer in its current format.
func (img *Image) SaveToWriter(writer io.Writer) error {
	return img.encodeTo(writer)
}

// ToBytes encodes the image into a byte slice in its current format.
func (img *Image) ToBytes() ([]byte, error) {
	var buf bytes.Buffer
	err := img.encodeTo(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ToBase64 encodes the image as a base64 string in its current format.
func (img *Image) ToBase64() (string, error) {
	b, err := img.ToBytes()
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// Resize creates a new resized copy of the image using the given dimensions and interpolation.
// The original image remains unchanged.
func (img *Image) Resize(width, height uint, interp resize.InterpolationFunction) *Image {
	resized := resize.Resize(width, height, img.Image, interp)

	return newImage(img.Format, resized)
}

// ResizeToWidth resizes the image to the given width while preserving aspect ratio.
// The original image remains unchanged.
func (img *Image) ResizeToWidth(width uint, interp resize.InterpolationFunction) *Image {
	return img.Resize(width, img.Height, interp)
}

// ResizeToHeight resizes the image to the given height while preserving aspect ratio.
// The original image remains unchanged.
func (img *Image) ResizeToHeight(height uint, interp resize.InterpolationFunction) *Image {
	return img.Resize(img.Width, height, interp)
}

// ResizeSelf resizes the image in place to the given dimensions using the specified interpolation.
func (img *Image) ResizeSelf(width, height uint, interp resize.InterpolationFunction) {
	img.Image = resize.Resize(width, height, img.Image, interp)
	img.Width = width
	img.Height = height
	img.ColorModel = img.Image.ColorModel()
}

// Scale creates a new scaled copy of the image by the given factor (e.g., 0.5 for half size).
// The original image remains unchanged.
func (img *Image) Scale(factor float64, interp resize.InterpolationFunction) (*Image, error) {
	if factor <= 0 {
		return nil, fmt.Errorf("factor must be positive")
	}

	w := uint(float64(img.Width) * factor)
	h := uint(float64(img.Height) * factor)

	return img.Resize(w, h, interp), nil
}

// ScaleSelf scales the image in place by the given factor (e.g., 2.0 for double size).
func (img *Image) ScaleSelf(factor float64, interp resize.InterpolationFunction) error {
	if factor <= 0 {
		return fmt.Errorf("factor must be positive")
	}

	w := uint(float64(img.Width) * factor)
	h := uint(float64(img.Height) * factor)
	img.ResizeSelf(w, h, interp)

	return nil
}

/*
ScaleDown creates a new image scaled down to fit within the given max width and height,
preserving aspect ratio. The original image remains unchanged.
*/
func (img *Image) ScaleDown(maxWidth, maxHeight uint, interp resize.InterpolationFunction) (*Image, error) {
	w, h := img.Width, img.Height
	if w <= maxWidth && h <= maxHeight {
		// No need to scale down
		return img, nil
	}

	ratioW := float64(maxWidth) / float64(w)
	ratioH := float64(maxHeight) / float64(h)
	scale := min(ratioW, ratioH)

	return img.Scale(scale, interp)
}

func decodeTo(format ImageFormat, r io.Reader) (image.Image, error) {
	switch format {
	case FormatJPG, FormatJPEG, FormatJFIF, FormatJPE:
		return jpeg.Decode(r)
	case FormatPNG:
		return png.Decode(r)
	case FormatGIF:
		return gif.Decode(r)
	case FormatTIF, FormatTIFF, FormatTIFF_FX:
		return tiff.Decode(r)
	case FormatWEBP:
		return webp.Decode(r)
	case FormatBMP, FormatDIB, FormatXBMP:
		return bmp.Decode(r)
	default:
		return nil, fmt.Errorf("unable to decode image in format %s", format)
	}
}

func (img *Image) encodeTo(w io.Writer) error {
	switch img.Format {
	case FormatJPG, FormatJPEG, FormatJFIF, FormatJPE:
		return jpeg.Encode(w, img.Image, nil)
	case FormatPNG:
		return png.Encode(w, img.Image)
	case FormatGIF:
		return gif.Encode(w, img.Image, nil)
	case FormatTIF, FormatTIFF, FormatTIFF_FX:
		return tiff.Encode(w, img.Image, nil)
	case FormatBMP, FormatDIB, FormatXBMP:
		return bmp.Encode(w, img.Image)
	default:
		return fmt.Errorf("unable to encode image in format %s", img.Format)
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
