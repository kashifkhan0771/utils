package templates

import (
	"bytes"
	"html/template"
)

// RenderHTMLTemplate processes an HTML template with the provided data.
func RenderHTMLTemplate(tmpl string, data interface{}) (string, error) {
	t, err := template.New("htmlTemplate").Funcs(GetCustomFuncMap()).Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
