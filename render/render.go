package render

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Render struct {
	Renderer   string
	RootPath   string
	Source     bool
	Port       string
	ServerName string
}

type TemplateData struct {
	IsAuthenticted bool
	Data           map[string]interface{}
	CSRFToken      string
	Port           string
	ServerName     string
	Secure         bool
}

func (n *Render) Page(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {

	switch strings.ToLower(n.Renderer) {
	case "go":
		return n.GoPage(w, r, view, data)
	case "jet":
		break
	}

	return nil
}

func (n *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/view/%s.page.tmpl", n.RootPath, view))
	if err != nil {
		return err
	}

	td := TemplateData{}
	if data != nil {
		td = data.(TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}
