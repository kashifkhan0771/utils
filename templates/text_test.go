package templates

import (
	"testing"
	"time"
)

var (
	testTemplate1 = `Welcome {{toUpper .Name}}! Today is {{formatDate .Date "2006-01-02"}}.`
	want1         = "Welcome ALICE! Today is 2024-10-28."

	testTemplate2 = `{{title "this"}} {{toLower "IS"}} an {{reverse "elpmaxe"}} text {{toUpper "template"}} which uses some {{split "avail-able" "-"}} custom string funcs for testing`
	want2         = "This is an example text TEMPLATE which uses some [avail able] custom string funcs for testing"
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
					Date: time.Date(2024, time.October, 28, 0, 0, 0, 0, time.UTC),
				},
			},
			want:    want1,
			wantErr: false,
		},
		{
			name: "success - strings funcs",
			args: args{
				tmpl: testTemplate2,
				data: nil,
			},
			want:    want2,
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
