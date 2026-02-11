package pdf

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-pdf/fpdf"
)

// createTestPDF creates a simple multi-page PDF in the given directory and returns its path.
func createTestPDF(t *testing.T, dir string, name string, pages int) string {
	t.Helper()

	p := fpdf.New("P", "mm", "A4", "")
	for i := 1; i <= pages; i++ {
		p.AddPage()
		p.SetFont("Arial", "", 12)
		p.CellFormat(0, 10, "Page content for page "+string(rune('0'+i)), "", 1, "", false, 0, "")
	}

	path := filepath.Join(dir, name)
	if err := p.OutputFileAndClose(path); err != nil {
		t.Fatalf("failed to create test PDF %q: %v", name, err)
	}

	return path
}

// createTestPDFWithText creates a PDF with specific text content for text extraction testing.
func createTestPDFWithText(t *testing.T, dir string, name string, text string) string {
	t.Helper()

	p := fpdf.New("P", "mm", "A4", "")
	p.AddPage()
	p.SetFont("Arial", "", 12)
	p.MultiCell(0, 10, text, "", "", false)

	path := filepath.Join(dir, name)
	if err := p.OutputFileAndClose(path); err != nil {
		t.Fatalf("failed to create test PDF %q: %v", name, err)
	}

	return path
}

func TestConvertHTMLToPDF(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name       string
		html       string
		outputName string
		wantErr    bool
	}{
		{
			name:       "basic_paragraph",
			html:       "<p>Hello, World!</p>",
			outputName: "basic.pdf",
			wantErr:    false,
		},
		{
			name:       "headings",
			html:       "<h1>Title</h1><h2>Subtitle</h2><p>Content</p>",
			outputName: "headings.pdf",
			wantErr:    false,
		},
		{
			name:       "bold_italic",
			html:       "<p><b>Bold</b> and <i>Italic</i> text</p>",
			outputName: "bold_italic.pdf",
			wantErr:    false,
		},
		{
			name:       "strong_em",
			html:       "<p><strong>Strong</strong> and <em>Emphasis</em></p>",
			outputName: "strong_em.pdf",
			wantErr:    false,
		},
		{
			name:       "line_break",
			html:       "<p>Line one<br>Line two</p>",
			outputName: "linebreak.pdf",
			wantErr:    false,
		},
		{
			name:       "link",
			html:       `<p>Visit <a href="https://example.com">example</a></p>`,
			outputName: "link.pdf",
			wantErr:    false,
		},
		{
			name:       "complex_html",
			html:       "<html><body><h1>Title</h1><p>A <b>bold</b> and <i>italic</i> paragraph.</p><h2>Section</h2><p>More content here.</p></body></html>",
			outputName: "complex.pdf",
			wantErr:    false,
		},
		{
			name:       "empty_html",
			html:       "",
			outputName: "empty.pdf",
			wantErr:    true,
		},
		{
			name:       "whitespace_only",
			html:       "   \n\t  ",
			outputName: "whitespace.pdf",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputPath := filepath.Join(tmpDir, tt.outputName)
			err := ConvertHTMLToPDF(tt.html, outputPath)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ConvertHTMLToPDF() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				info, err := os.Stat(outputPath)
				if err != nil {
					t.Fatalf("output file not created: %v", err)
				}
				if info.Size() == 0 {
					t.Error("output file is empty")
				}
			}
		})
	}
}

func TestExtractTextFromPDF(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		setupFunc   func(t *testing.T) string
		wantContain string
		wantErr     bool
	}{
		{
			name: "extract_simple_text",
			setupFunc: func(t *testing.T) string {
				return createTestPDFWithText(t, tmpDir, "simple.pdf", "Hello World")
			},
			wantContain: "Hello World",
			wantErr:     false,
		},
		{
			name: "extract_multiline_text",
			setupFunc: func(t *testing.T) string {
				return createTestPDFWithText(t, tmpDir, "multiline.pdf", "Line one\nLine two\nLine three")
			},
			wantContain: "Line one",
			wantErr:     false,
		},
		{
			name: "nonexistent_file",
			setupFunc: func(_ *testing.T) string {
				return filepath.Join(tmpDir, "nonexistent.pdf")
			},
			wantContain: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPath := tt.setupFunc(t)
			text, err := ExtractTextFromPDF(inputPath)

			if (err != nil) != tt.wantErr {
				t.Fatalf("ExtractTextFromPDF() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr && tt.wantContain != "" {
				if !strings.Contains(text, tt.wantContain) {
					t.Errorf("ExtractTextFromPDF() text = %q, want to contain %q", text, tt.wantContain)
				}
			}
		})
	}
}

func TestMergePDFs(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) []string
		outputName string
		wantErr    bool
	}{
		{
			name: "merge_two_pdfs",
			setupFunc: func(t *testing.T) []string {
				return []string{
					createTestPDF(t, tmpDir, "merge_a1.pdf", 1),
					createTestPDF(t, tmpDir, "merge_a2.pdf", 1),
				}
			},
			outputName: "merged_two.pdf",
			wantErr:    false,
		},
		{
			name: "merge_three_pdfs",
			setupFunc: func(t *testing.T) []string {
				return []string{
					createTestPDF(t, tmpDir, "merge_b1.pdf", 2),
					createTestPDF(t, tmpDir, "merge_b2.pdf", 1),
					createTestPDF(t, tmpDir, "merge_b3.pdf", 3),
				}
			},
			outputName: "merged_three.pdf",
			wantErr:    false,
		},
		{
			name: "merge_single_file",
			setupFunc: func(t *testing.T) []string {
				return []string{
					createTestPDF(t, tmpDir, "merge_single.pdf", 1),
				}
			},
			outputName: "merged_single.pdf",
			wantErr:    true,
		},
		{
			name: "merge_empty_list",
			setupFunc: func(_ *testing.T) []string {
				return []string{}
			},
			outputName: "merged_empty.pdf",
			wantErr:    true,
		},
		{
			name: "merge_nonexistent_file",
			setupFunc: func(t *testing.T) []string {
				return []string{
					createTestPDF(t, tmpDir, "merge_exists.pdf", 1),
					filepath.Join(tmpDir, "nonexistent.pdf"),
				}
			},
			outputName: "merged_nonexistent.pdf",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFiles := tt.setupFunc(t)
			outputPath := filepath.Join(tmpDir, tt.outputName)

			err := MergePDFs(inputFiles, outputPath)

			if (err != nil) != tt.wantErr {
				t.Fatalf("MergePDFs() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				info, err := os.Stat(outputPath)
				if err != nil {
					t.Fatalf("output file not created: %v", err)
				}
				if info.Size() == 0 {
					t.Error("merged output is empty")
				}
			}
		})
	}
}

func TestSplitPDF(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name       string
		setupFunc  func(t *testing.T) string
		pageRanges []string
		wantFiles  []string
		wantErr    bool
	}{
		{
			name: "split_single_range",
			setupFunc: func(t *testing.T) string {
				return createTestPDF(t, tmpDir, "split_a.pdf", 5)
			},
			pageRanges: []string{"1-3"},
			wantFiles:  []string{"pages_1-3.pdf"},
			wantErr:    false,
		},
		{
			name: "split_multiple_ranges",
			setupFunc: func(t *testing.T) string {
				return createTestPDF(t, tmpDir, "split_b.pdf", 5)
			},
			pageRanges: []string{"1-2", "4-5"},
			wantFiles:  []string{"pages_1-2.pdf", "pages_4-5.pdf"},
			wantErr:    false,
		},
		{
			name: "split_single_page",
			setupFunc: func(t *testing.T) string {
				return createTestPDF(t, tmpDir, "split_c.pdf", 3)
			},
			pageRanges: []string{"2"},
			wantFiles:  []string{"pages_2-2.pdf"},
			wantErr:    false,
		},
		{
			name: "split_empty_ranges",
			setupFunc: func(t *testing.T) string {
				return createTestPDF(t, tmpDir, "split_d.pdf", 3)
			},
			pageRanges: []string{},
			wantFiles:  nil,
			wantErr:    true,
		},
		{
			name: "split_invalid_range",
			setupFunc: func(t *testing.T) string {
				return createTestPDF(t, tmpDir, "split_e.pdf", 3)
			},
			pageRanges: []string{"abc"},
			wantFiles:  nil,
			wantErr:    true,
		},
		{
			name: "split_reversed_range",
			setupFunc: func(t *testing.T) string {
				return createTestPDF(t, tmpDir, "split_f.pdf", 5)
			},
			pageRanges: []string{"5-3"},
			wantFiles:  nil,
			wantErr:    true,
		},
		{
			name: "split_nonexistent_file",
			setupFunc: func(_ *testing.T) string {
				return filepath.Join(tmpDir, "nonexistent.pdf")
			},
			pageRanges: []string{"1-2"},
			wantFiles:  nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile := tt.setupFunc(t)
			outputDir := filepath.Join(tmpDir, "split_out_"+tt.name)

			err := SplitPDF(inputFile, tt.pageRanges, outputDir)

			if (err != nil) != tt.wantErr {
				t.Fatalf("SplitPDF() error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				for _, expectedFile := range tt.wantFiles {
					path := filepath.Join(outputDir, expectedFile)
					info, err := os.Stat(path)
					if err != nil {
						t.Errorf("expected output file %q not found: %v", expectedFile, err)
						continue
					}
					if info.Size() == 0 {
						t.Errorf("output file %q is empty", expectedFile)
					}
				}
			}
		})
	}
}

func TestParsePageRange(t *testing.T) {
	tests := []struct {
		name      string
		rangeStr  string
		wantStart int
		wantEnd   int
		wantErr   bool
	}{
		{"valid_range", "1-5", 1, 5, false},
		{"single_page", "3", 3, 3, false},
		{"same_start_end", "4-4", 4, 4, false},
		{"large_range", "1-100", 1, 100, false},
		{"with_spaces", " 2 - 8 ", 2, 8, false},
		{"empty_string", "", 0, 0, true},
		{"invalid_start", "abc-5", 0, 0, true},
		{"invalid_end", "1-xyz", 0, 0, true},
		{"reversed_range", "5-1", 0, 0, true},
		{"zero_start", "0-5", 0, 0, true},
		{"negative_start", "-1-5", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, err := parsePageRange(tt.rangeStr)

			if (err != nil) != tt.wantErr {
				t.Fatalf("parsePageRange(%q) error = %v, wantErr = %v", tt.rangeStr, err, tt.wantErr)
			}

			if !tt.wantErr {
				if start != tt.wantStart {
					t.Errorf("parsePageRange(%q) start = %d, want %d", tt.rangeStr, start, tt.wantStart)
				}
				if end != tt.wantEnd {
					t.Errorf("parsePageRange(%q) end = %d, want %d", tt.rangeStr, end, tt.wantEnd)
				}
			}
		})
	}
}
