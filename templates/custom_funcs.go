package templates

import (
	"strings"
	"text/template"
	"time"

	strutils "github.com/kashifkhan0771/utils/strings"
)

// custom functions for templates
var customFuncsMap = template.FuncMap{
	// string function
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"title":    strutils.Title,
	"contains": strings.Contains,
	"replace":  strings.ReplaceAll,
	"trim":     strings.TrimSpace,
	"split":    strings.Split,
	"reverse":  strutils.Reverse,

	// date and time functions
	"formatDate": func(t time.Time, layout string) string { return t.Format(layout) },

	// numeric and arithmetic functions
	"add": func(a, b int) int { return a + b },
	"sub": func(a, b int) int { return a - b },
	"mul": func(a, b int) int { return a * b },
	"div": func(a, b int) int { return a / b },
	"mod": func(a, b int) int { return a % b },
}
