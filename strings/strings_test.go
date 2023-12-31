package strings

import (
	"reflect"
	"testing"
)

func TestTitleCase(t *testing.T) {
	type args struct {
		input      string
		exceptions []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - convert a fully lower case string to title case without exceptions",
			args: args{input: "lower case", exceptions: []string{}},
			want: "Lower Case",
		},
		{
			name: "success - convert a fully lower case string to title case with exceptions",
			args: args{input: "lower case", exceptions: []string{"case"}},
			want: "Lower case",
		},
		{
			name: "success - convert a camel lower case string to title case without exceptions",
			args: args{input: "lower Case", exceptions: []string{}},
			want: "Lower Case",
		},
		{
			name: "success - convert a camel lower case string to title case with exceptions",
			args: args{input: "lower Case to camel CASE", exceptions: []string{"CASE"}},
			want: "Lower Case To Camel CASE",
		},
		{
			name: "success - convert a upper case string to title case without exceptions",
			args: args{input: "UPPER CASE", exceptions: []string{}},
			want: "Upper Case",
		},
		{
			name: "success - convert a upper case string to title case with exceptions",
			args: args{input: "UPPER CASE WITH exception", exceptions: []string{"exception"}},
			want: "Upper Case With exception",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTitle(tt.args.input, tt.args.exceptions); got != tt.want {
				t.Errorf("ToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenize(t *testing.T) {
	type args struct {
		input            string
		customDelimiters string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success - tokenize a string",
			args: args{input: "This is a custom-tokenization!example", customDelimiters: "-!"},
			want: []string{"This", "is", "a", "custom", "tokenization", "example"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tokenize(tt.args.input, tt.args.customDelimiters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstringSearch(t *testing.T) {
	type args struct {
		input     string
		substring string
		options   SubstringSearchOptions
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success - search a substring without case insensitivity and indexes",
			args: args{input: "find this in a string", substring: "this", options: SubstringSearchOptions{
				CaseInsensitive: false,
				ReturnIndexes:   false,
			}},
			want: []string{"this"},
		},
		{
			name: "success - search a substring with case insensitivity and without indexes",
			args: args{input: "find THIS in a string", substring: "THIS", options: SubstringSearchOptions{
				CaseInsensitive: true,
				ReturnIndexes:   false,
			}},
			want: []string{"THIS"},
		},
		{
			name: "success - search a substring without case insensitivity and with indexes",
			args: args{input: "find this in a string", substring: "this", options: SubstringSearchOptions{
				CaseInsensitive: false,
				ReturnIndexes:   true,
			}},
			want: []string{"this in a string"},
		},
		{
			name: "success - search a multiple substring without case insensitivity and indexes",
			args: args{input: "find this in a string, and this again in a string", substring: "this", options: SubstringSearchOptions{
				CaseInsensitive: false,
				ReturnIndexes:   false,
			}},
			want: []string{"this", "this"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstringSearch(tt.args.input, tt.args.substring, tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubstringSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRot13Encode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - encode a string using Rot13 cipher",
			args: args{input: "Hello, World!"},
			want: "Uryyb, Jbeyq!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rot13Encode(tt.args.input); got != tt.want {
				t.Errorf("Rot13Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRot13Decode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - decode a string using Rot13 cipher",
			args: args{input: "Uryyb, Jbeyq!"},
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rot13Decode(tt.args.input); got != tt.want {
				t.Errorf("Rot13Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaesarEncrypt(t *testing.T) {
	type args struct {
		input string
		shift int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - encrypt a string using caesar cipher",
			args: args{input: "Hello, World!", shift: 3},
			want: "Khoor, Zruog!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaesarEncrypt(tt.args.input, tt.args.shift); got != tt.want {
				t.Errorf("CaesarEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaesarDecrypt(t *testing.T) {
	type args struct {
		input string
		shift int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - decrypt a string using caesar cipher",
			args: args{input: "Khoor, Zruog!", shift: 3},
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaesarDecrypt(tt.args.input, tt.args.shift); got != tt.want {
				t.Errorf("CaesarDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - valid email",
			args: args{email: "test-email@test.com"},
			want: true,
		},
		{
			name: "fail - invalid email",
			args: args{email: "test-email@test"},
			want: false,
		},
		{
			name: "fail - invalid email",
			args: args{email: "test-email#test.com"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSanitizeEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - sanitize an email",
			args: args{email: " test@test.com "},
			want: "test@test.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeEmail(tt.args.email); got != tt.want {
				t.Errorf("SanitizeEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - convert full uppercase to title case",
			args: args{input: "UPPERCASE"},
			want: "Uppercase",
		},
		{
			name: "success - convert full lowercase to title case",
			args: args{input: "lowercase"},
			want: "Lowercase",
		},
		{
			name: "success - empty string",
			args: args{input: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Title(tt.args.input); got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}
