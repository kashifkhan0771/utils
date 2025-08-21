/*
Package url defines url utilities helpers.
*/
package url

import (
	"testing"
)

func TestBuildURL(t *testing.T) {
	t.Parallel()

	type args struct {
		scheme      string
		host        string
		path        string
		queryParams map[string]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - simple URL with single query param",
			args: args{scheme: "https", host: "example.com", path: "onePath", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne"}},
			want: "https://example.com/onePath?queryParamOne=valueQueryParamOne",
		},
		{
			name: "success - URL with multiple path segments and query params",
			args: args{scheme: "https", host: "example.com", path: "onePath/otherPath/other", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "https://example.com/onePath/otherPath/other?queryParamOne=valueQueryParamOne&queryParamTwo=valueQueryParamTwo",
		},
		{
			name: "success - subdomain URL with multiple query params",
			args: args{scheme: "https", host: "subdomain.example.com", path: "onePath", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "https://subdomain.example.com/onePath?queryParamOne=valueQueryParamOne&queryParamTwo=valueQueryParamTwo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if value, _ := BuildURL(tt.args.scheme, tt.args.host, tt.args.path, tt.args.queryParams); value != tt.want {
				t.Errorf("BuildURL() = got %v, want %v", value, tt.want)
			}
		})
	}
}

func TestBuildURLError(t *testing.T) {
	t.Parallel()

	type args struct {
		scheme      string
		host        string
		path        string
		queryParams map[string]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "error - build URL",
			args: args{scheme: "", host: "example.com", path: "onePath", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne"}},
			want: "scheme is required",
		},
		{
			name: "error - build URL",
			args: args{scheme: "https", host: "", path: "onePath", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "host is required",
		},
		{
			name: "error - build URL",
			args: args{scheme: "https", host: "example.com", path: "one1Path2@", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "path is permitted with a-z, 0-9, - and _ characters and multiple path segments",
		},
		{
			name: "error - build URL",
			args: args{scheme: "", host: "", path: "onePath", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "scheme is required; host is required",
		},
		{
			name: "error - build URL",
			args: args{scheme: "https", host: "ex@ample.com", path: "onePath", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "the host is not valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if _, err := BuildURL(tt.args.scheme, tt.args.host, tt.args.path, tt.args.queryParams); err.Error() != tt.want {
				t.Errorf("BuildURL() = got %v, want %v", err.Error(), tt.want)
			}
		})
	}
}

func TestAddQueryParams(t *testing.T) {
	t.Parallel()

	type args struct {
		urlStr      string
		queryParams map[string]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - add query params",
			args: args{urlStr: "https://example.com", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne"}},
			want: "https://example.com?queryParamOne=valueQueryParamOne",
		},
		{
			name: "success - add query params",
			args: args{urlStr: "https://subdomain.example.com", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "https://subdomain.example.com?queryParamOne=valueQueryParamOne&queryParamTwo=valueQueryParamTwo",
		},
		{
			name: "success - add query params",
			args: args{urlStr: "https://subdomain.example.com?firstQueryParam=anyValidValue", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: "https://subdomain.example.com?firstQueryParam=anyValidValue&queryParamOne=valueQueryParamOne&queryParamTwo=valueQueryParamTwo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if valid, _ := AddQueryParams(tt.args.urlStr, tt.args.queryParams); valid != tt.want {
				t.Errorf("AddQueryParams() = got %v, want %v", valid, tt.want)
			}
		})
	}
}

func TestAddQueryParamsError(t *testing.T) {
	t.Parallel()

	type args struct {
		urlStr      string
		queryParams map[string]string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "error - add query params",
			args: args{urlStr: "htt@p://example.com", queryParams: map[string]string{"queryParamOne": "valueQueryParamOne"}},
			want: "URL htt@p://example.com could not be parsed. err: parse \"htt@p://example.com\": first path segment in URL cannot contain colon",
		},
		{
			name: "error - add query params",
			args: args{urlStr: "https://subdomain.example.com", queryParams: map[string]string{"queryParam@One": "valueQueryParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: `query parameter key "queryParam@One" must contain only letters, digits, or hyphens`,
		},
		{
			name: "error - add query params",
			args: args{urlStr: "https://subdomain.example.com", queryParams: map[string]string{"queryParamOne": "valueQuery@ParamOne", "queryParamTwo": "valueQueryParamTwo"}},
			want: `query parameter value "valueQuery@ParamOne" for key "queryParamOne" must contain only letters, digits, or hyphens`,
		},
		{
			name: "error - add query params",
			args: args{urlStr: "https://subdomain.example.com", queryParams: map[string]string{"queryParamOne": ""}},
			want: "query parameter value for key queryParamOne cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if _, err := AddQueryParams(tt.args.urlStr, tt.args.queryParams); err.Error() != tt.want {
				t.Errorf("AddQueryParams() = got %v, want %v", err.Error(), tt.want)
			}
		})
	}
}

func TestIsValidURL(t *testing.T) {
	t.Parallel()

	type args struct {
		urlStr      string
		validScheme []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - is valid URL",
			args: args{urlStr: "https://example.com?someQuery=oneValue&otherQuery=otherValue", validScheme: []string{"http"}},
			want: false,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "https://subdomain.example.com?someQuery=oneValue&otherQuery=otherValue", validScheme: []string{"http", "https"}},
			want: true,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "ws://subdomain.example.com?someQuery=oneValue&otherQuery=otherValue", validScheme: []string{"http", "https", "ftp", "ws", "wss"}},
			want: true,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "ftp://subdomain.example.com?someQuery=oneValue&otherQuery=otherValue", validScheme: []string{"http", "https", "ftp", "ws", "wss"}},
			want: true,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "wss://subdomain.example.com?someQuery=oneValue&otherQuery=otherValue", validScheme: []string{"http", "https", "ftp", "ws", "wss"}},
			want: true,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "", validScheme: []string{"http", "https"}},
			want: false,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "example.com", validScheme: []string{"http", "https"}},
			want: false,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "example.com", validScheme: []string{""}},
			want: false,
		},
		{
			name: "success - is valid URL",
			args: args{urlStr: "exam@ple.com", validScheme: []string{"http", "https"}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if valid := IsValidURL(tt.args.urlStr, tt.args.validScheme); valid != tt.want {
				t.Errorf("IsValidURL() = got %v, want %v", valid, tt.want)
			}
		})
	}
}

func TestExtractDomain(t *testing.T) {
	t.Parallel()

	type args struct {
		urlStr string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - extract domain",
			args: args{urlStr: "https://example.com?someQuery=oneValue&otherQuery=otherValue"},
			want: "example.com",
		},
		{
			name: "success - domain with multiple subdomains",
			args: args{urlStr: "https://a.b.c.example.com"},
			want: "example.com",
		},
		{
			name: "success - domain with port",
			args: args{urlStr: "https://example.com:8080"},
			want: "example.com",
		},
		{
			name: "success - international domain",
			args: args{urlStr: "https://münchen.de"},
			want: "münchen.de",
		},
		{
			name: "error - invalid URL",
			args: args{urlStr: "not-a-url"},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got, _ := ExtractDomain(tt.args.urlStr); got != tt.want {
				t.Errorf("ExtractDomain() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestExtractDomainError(t *testing.T) {
	t.Parallel()

	type args struct {
		urlStr string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success - domain with multiple subdomains",
			args: args{urlStr: "htt@ps://exam@ple.com"},
			want: "URL htt@ps://exam@ple.com could not be parsed. err: parse \"htt@ps://exam@ple.com\": first path segment in URL cannot contain colon",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if _, err := ExtractDomain(tt.args.urlStr); err.Error() != tt.want {
				t.Errorf("ExtractDomain() = %s, want %s", err.Error(), tt.want)
			}
		})
	}
}

func TestGetQueryParam(t *testing.T) {
	t.Parallel()

	type args struct {
		urlStr string
		param  string
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "success - get query param",
			args: args{urlStr: "https://someurl.com?paramOne=oneValue&paramTwo=otherValue", param: "paramOne"},
			want: "oneValue",
		},
		{
			name: "success - get query param",
			args: args{urlStr: "https://someurl.com?paramOne=oneValue&paramTwo=otherValue&oneQuery=value&otherQuery=otherValue", param: "paramTwo"},
			want: "otherValue",
		},
		{
			name: "success - get query param",
			args: args{urlStr: "https://someurl.com?paramOne=oneValue&paramTwo=otherValue&oneQuery=valueOneQuery&otherQuery=otherValue", param: "oneQuery"},
			want: "valueOneQuery",
		},
		{
			name: "success - get query param",
			args: args{urlStr: "https://someurl.com?paramOne=oneValue&paramTwo=otherValue&oneQuery=valueOneQuery&otherQuery=otherQueryValue", param: "otherQuery"},
			want: "otherQueryValue",
		},
		{
			name: "success - simple parameter",
			args: args{
				urlStr: "https://example.com?key=value",
				param:  "key",
			},
			want: "value",
		},
		{
			name: "success - encoded parameter",
			args: args{
				urlStr: "https://example.com?key=value+with+spaces%26special%3Dchars",
				param:  "key",
			},
			want: "value with spaces&special=chars",
		},
		{
			name: "success - empty parameter value",
			args: args{
				urlStr: "https://example.com?key=",
				param:  "key",
			},
			want: "",
		},
		{
			name: "error - parameter not found",
			args: args{
				urlStr: "https://example.com?key=value",
				param:  "missing",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got, _ := GetQueryParam(tt.args.urlStr, tt.args.param); got != tt.want {
				t.Errorf("GetQueryParam() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestGetQueryParamError(t *testing.T) {
	t.Parallel()

	var (
		inputURL   = "https://someurl.com?paramOne=oneValue&paramTwo=otherValue"
		inputParam = "none"

		wantErrMsg = "parameter none not found in URL https://someurl.com?paramOne=oneValue&paramTwo=otherValue"
	)

	if _, err := GetQueryParam(inputURL, inputParam); err == nil {
		t.Errorf("GetQueryParam() expected error, got nil")
	} else if err.Error() != wantErrMsg {
		t.Errorf("GetQueryParam() error = %q, want %q", err.Error(), wantErrMsg)
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkBuildURL(b *testing.B) {
	var (
		scheme, host, path = "http", "example.com", "onePath"
		qp                 = map[string]string{"queryParamOne": "valueQueryParamOne"}
	)

	b.ReportAllocs()
	for b.Loop() {
		_, _ = BuildURL(scheme, host, path, qp)
	}
}

func BenchmarkAddQueryParams(b *testing.B) {
	var (
		urlStr = "https://example.com"
		qp     = map[string]string{"queryParamOne": "valueQueryParamOne"}
	)

	b.ReportAllocs()
	for b.Loop() {
		_, _ = AddQueryParams(urlStr, qp)
	}
}

func BenchmarkIsValidURL(b *testing.B) {
	var (
		urlStr  = "https://example.com"
		schemes = []string{"http", "https"}
	)

	b.ReportAllocs()
	for b.Loop() {
		IsValidURL(urlStr, schemes)
	}
}

func BenchmarkExtractDomain(b *testing.B) {
	urlStr := "https://example.com"

	b.ReportAllocs()
	for b.Loop() {
		_, _ = ExtractDomain(urlStr)
	}
}

func BenchmarkGetQueryParam(b *testing.B) {
	urlStr, key := "https://example.com?key=value", "key"

	b.ReportAllocs()
	for b.Loop() {
		_, _ = GetQueryParam(urlStr, key)
	}
}
