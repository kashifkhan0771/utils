package templates

import (
	"html"
	htmlTemplate "html/template"
	"regexp"
	"strings"
	"testing"
	"time"
)

func normalizeWhitespace(input string) string {
	// Remove leading and trailing whitespace
	input = strings.TrimSpace(input)

	// Normalize whitespace by replacing multiple spaces and newlines with a single space
	whitespace := regexp.MustCompile(`\s+`)
	input = whitespace.ReplaceAllString(input, " ")

	// Remove spaces before and after HTML tags to ensure consistent formatting
	input = regexp.MustCompile(`\s*(<[^>]+>)\s*`).ReplaceAllString(input, "$1")

	return input
}

var (
	htmlTestTemplate1 = `<!DOCTYPE html><html><body>Welcome {{toUpper .Name}}! Today is {{formatDate .Date "2006-01-02"}}.</body></html>`
	htmlWant1         = `<!DOCTYPE html><html><body>Welcome ALICE! Today is 2024-10-01.</body></html>`

	htmlTestTemplate2 = `<!DOCTYPE html><html><body>
		{{- $upper := toUpper "hello" -}}
		{{- $lower := toLower "WORLD" -}}
		{{- $title := title "hello world" -}}
		{{- $contains := contains "hello world" "world" -}}
		{{- $replace := replace "go gophers" "go" "GoLang" -}}
		{{- $trim := trim "   trimmed   " -}}
		{{- $split := index (split "one,two,three" ",") 1 -}}
		{{- $reverse := reverse "abcde" -}}

		Uppercase: {{$upper}}<br>
		Lowercase: {{$lower}}<br>
		Title Case: {{$title}}<br>
		Contains 'world': {{$contains}}<br>
		Replace 'go' with 'GoLang': {{$replace}}<br>
		Trimmed: {{$trim}}<br>
		Split Result [1]: {{$split}}<br>
		Reversed: {{$reverse}}
	</body></html>`
	htmlWant2 = `<!DOCTYPE html><html><body>
		Uppercase: HELLO<br>
		Lowercase: world<br>
		Title Case: Hello World<br>
		Contains 'world': true<br>
		Replace 'go' with 'GoLang': GoLang GoLangphers<br>
		Trimmed: trimmed<br>
		Split Result [1]: two<br>
		Reversed: edcba
	</body></html>`

	htmlTestTemplate3 = `<!DOCTYPE html><html><body>
        {{- $sum := add 1 4 -}}
        {{- $sub := sub 4 1 -}}
	   {{- $mul := mul 2 2 -}}
	   {{- $div := div 2 2 -}}
	   {{- $mod := mod 3 2 -}}

     	Addition: {{$sum}}<br>
        	Subtraction: {{$sub}}<br>
	   	Multiplication: {{$mul}}<br>
	   	Division: {{$div}}<br>
	   	Mod: {{$mod}}
	</body></html>`
	htmlWant3 = `<!DOCTYPE html><html><body>
		Addition: 5<br>
        	Subtraction: 3<br>
	   	Multiplication: 4<br>
	   	Division: 1<br>
	   	Mod: 1
	</body></html>`

	htmlTestTemplate4 = `<!DOCTYPE html><html><body>
		{{- $isNil := isNil .NilValue -}}
		{{- $notNil := isNil .NotNilValue -}}
	     {{- $notTrue := not true -}}

	     Is Nil: {{$isNil}}<br>
		Is Nil: {{$notNil}}<br>
	     Not True: {{$notTrue}}
	</body></html>`
	htmlWant4 = `<!DOCTYPE html><html><body>
		Is Nil: true<br>
		Is Nil: false<br>
	     Not True: false
	</body></html>`

	htmlTestTemplate5 = `<!DOCTYPE html><html><body>
     	{{- $dumpValue := dump .SampleMap -}}
        	{{- $typeOfValue := typeOf .SampleMap -}}

     	Dump: {{$dumpValue}}<br>
     	Type Of: {{$typeOfValue}}
	</body></html>`
	htmlWant5 = `<!DOCTYPE html><html><body>
		Dump: map[string]int{"a":1, "b":2}<br>
     	Type Of: map[string]int
	</body></html>`
)

func TestRenderHTML(t *testing.T) {
	type args struct {
		tmpl string
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success - simple HTML",
			args: args{
				tmpl: htmlTestTemplate1,
				data: struct {
					Name string
					Date time.Time
				}{
					Name: "alice",
					Date: time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want:    htmlWant1,
			wantErr: false,
		},
		{
			name: "success - string funcs in HTML",
			args: args{
				tmpl: normalizeWhitespace(htmlTestTemplate2),
				data: nil,
			},
			want:    normalizeWhitespace(htmlWant2),
			wantErr: false,
		},
		{
			name: "success - numeric and arithmetic funcs in HTML",
			args: args{
				tmpl: normalizeWhitespace(htmlTestTemplate3),
				data: nil,
			},
			want:    normalizeWhitespace(htmlWant3),
			wantErr: false,
		},
		{
			name: "success - conditional and logical funcs in HTML",
			args: args{
				tmpl: normalizeWhitespace(htmlTestTemplate4),
				data: struct {
					NilValue    interface{}
					NotNilValue interface{}
				}{
					NilValue:    nil,
					NotNilValue: "example",
				},
			},
			want:    normalizeWhitespace(htmlWant4),
			wantErr: false,
		},
		{
			name: "success - debugging funcs in HTML",
			args: args{
				tmpl: normalizeWhitespace(htmlTestTemplate5),
				data: struct {
					SampleMap map[string]int
				}{
					SampleMap: map[string]int{"a": 1, "b": 2},
				},
			},
			want:    normalizeWhitespace(htmlWant5),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := htmlTemplate.New("htmlTestTemplate").Funcs(GetCustomFuncMap()).Parse(tt.args.tmpl)
			if err != nil {
				t.Fatalf("failed to parse template: %v", err)
			}
			var sb strings.Builder
			err = tmpl.Execute(&sb, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := sb.String()

			//Trim whitespace for comparison
			gotTrimmed := strings.TrimSpace(got)
			wantTrimmed := strings.TrimSpace(tt.want)

			// Handle HTML-escaped characters
			gotUnescaped := html.UnescapeString(gotTrimmed)
			wantUnescaped := html.UnescapeString(wantTrimmed)

			if gotUnescaped != wantUnescaped {
				t.Errorf("Execute() = %v, want %v", gotUnescaped, wantUnescaped)
			}
		})
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkRenderHTML(b *testing.B) {
	tmpl, _ := htmlTemplate.New("htmlTestTemplate").Funcs(GetCustomFuncMap()).Parse(normalizeWhitespace(htmlTestTemplate1))
	data := struct {
		Name string
		Date time.Time
	}{
		Name: "alice",
		Date: time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		_ = tmpl.Execute(&sb, data)
	}
}
