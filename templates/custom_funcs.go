package templates

import (
	"fmt"
	htmlTemplate "html/template"
	"strings"
	textTemplate "text/template"
	"time"

	strutils "github.com/kashifkhan0771/utils/strings"
)

// custom functions for templates
var customFuncsMap = textTemplate.FuncMap{
	// string functions
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"title":    strutils.Title,
	"contains": strings.Contains,
	"replace":  strings.ReplaceAll,
	"trim":     strings.TrimSpace,
	"split":    strings.Split,
	"reverse":  strutils.Reverse,
	"toString": func(v interface{}) string { return fmt.Sprintf("%v", v) },

	// date and time functions
	"formatDate": func(t time.Time, layout string) string { return t.Format(layout) },
	"now":        time.Now,

	// numeric and arithmetic functions
	"add": func(a, b int) int { return a + b },
	"sub": func(a, b int) int { return a - b },
	"mul": func(a, b int) int { return a * b },
	"div": func(a, b int) int { return a / b },
	"mod": func(a, b int) int { return a % b },

	// conditional and logical functions
	"isNil": func(v interface{}) bool { return v == nil },
	"not":   func(v bool) bool { return !v },

	// debugging functions
	"dump":   func(v interface{}) string { return fmt.Sprintf("%#v", v) },
	"typeOf": func(v interface{}) string { return fmt.Sprintf("%T", v) },

	// safe HTML rendering for trusted content	
	"safeHTML": func(s string) htmlTemplate.HTML {
		return htmlTemplate.HTML(s)
	},
}

// GetCustomFuncMap returns a map of custom functions
func GetCustomFuncMap() textTemplate.FuncMap {
	return customFuncsMap
}
