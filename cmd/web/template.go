package main

import (
	"io/fs"
	"path/filepath"
	"text/template"

	"github.com/Torkel-Aannestad/torkel-dev/ui"
)

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/**/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		pattern := []string{
			"html/layouts/*.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, pattern...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil

}
