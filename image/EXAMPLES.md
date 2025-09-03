## Image Utilities examples

### Load image from a file

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
)

func main() {
	path := "example.png"

	img, err := image.LoadFromFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded image: format=%s, size=%dx%d\n", img.Format, img.Width, img.Height)
}
```

#### Output
`Loaded image: format=png, size=150x150`


### Load an image from URL

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
)

func main() {
	url := "https://via.placeholder.com/150.png"

	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded image: format=%s, size=%dx%d\n", img.Format, img.Width, img.Height)
}
```

#### Output
`Loaded image: format=png, size=150x150`


### Convert image to bytes and Base64

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
)

func main() {
	url := "https://via.placeholder.com/100.jpg"
	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := img.ToBytes()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Byte length:", len(data))

	base64Str, err := img.ToBase64()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Base64 length:", len(base64Str))
}
```

#### Output
`Byte length: 829
Base64 length: 1108`

### Resize an image

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
	"github.com/nfnt/resize"
)

func main() {
	url := "https://via.placeholder.com/200.png"
	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	resized := img.Resize(100, 50, resize.Bilinear)
	fmt.Printf("Resized image: %dx%d\n", resized.Width, resized.Height)
}
```

#### Output
`Resized image: 100x50`

### Resize image in place

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
	"github.com/nfnt/resize"
)

func main() {
	url := "https://via.placeholder.com/120.png"
	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	img.ResizeSelf(60, 30, resize.NearestNeighbor)
	fmt.Printf("Image resized in place: %dx%d\n", img.Width, img.Height)
}
```

#### Output
`Image resized in place: 60x30`

### Scale an image

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
	"github.com/nfnt/resize"
)

func main() {
	url := "https://via.placeholder.com/200.png"
	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	scaled, err := img.Scale(0.5, resize.Bilinear)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Scaled image: %dx%d\n", scaled.Width, scaled.Height)
}
```

#### Output
`Scaled image: 100x100`


### Scale image in place

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
	"github.com/nfnt/resize"
)

func main() {
	url := "https://via.placeholder.com/150.png"
	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	err = img.ScaleSelf(0.25, resize.Bicubic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Scaled in place: %dx%d\n", img.Width, img.Height)
}
```

#### Output
`Scaled in place: 37x37`


### Scale down to fit maximum size

```go
package main

import (
	"fmt"
	"log"

	"github.com/kashifkhan0771/utils/image"
	"github.com/nfnt/resize"
)

func main() {
	url := "https://via.placeholder.com/400x200.png"
	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	scaled, err := img.ScaleDown(150, 150, resize.Lanczos3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Scaled down: %dx%d\n", scaled.Width, scaled.Height)
}
```

#### Output
`Scaled down: 150x75`

### Save image to file

```go
package main

import (
	"log"

	"github.com/kashifkhan0771/utils/image"
)

func main() {
	url := "https://via.placeholder.com/100.png"
	img, err := image.LoadFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	if err := img.SaveToFile("output.png"); err != nil {
		log.Fatal(err)
	}
}
```
