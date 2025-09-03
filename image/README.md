# Image Utilities

The `image` package provides a set of utilities for loading, manipulating, encoding, and decoding images in various formats. Below are the main features and methods of the package:

---

### **LoadFromFile**

- **`LoadFromFile(path string) (*Image, error)`**  
  Loads an image from a local file.
  - **`path`**: Path to the image file.
  - Returns an `*Image` object and an error if loading fails.
  - Supports formats: JPG, JPEG, PNG, GIF, BMP, TIFF, WEBP, and more.

---

### **LoadFromURL**

- **`LoadFromURL(url string) (*Image, error)`**  
  Loads an image from a URL.
  - **`url`**: The image URL.
  - Returns an `*Image` object and an error if loading fails.
  - Supports content types: PNG, JPEG, GIF, BMP, TIFF, WEBP.

---

### **LoadFromBytes**

- **`LoadFromBytes(format ImageFormat, data []byte) (*Image, error)`**  
  Loads an image from a byte slice.
  - **`format`**: The format of the image (e.g., `FormatPNG`, `FormatJPEG`).
  - **`data`**: Byte slice containing image data.
  - Returns an `*Image` object.

---

### **SaveToFile**

- **`(img *Image) SaveToFile(path string) error`**  
  Saves an image to a local file.
  - **`path`**: Path to save the image.
  - Supports all formats supported by the `Image` object.

---

### **SaveToWriter**

- **`(img *Image) SaveToWriter(writer io.Writer) error`**  
  Saves an image to any `io.Writer`.
  - **`writer`**: Destination writer.

---

### **ToBytes**

- **`(img *Image) ToBytes() ([]byte, error)`**  
  Converts the image into a byte slice.
  - Returns the encoded image data.

---

### **ToBase64**

- **`(img *Image) ToBase64() (string, error)`**  
  Converts the image to a base64-encoded string.
  - Returns the base64 string representation of the image.

---

### **Resize**

- **`(img *Image) Resize(width, height uint, interp resize.InterpolationFunction) *Image`**  
  Returns a resized copy of the image.
  - **`width`**: Target width.
  - **`height`**: Target height.
  - **`interp`**: Interpolation function (e.g., `resize.NearestNeighbor`, `resize.Bilinear`).

---

### **ResizeToWidth**

- **`(img *Image) ResizeToWidth(width uint, interp resize.InterpolationFunction) *Image`**  
  Returns a resized copy with a new width; height remains unchanged.

---

### **ResizeToHeight**

- **`(img *Image) ResizeToHeight(height uint, interp resize.InterpolationFunction) *Image`**  
  Returns a resized copy with a new height; width remains unchanged.

---

### **ResizeSelf**

- **`(img *Image) ResizeSelf(width, height uint, interp resize.InterpolationFunction)`**  
  Resizes the original image in place.

---

### **Scale**

- **`(img *Image) Scale(factor float64, interp resize.InterpolationFunction) (*Image, error)`**  
  Returns a resized copy scaled by the given factor.
  - **`factor`**: Scaling factor (must be positive).

---

### **ScaleSelf**

- **`(img *Image) ScaleSelf(factor float64, interp resize.InterpolationFunction) error`**  
  Scales the original image in place by the given factor.

---

### **ScaleDown**

- **`(img *Image) ScaleDown(maxWidth, maxHeight uint, interp resize.InterpolationFunction) (*Image, error)`**  
  Scales down the image to fit within `maxWidth` and `maxHeight` while maintaining the aspect ratio.
  - If the image is already within the bounds, the original object is returned.

---

## Examples

For examples of using each function, please check out [EXAMPLES.md](/image/EXAMPLES.md)

