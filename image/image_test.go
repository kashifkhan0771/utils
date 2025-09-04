package image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"

	"github.com/nfnt/resize"
)

var makeRGBA = func(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range height {
		for x := range width {
			img.Set(x, y, color.RGBA{R: 255, G: 128, B: 64, A: 255})
		}
	}
	return img
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		name    string
		width   int
		height  int
		format  ImageFormat
		wantErr bool
	}{
		{"PNG_1x1", 1, 1, FormatPNG, false},
		{"PNG_2x2", 2, 2, FormatPNG, false},
		{"JPEG_5x3", 5, 3, FormatJPEG, false},
		{"JFIF_3x5", 3, 5, FormatJFIF, false},
		{"JPE_10x5", 10, 5, FormatJPE, false},
		{"GIF_8x8", 8, 8, FormatGIF, false},
		{"TIF_16x16", 16, 16, FormatTIF, false},
		{"TIFF_32x16", 32, 16, FormatTIFF, false},
		{"TIFF_FX_64x64", 64, 64, FormatTIFF_FX, false},
		{"WEBP_2x2", 2, 2, FormatWEBP, true},
		{"BMP_5x3", 5, 3, FormatBMP, false},
		{"DIB_3x5", 3, 5, FormatDIB, false},
		{"XBMP_10x5", 10, 5, FormatXBMP, false},
		{"Unsupported_2x2", 2, 2, "unsupported", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				Image:  makeRGBA(tt.width, tt.height),
				Format: tt.format,
				// #nosec G115
				Width: uint(tt.width),
				// #nosec G115
				Height: uint(tt.height),
			}

			data, err := img.ToBytes()
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToBytes() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if err == nil && len(data) == 0 {
				t.Errorf("ToBytes() returned empty byte slice")
			}
		})
	}
}

func TestToBase64(t *testing.T) {
	tests := []struct {
		name    string
		width   int
		height  int
		format  ImageFormat
		wantErr bool
	}{
		{"PNG_1x1", 1, 1, FormatPNG, false},
		{"PNG_2x2", 2, 2, FormatPNG, false},
		{"JPEG_5x3", 5, 3, FormatJPEG, false},
		{"JFIF_3x5", 3, 5, FormatJFIF, false},
		{"JPE_10x5", 10, 5, FormatJPE, false},
		{"GIF_8x8", 8, 8, FormatGIF, false},
		{"TIF_16x16", 16, 16, FormatTIF, false},
		{"TIFF_32x16", 32, 16, FormatTIFF, false},
		{"TIFF_FX_64x64", 64, 64, FormatTIFF_FX, false},
		{"WEBP_2x2", 2, 2, FormatWEBP, true},
		{"BMP_5x3", 5, 3, FormatBMP, false},
		{"DIB_3x5", 3, 5, FormatDIB, false},
		{"XBMP_10x5", 10, 5, FormatXBMP, false},
		{"Unsupported_2x2", 2, 2, "unsupported", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				Image:  makeRGBA(tt.width, tt.height),
				Format: tt.format,
				// #nosec G115
				Width: uint(tt.width),
				// #nosec G115
				Height: uint(tt.height),
			}

			b64, err := img.ToBase64()
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToBase64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if len(b64) == 0 {
					t.Errorf("ToBase64() returned empty string")
				}

				data, err := base64.StdEncoding.DecodeString(b64)
				if err != nil {
					t.Errorf("base64 decoding failed: %v", err)
				}
				if len(data) == 0 {
					t.Errorf("decoded base64 is empty")
				}
			}
		})
	}
}

func TestResize(t *testing.T) {
	tests := []struct {
		name    string
		origW   uint
		origH   uint
		targetW uint
		targetH uint
	}{
		{"smaller", 8, 8, 4, 4},
		{"larger", 4, 4, 8, 8},
		{"rectangular", 5, 10, 10, 5},
		{"same", 3, 3, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				// #nosec G115
				Image:  makeRGBA(int(tt.origW), int(tt.origH)),
				Format: FormatPNG,
				Width:  tt.origW,
				Height: tt.origH,
			}

			newImg := img.Resize(tt.targetW, tt.targetH, resize.NearestNeighbor)

			if newImg.Width != tt.targetW || newImg.Height != tt.targetH {
				t.Errorf("Resize() dimensions = %dx%d, want %dx%d", newImg.Width, newImg.Height, tt.targetW, tt.targetH)
			}

			// Original image should remain unchanged
			if img.Width != tt.origW || img.Height != tt.origH {
				t.Errorf("Resize() modified original image dimensions")
			}
		})
	}
}

func TestResizeSelf(t *testing.T) {
	tests := []struct {
		name    string
		origW   uint
		origH   uint
		targetW uint
		targetH uint
	}{
		{"smaller", 8, 8, 4, 4},
		{"larger", 4, 4, 8, 8},
		{"rectangular", 5, 10, 10, 5},
		{"same", 3, 3, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				// #nosec G115
				Image:  makeRGBA(int(tt.origW), int(tt.origH)),
				Format: FormatPNG,
				Width:  tt.origW,
				Height: tt.origH,
			}

			img.ResizeSelf(tt.targetW, tt.targetH, resize.NearestNeighbor)

			if img.Width != tt.targetW || img.Height != tt.targetH {
				t.Errorf("ResizeSelf() dimensions = %dx%d, want %dx%d", img.Width, img.Height, tt.targetW, tt.targetH)
			}
		})
	}
}

func TestResizeToWidth(t *testing.T) {
	tests := []struct {
		name    string
		origW   uint
		origH   uint
		targetW uint
		targetH uint
	}{
		{"increase_width", 4, 4, 8, 8},
		{"decrease_width", 8, 4, 4, 2},
		{"same_width", 5, 3, 5, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				// #nosec G115
				Image:  makeRGBA(int(tt.origW), int(tt.origH)),
				Format: FormatPNG,
				Width:  tt.origW,
				Height: tt.origH,
			}

			newImg := img.ResizeToWidth(tt.targetW, resize.NearestNeighbor)

			if newImg.Width != tt.targetW {
				t.Errorf("ResizeToWidth() width = %d, want %d", newImg.Width, tt.targetW)
			}

			if newImg.Height != tt.targetH {
				t.Errorf("ResizeToWidth() height = %d, want %d", newImg.Height, tt.targetH)
			}

			// Original image should stay unchanged
			if img.Width != tt.origW || img.Height != tt.origH {
				t.Errorf("ResizeToWidth() modified original image")
			}
		})
	}
}

func TestResizeToHeight(t *testing.T) {
	tests := []struct {
		name    string
		origW   uint
		origH   uint
		targetH uint
		targetW uint
	}{
		{"increase_height", 4, 4, 8, 8},
		{"decrease_height", 4, 8, 4, 2},
		{"same_height", 5, 3, 3, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				// #nosec G115
				Image:  makeRGBA(int(tt.origW), int(tt.origH)),
				Format: FormatPNG,
				Width:  tt.origW,
				Height: tt.origH,
			}

			newImg := img.ResizeToHeight(tt.targetH, resize.NearestNeighbor)

			if newImg.Height != tt.targetH {
				t.Errorf("ResizeToHeight() height = %d, want %d", newImg.Height, tt.targetH)
			}

			if newImg.Width != tt.targetW {
				t.Errorf("ResizeToHeight() width = %d, want %d", newImg.Width, tt.targetW)
			}

			// Original image should stay unchanged
			if img.Width != tt.origW || img.Height != tt.origH {
				t.Errorf("ResizeToHeight() modified original image")
			}
		})
	}
}

func TestScale(t *testing.T) {
	tests := []struct {
		name    string
		origW   uint
		origH   uint
		factor  float64
		wantW   uint
		wantH   uint
		wantErr bool
	}{
		{"scale_up", 4, 4, 2.0, 8, 8, false},
		{"scale_down", 10, 10, 0.5, 5, 5, false},
		{"scale_same", 6, 3, 1.0, 6, 3, false},
		{"scale_invalid_zero", 5, 5, 0.0, 0, 0, true},
		{"scale_invalid_negative", 5, 5, -1.0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				// #nosec G115
				Image:  makeRGBA(int(tt.origW), int(tt.origH)),
				Format: FormatPNG,
				Width:  tt.origW,
				Height: tt.origH,
			}

			newImg, err := img.Scale(tt.factor, resize.NearestNeighbor)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Scale() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if newImg.Width != tt.wantW || newImg.Height != tt.wantH {
					t.Errorf("Scale() dimensions = %dx%d, want %dx%d", newImg.Width, newImg.Height, tt.wantW, tt.wantH)
				}

				// Original should remain unchanged
				if img.Width != tt.origW || img.Height != tt.origH {
					t.Errorf("Scale() modified original image")
				}
			}
		})
	}
}

func TestScaleSelf(t *testing.T) {
	tests := []struct {
		name    string
		origW   uint
		origH   uint
		factor  float64
		wantW   uint
		wantH   uint
		wantErr bool
	}{
		{"scale_up", 4, 4, 2.0, 8, 8, false},
		{"scale_down", 10, 10, 0.5, 5, 5, false},
		{"scale_same", 6, 3, 1.0, 6, 3, false},
		{"scale_invalid_zero", 5, 5, 0.0, 0, 0, true},
		{"scale_invalid_negative", 5, 5, -1.0, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				// #nosec G115
				Image:  makeRGBA(int(tt.origW), int(tt.origH)),
				Format: FormatPNG,
				Width:  tt.origW,
				Height: tt.origH,
			}

			err := img.ScaleSelf(tt.factor, resize.NearestNeighbor)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ScaleSelf() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if img.Width != tt.wantW || img.Height != tt.wantH {
					t.Errorf("ScaleSelf() dimensions = %dx%d, want %dx%d", img.Width, img.Height, tt.wantW, tt.wantH)
				}
			}
		})
	}
}

func TestScaleDown(t *testing.T) {
	tests := []struct {
		name     string
		origW    uint
		origH    uint
		maxW     uint
		maxH     uint
		wantW    uint
		wantH    uint
		wantSame bool
	}{
		{"within_limits", 100, 50, 200, 200, 100, 50, true},
		{"scale_down_width", 300, 100, 150, 200, 150, 50, false},
		{"scale_down_height", 100, 400, 200, 200, 50, 200, false},
		{"scale_down_both", 400, 400, 100, 100, 100, 100, false},
		{"exact_match", 200, 200, 200, 200, 200, 200, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := &Image{
				// #nosec G115
				Image:  makeRGBA(int(tt.origW), int(tt.origH)),
				Format: FormatPNG,
				Width:  tt.origW,
				Height: tt.origH,
			}

			newImg, err := img.ScaleDown(tt.maxW, tt.maxH, resize.NearestNeighbor)
			if err != nil {
				t.Fatalf("ScaleDown() error = %v", err)
			}

			if newImg.Width != tt.wantW || newImg.Height != tt.wantH {
				t.Errorf("ScaleDown() dimensions = %dx%d, want %dx%d",
					newImg.Width, newImg.Height, tt.wantW, tt.wantH)
			}

			// If within limits, should return original object (not new one)
			if tt.wantSame && newImg != img {
				t.Errorf("ScaleDown() returned new object, want same original")
			}
		})
	}
}

func BenchmarkDecodeTo(b *testing.B) {
	img := makeRGBA(128, 128)

	var pngBuf, jpegBuf, gifBuf bytes.Buffer
	if err := png.Encode(&pngBuf, img); err != nil {
		b.Fatalf("png.Encode() failed: %v", err)
	}

	if err := jpeg.Encode(&jpegBuf, img, nil); err != nil {
		b.Fatalf("jpeg.Encode() failed: %v", err)
	}

	if err := gif.Encode(&gifBuf, img, nil); err != nil {
		b.Fatalf("gif.Encode() failed: %v", err)
	}

	benches := []struct {
		name   string
		format ImageFormat
		data   []byte
	}{
		{"PNG", FormatPNG, pngBuf.Bytes()},
		{"JPEG", FormatJPEG, jpegBuf.Bytes()},
		{"GIF", FormatGIF, gifBuf.Bytes()},
	}

	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := bytes.NewReader(bb.data)
				if _, err := decodeTo(bb.format, r); err != nil {
					b.Fatalf("decodeTo(%s) failed: %v", bb.name, err)
				}
			}
		})
	}
}

func BenchmarkEncodeTo(b *testing.B) {
	img := makeRGBA(128, 128)

	benches := []struct {
		name   string
		format ImageFormat
	}{
		{"PNG", FormatPNG},
		{"JPEG", FormatJPEG},
		{"GIF", FormatGIF},
	}

	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			wrapped := &Image{
				Image:  img,
				Format: bb.format,
				Width:  128,
				Height: 128,
			}

			var buf bytes.Buffer
			for i := 0; i < b.N; i++ {
				buf.Reset()
				if err := wrapped.encodeTo(&buf); err != nil {
					b.Fatalf("encodeTo(%s) failed: %v", bb.name, err)
				}
			}
		})
	}
}
