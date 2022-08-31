package render

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gabrielyea/mini-booking-go/pkg/config"
	"github.com/gabrielyea/mini-booking-go/pkg/models"
)

var app *config.AppConfig

// NewTemplates set config for app config
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	var err error

	// get requested template
	t, ok := tc[tmpl]
	if !ok {
		fmt.Printf("err: %v\n", err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, td)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// render template
	_, err1 := buf.WriteTo(w)
	if err != nil {
		fmt.Printf("err: %v\n", err1)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files .page.html
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Printf("name: %v\n", name)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
