package validation

import (
	"errors"
	"testing"
)

func TestStringNotEmpty(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "success - non-empty value",
			value: "jane",
			want:  "",
		},
		{
			name:  "failure - whitespace only value",
			value: "   ",
			want:  "name: must not be empty",
		},
		{
			name:  "failure - empty value",
			value: "",
			want:  "name: must not be empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StringNotEmpty("name", tt.value)
			got := ""
			if err != nil {
				got = err.Error()
			}
			if got != tt.want {
				t.Errorf("StringNotEmpty() error = %v, want %q", err, tt.want)
			}
		})
	}
}

func TestStringIn(t *testing.T) {
	t.Parallel()

	type args struct {
		value   string
		allowed []string
	}

	tests := []struct {
		name string
		arg  args
		want string
	}{
		{
			name: "success - value in allowed list",
			arg:  args{value: "admin", allowed: []string{"superuser", "admin"}},
			want: "",
		},
		{
			name: "failure - value not in allowed list",
			arg:  args{value: "hacker", allowed: []string{"superuser", "admin", "user"}},
			want: "role: must be one of [superuser, admin, user]",
		},
		{
			name: "success - no constraints when allowed is empty",
			arg:  args{value: "anything", allowed: []string{}},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StringIn("role", tt.arg.value, tt.arg.allowed...)
			got := ""
			if err != nil {
				got = err.Error()
			}
			if got != tt.want {
				t.Errorf("StringIn() error = %v, want %q", err, tt.want)
			}
		})
	}
}

func TestIntBetween(t *testing.T) {
	t.Parallel()

	type args struct {
		value int
		min   int
		max   int
	}

	tests := []struct {
		name string
		arg  args
		want string
	}{
		{
			name: "success - value within range",
			arg:  args{value: 23, min: 18, max: 67},
			want: "",
		},
		{
			name: "success - value exactly at min boundary",
			arg:  args{value: 18, min: 18, max: 67},
			want: "",
		},
		{
			name: "success - value exactly at max boundary",
			arg:  args{value: 67, min: 18, max: 67},
			want: "",
		},
		{
			name: "failure - value above max",
			arg:  args{value: 200, min: 18, max: 75},
			want: "age: integer must be between 18 and 75",
		},
		{
			name: "failure - value below min",
			arg:  args{value: 10, min: 18, max: 75},
			want: "age: integer must be between 18 and 75",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IntBetween("age", tt.arg.value, tt.arg.min, tt.arg.max)
			got := ""
			if err != nil {
				got = err.Error()
			}
			if got != tt.want {
				t.Errorf("IntBetween() error = %v, want %q", err, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "success - valid email",
			value: "jane@example.com",
			want:  "",
		},
		{
			name:  "failure - invalid email",
			value: "not-an-email",
			want:  "email: mail: missing '@' or angle-addr",
		},
		{
			name:  "failure - display name not allowed",
			value: "John <jane@example.com>",
			want:  "email: email must be valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsEmail("email", tt.value)
			got := ""
			if err != nil {
				got = err.Error()
			}
			if got != tt.want {
				t.Errorf("IsEmail() error = %v, want %q", err, tt.want)
			}
		})
	}
}

func TestIsURL(t *testing.T) {
	t.Parallel()

	type args struct {
		value          string
		allowedSchemes []string
	}

	tests := []struct {
		name string
		arg  args
		want string
	}{
		{
			name: "success - valid URL with default schemes",
			arg:  args{value: "https://example.com"},
			want: "",
		},
		{
			name: "success - custom scheme passed explicitly",
			arg:  args{value: "ftp://example.com", allowedSchemes: []string{"ftp"}},
			want: "",
		},
		{
			name: "failure - scheme not in default list",
			arg:  args{value: "ftp://example.com"},
			want: "url: must be a valid URL",
		},
		{
			name: "failure - scheme not in custom list",
			arg:  args{value: "ftp://example.com", allowedSchemes: []string{"https"}},
			want: "url: must be a valid URL",
		},
		{
			name: "failure - no scheme at all",
			arg:  args{value: "example.com"},
			want: "url: must be a valid URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsURL("url", tt.arg.value, tt.arg.allowedSchemes...)
			got := ""
			if err != nil {
				got = err.Error()
			}
			if got != tt.want {
				t.Errorf("IsURL() error = %v, want %q", err, tt.want)
			}
		})
	}
}

func TestIsSlug(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "success - valid slug",
			value: "my-new-post",
			want:  "",
		},
		{
			name:  "success - slug with digits",
			value: "post-42",
			want:  "",
		},
		{
			name:  "failure - uppercase characters",
			value: "My-New-Post",
			want:  "slug: must be a valid slug",
		},
		{
			name:  "failure - spaces",
			value: "my new post",
			want:  "slug: must be a valid slug",
		},
		{
			name:  "failure - underscores",
			value: "my_new_post",
			want:  "slug: must be a valid slug",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsSlug("slug", tt.value)
			got := ""
			if err != nil {
				got = err.Error()
			}
			if got != tt.want {
				t.Errorf("IsSlug() error = %v, want %q", err, tt.want)
			}
		})
	}
}

func TestErrors_Add(t *testing.T) {
	t.Parallel()

	t.Run("add single error", func(t *testing.T) {
		var errs Errors

		errs.Add(errors.New("first error"))

		if err := errs.AsError(); err == nil || err.Error() != "first error" {
			t.Errorf("AsError() = %v, want %q", err, "first error")
		}
	})

	t.Run("add multiple errors", func(t *testing.T) {
		var errs Errors

		errs.Add(errors.New("first error"))
		errs.Add(errors.New("second error"))
		errs.Add(errors.New("third error"))

		want := "first error; second error; third error"
		if err := errs.AsError(); err == nil || err.Error() != want {
			t.Errorf("AsError() = %v, want %q", err, want)
		}
	})

	t.Run("nil errors are ignored", func(t *testing.T) {
		var errs Errors

		errs.Add(nil)
		errs.Add(nil)

		if err := errs.AsError(); err != nil {
			t.Errorf("AsError() = %v, want nil", err)
		}
	})
}

func TestErrors_AsError(t *testing.T) {
	t.Parallel()

	t.Run("no errors added returns nil", func(t *testing.T) {
		var errs Errors

		if err := errs.AsError(); err != nil {
			t.Errorf("AsError() = %v, want nil", err)
		}
	})

	t.Run("errors joined with semicolon", func(t *testing.T) {
		var errs Errors

		errs.Add(errors.New("one"))
		errs.Add(errors.New("two"))
		errs.Add(errors.New("three"))

		want := "one; two; three"
		if err := errs.AsError(); err == nil || err.Error() != want {
			t.Errorf("AsError() = %v, want %q", err, want)
		}
	})
}

type testStruct struct {
	name string
}

func (ts testStruct) Validate() error {
	return StringNotEmpty("name", ts.name)
}

func TestValidateStruct(t *testing.T) {
	t.Parallel()

	t.Run("success - valid struct", func(t *testing.T) {
		if err := ValidateStruct(testStruct{name: "jane"}); err != nil {
			t.Errorf("ValidateStruct() = %v, want nil", err)
		}
	})

	t.Run("failure - invalid struct", func(t *testing.T) {
		want := "name: must not be empty"
		if err := ValidateStruct(testStruct{name: ""}); err == nil || err.Error() != want {
			t.Errorf("ValidateStruct() = %v, want %q", err, want)
		}
	})

	t.Run("nil validator returns nil", func(t *testing.T) {
		if err := ValidateStruct(nil); err != nil {
			t.Errorf("ValidateStruct() = %v, want nil", err)
		}
	})
}
