// Package pdf provides utilities for creating and manipulating PDF documents.
//
// This package includes functions for:
//   - Converting HTML content to PDF documents
//   - Extracting plain text from PDF files
//   - Merging multiple PDF files into a single document
//   - Splitting a PDF into multiple files based on page ranges
package pdf

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/ledongthuc/pdf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"golang.org/x/net/html"
)

// fontState holds font properties for the context stack used during HTML rendering.
type fontState struct {
	family string
	style  string
	size   float64
}

// ConvertHTMLToPDF converts an HTML string into a PDF document and saves it to
// the specified output path.
//
// Supported HTML elements:
//   - Headings: <h1> through <h6>
//   - Paragraphs: <p>
//   - Bold: <b>, <strong>
//   - Italic: <i>, <em>
//   - Line breaks: <br>
//   - Links: <a href="..."> (rendered as underlined text)
//   - Images: <img src="..."> (local file paths only)
//
// This function does not support CSS stylesheets, JavaScript, or complex HTML
// layouts such as tables or flexbox.
func ConvertHTMLToPDF(htmlContent string, outputPath string) error {
	if strings.TrimSpace(htmlContent) == "" {
		return fmt.Errorf("HTML content cannot be empty")
	}

	outputPath = filepath.Clean(outputPath)

	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return fmt.Errorf("failed to parse HTML: %w", err)
	}

	p := fpdf.New("P", "mm", "A4", "")
	p.SetAutoPageBreak(true, 15)
	p.AddPage()
	p.SetFont("Arial", "", 12)

	renderNode(p, doc, nil)

	if p.Err() {
		return fmt.Errorf("PDF generation error: %w", p.Error())
	}

	return p.OutputFileAndClose(outputPath)
}

// ConvertHTMLFileToPDF reads an HTML file at inputPath and converts it to a PDF
// saved at outputPath. Supports the same HTML elements as ConvertHTMLToPDF.
func ConvertHTMLFileToPDF(inputPath string, outputPath string) error {
	inputPath = filepath.Clean(inputPath)

	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %w", err)
	}

	return ConvertHTMLToPDF(string(data), outputPath)
}

// renderNode recursively walks the HTML node tree and renders content to the PDF.
// stack carries the font state of ancestor elements so inline elements can restore
// the parent font on exit rather than blindly resetting to a default.
func renderNode(p *fpdf.Fpdf, n *html.Node, stack []fontState) {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			p.Write(6, toLatin1(text)+" ")
		}
		return
	}

	if n.Type != html.ElementNode {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			renderNode(p, c, stack)
		}
		return
	}

	// Skip non-visual elements — their text must not appear in output.
	switch n.Data {
	case "script", "style", "head", "noscript", "iframe", "svg", "canvas":
		return
	}

	// Derive current font from stack; fall back to document default.
	cur := fontState{"Arial", "", 12}
	if len(stack) > 0 {
		cur = stack[len(stack)-1]
	}

	// lineH returns line height matching font size.
	lineH := func(size float64) float64 { return size * 0.45 }

	var pushed bool
	var isBlock bool

	switch n.Data {
	case "h1":
		p.Ln(lineH(24) + 2)
		p.SetFont("Arial", "B", 24)
		stack = append(stack, fontState{"Arial", "B", 24})
		pushed, isBlock = true, true
	case "h2":
		p.Ln(lineH(20) + 2)
		p.SetFont("Arial", "B", 20)
		stack = append(stack, fontState{"Arial", "B", 20})
		pushed, isBlock = true, true
	case "h3":
		p.Ln(lineH(16) + 1)
		p.SetFont("Arial", "B", 16)
		stack = append(stack, fontState{"Arial", "B", 16})
		pushed, isBlock = true, true
	case "h4":
		p.Ln(lineH(14) + 1)
		p.SetFont("Arial", "B", 14)
		stack = append(stack, fontState{"Arial", "B", 14})
		pushed, isBlock = true, true
	case "h5":
		p.Ln(lineH(12))
		p.SetFont("Arial", "B", 12)
		stack = append(stack, fontState{"Arial", "B", 12})
		pushed, isBlock = true, true
	case "h6":
		p.Ln(lineH(10))
		p.SetFont("Arial", "B", 10)
		stack = append(stack, fontState{"Arial", "B", 10})
		pushed, isBlock = true, true
	case "p":
		p.Ln(lineH(12) + 2)
		p.SetFont("Arial", "", 12)
		stack = append(stack, fontState{"Arial", "", 12})
		pushed, isBlock = true, true
	case "div", "section", "article", "main", "header", "footer", "nav":
		p.Ln(3)
		isBlock = true
	case "ul", "ol":
		p.Ln(2)
		isBlock = true
	case "li":
		p.Ln(1)
		p.Write(6, "  - ")
		isBlock = true
	case "b", "strong":
		p.SetFont(cur.family, "B", cur.size)
		stack = append(stack, fontState{cur.family, "B", cur.size})
		pushed = true
	case "i", "em":
		p.SetFont(cur.family, "I", cur.size)
		stack = append(stack, fontState{cur.family, "I", cur.size})
		pushed = true
	case "a":
		p.SetFont(cur.family, "U", cur.size)
		stack = append(stack, fontState{cur.family, "U", cur.size})
		pushed = true
	case "br":
		p.Ln(5)
	case "hr":
		p.Ln(3)
		pw, _ := p.GetPageSize()
		p.Line(10, p.GetY(), pw-10, p.GetY())
		p.Ln(3)
	case "img":
		renderImage(p, n)
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		renderNode(p, c, stack)
	}

	if pushed {
		stack = stack[:len(stack)-1]
		parent := fontState{"Arial", "", 12}
		if len(stack) > 0 {
			parent = stack[len(stack)-1]
		}
		p.SetFont(parent.family, parent.style, parent.size)

		if isBlock {
			p.Ln(lineH(cur.size) + 1)
			p.SetFont("Arial", "", 12)
		}
	} else if isBlock {
		p.Ln(2)
	}
}

// renderImage attempts to add an image from a local file path to the PDF.
func renderImage(p *fpdf.Fpdf, n *html.Node) {
	var src string
	for _, attr := range n.Attr {
		if attr.Key == "src" {
			src = attr.Val
			break
		}
	}

	if src == "" {
		return
	}

	// Only support local file paths
	src = filepath.Clean(src)
	if _, err := os.Stat(src); err != nil {
		return
	}

	// Constrain image to page width (margins ~10mm each side).
	pageWidth, _ := p.GetPageSize()
	pageWidth -= 20
	p.Ln(4)
	p.Image(src, -1, -1, pageWidth, 0, true, "", 0, "")
	p.Ln(4)
}

// ExtractTextFromPDF extracts the plain text content from all pages of a PDF
// file and returns it as a single string. Pages are separated by a newline.
func ExtractTextFromPDF(inputPath string) (string, error) {
	inputPath = filepath.Clean(inputPath)

	f, r, err := pdf.Open(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to open PDF: %w", err)
	}
	defer f.Close()

	totalPages := r.NumPage()
	if totalPages == 0 {
		return "", nil
	}

	var buf bytes.Buffer
	for i := 1; i <= totalPages; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return "", fmt.Errorf("failed to extract text from page %d: %w", i, err)
		}

		if buf.Len() > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(text)
	}

	return buf.String(), nil
}

// MergePDFs combines multiple PDF files into a single output document.
// The pages from each input file are appended in the order provided.
// At least two input files must be specified.
func MergePDFs(inputFiles []string, outputFile string) error {
	if len(inputFiles) < 2 {
		return fmt.Errorf("at least two input files are required for merging")
	}

	outputFile = filepath.Clean(outputFile)

	cleaned := make([]string, len(inputFiles))
	for i, f := range inputFiles {
		cleaned[i] = filepath.Clean(f)
		if _, err := os.Stat(cleaned[i]); err != nil {
			return fmt.Errorf("input file %q does not exist: %w", cleaned[i], err)
		}
	}

	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed

	return api.MergeCreateFile(cleaned, outputFile, false, conf)
}

// SplitPDF splits a PDF file into multiple smaller PDF files based on the
// specified page ranges. Each range produces a separate output PDF.
//
// Page ranges are specified as strings such as "1-3", "5", or "7-10".
// Output files are saved in the specified directory with names like
// "pages_1-3.pdf", "pages_5-5.pdf", etc.
func SplitPDF(inputFile string, pageRanges []string, outputDir string) error {
	if len(pageRanges) == 0 {
		return fmt.Errorf("at least one page range is required")
	}

	inputFile = filepath.Clean(inputFile)
	outputDir = filepath.Clean(outputDir)

	if _, err := os.Stat(inputFile); err != nil {
		return fmt.Errorf("input file %q does not exist: %w", inputFile, err)
	}

	// Get total page count for range validation.
	totalPages, err := api.PageCountFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read page count: %w", err)
	}

	if err := os.MkdirAll(outputDir, 0750); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed

	for _, rangeStr := range pageRanges {
		start, end, err := parsePageRange(rangeStr)
		if err != nil {
			return fmt.Errorf("invalid page range %q: %w", rangeStr, err)
		}

		if end > totalPages {
			return fmt.Errorf("page range %q exceeds document page count (%d)", rangeStr, totalPages)
		}

		outputFile := filepath.Join(outputDir, fmt.Sprintf("pages_%d-%d.pdf", start, end))

		// pdfcpu uses page selection strings like "1-3"
		pageSelection := fmt.Sprintf("%d-%d", start, end)

		if err := api.TrimFile(inputFile, outputFile, []string{pageSelection}, conf); err != nil {
			return fmt.Errorf("failed to split pages %s: %w", rangeStr, err)
		}
	}

	return nil
}
