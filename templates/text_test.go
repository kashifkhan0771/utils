package templates

import (
	"testing"
	"time"
)

var (
	testTemplate1 = `Welcome {{toUpper .Name}}! Today is {{formatDate .Date "2006-01-02"}}.`
	want1         = "Welcome ALICE! Today is 2024-10-01."

	testTemplate2 = `
		{{- $upper := toUpper "hello" -}}
		{{- $lower := toLower "WORLD" -}}
		{{- $title := title "hello world" -}}
		{{- $contains := contains "hello world" "world" -}}
		{{- $replace := replace "go gophers" "go" "GoLang" -}}
		{{- $trim := trim "   trimmed   " -}}
		{{- $split := index (split "one,two,three" ",") 1 -}}
		{{- $reverse := reverse "abcde" -}}

		Uppercase: {{$upper}}
		Lowercase: {{$lower}}
		Title Case: {{$title}}
		Contains 'world': {{$contains}}
		Replace 'go' with 'GoLang': {{$replace}}
		Trimmed: {{$trim}}
		Split Result [1]: {{$split}}
		Reversed: {{$reverse}}
	`
	want2 = `Uppercase: HELLO
		Lowercase: world
		Title Case: Hello World
		Contains 'world': true
		Replace 'go' with 'GoLang': GoLang GoLangphers
		Trimmed: trimmed
		Split Result [1]: two
		Reversed: edcba
	`

	testTemplate3 = `
        {{- $sum := add 1 4 -}}
        {{- $sub := sub 4 1 -}}
	   {{- $mul := mul 2 2 -}}
	   {{- $div := div 2 2 -}}
	   {{- $mod := mod 3 2 -}}

     	Addition: {{$sum}}
        	Subtraction: {{$sub}}
	   	Multiplication: {{$mul}}
	   	Division: {{$div}}
	   	Mod: {{$mod}}
	`
	want3 = `Addition: 5
        	Subtraction: 3
	   	Multiplication: 4
	   	Division: 1
	   	Mod: 1
	`

	testTemplate4 = `
		{{- $isNil := isNil .NilValue -}}
		{{- $notNil := isNil .NotNilValue -}}
	     {{- $notTrue := not true -}}

	     Is Nil: {{$isNil}}
		Is Nil: {{$notNil}}
	     Not True: {{$notTrue}}
	`
	want4 = `Is Nil: true
		Is Nil: false
	     Not True: false
	`

	testTemplate5 = `
     	{{- $dumpValue := dump .SampleMap -}}
        	{{- $typeOfValue := typeOf .SampleMap -}}

     	Dump: {{$dumpValue}}
     	Type Of: {{$typeOfValue}}
	`
	want5 = `Dump: map[string]int{"a":1, "b":2}
     	Type Of: map[string]int
	`
)

func TestRenderText(t *testing.T) {
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
			name: "success - simple text",
			args: args{
				tmpl: testTemplate1,
				data: struct {
					Name string
					Date time.Time
				}{
					Name: "alice",
					Date: time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want:    want1,
			wantErr: false,
		},
		{
			name: "success - string funcs",
			args: args{
				tmpl: testTemplate2,
				data: nil,
			},
			want:    want2,
			wantErr: false,
		},
		{
			name: "success - numeric and arithmetic funcs",
			args: args{
				tmpl: testTemplate3,
				data: nil,
			},
			want:    want3,
			wantErr: false,
		},
		{
			name: "success - conditional and logical funcs",
			args: args{
				tmpl: testTemplate4,
				data: struct {
					NilValue    interface{}
					NotNilValue interface{}
				}{
					NilValue:    nil,
					NotNilValue: "example",
				},
			},
			want:    want4,
			wantErr: false,
		},
		{
			name: "success - debugging funcs",
			args: args{
				tmpl: testTemplate5,
				data: struct {
					SampleMap map[string]int
				}{
					SampleMap: map[string]int{"a": 1, "b": 2},
				},
			},
			want:    want5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderText(tt.args.tmpl, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderText() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if got != tt.want {
				t.Errorf("RenderText() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkRenderText(b *testing.B) {
	data := struct {
		Name string
		Date time.Time
	}{
		Name: "alice",
		Date: time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC),
	}

	b.ReportAllocs()
	for b.Loop() {
		_, _ = RenderText(testTemplate1, data)
	}
}
