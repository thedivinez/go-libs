package gothex

import (
	"html/template"
	"io"
)

//go:embed error_page_template.html
var errorPageTemplate string

var tmpl = template.Must(template.New("errorpage").Parse(errorPageTemplate))

func RenderErrorPage(w io.Writer, page ErrorPageContent) error {
	return tmpl.Execute(w, page)
}
