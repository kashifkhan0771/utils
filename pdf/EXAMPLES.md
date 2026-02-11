## PDF Function Examples

### Convert HTML to PDF

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pdf"
)

func main() {
	html := `
		<h1>Invoice #1234</h1>
		<p>Date: 2026-01-15</p>
		<p><b>Bill To:</b> John Doe</p>
		<p>Thank you for your purchase!</p>
	`

	err := pdf.ConvertHTMLToPDF(html, "invoice.pdf")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PDF created successfully!")
}
```

#### Output:

```
PDF created successfully!
```

---

### Extract Text from PDF

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pdf"
)

func main() {
	text, err := pdf.ExtractTextFromPDF("document.pdf")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Extracted text:")
	fmt.Println(text)
}
```

#### Output:

```
Extracted text:
This is the content of the PDF document.
It can span multiple pages.
```

---

### Merge Multiple PDFs

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pdf"
)

func main() {
	inputFiles := []string{
		"chapter1.pdf",
		"chapter2.pdf",
		"chapter3.pdf",
	}

	err := pdf.MergePDFs(inputFiles, "combined_book.pdf")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PDFs merged successfully!")
}
```

#### Output:

```
PDFs merged successfully!
```

---

### Split a PDF by Page Ranges

```go
package main

import (
	"fmt"

	"github.com/kashifkhan0771/utils/pdf"
)

func main() {
	pageRanges := []string{"1-3", "5", "7-10"}

	err := pdf.SplitPDF("large_document.pdf", pageRanges, "./split_output")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PDF split successfully!")
	fmt.Println("Output files: pages_1-3.pdf, pages_5-5.pdf, pages_7-10.pdf")
}
```

#### Output:

```
PDF split successfully!
Output files: pages_1-3.pdf, pages_5-5.pdf, pages_7-10.pdf
```

---
