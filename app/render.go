package app

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"sync"
)

var (
	viewsDir = "views"
	dataDir  = "data"

	tmpl     *template.Template
	tmplOnce sync.Once
	tmplErr  error
)

// SetPaths overrides the default views/data directory locations. Call
// before the first RenderPage or LoadPageContent — the defaults are
// relative to cwd, which is what the gin server expects when run from
// src/. The export CLI passes absolute paths so it can be invoked from
// any working directory.
func SetPaths(views, data string) {
	viewsDir = views
	dataDir = data
}

func loadTemplates() {
	t := template.New("")
	patterns := []string{
		filepath.Join(viewsDir, "layouts", "*.html"),
		filepath.Join(viewsDir, "partials", "*.html"),
	}
	for _, p := range patterns {
		parsed, err := t.ParseGlob(p)
		if err != nil {
			tmplErr = fmt.Errorf("parse %s: %w", p, err)
			return
		}
		t = parsed
	}
	tmpl = t
}

// RenderPage executes the base template against the given content and
// writes the rendered HTML to w. Templates are parsed lazily on first
// call; both the gin server and the export CLI share this single path.
func RenderPage(pc *PageContent, w io.Writer) error {
	tmplOnce.Do(loadTemplates)
	if tmplErr != nil {
		return tmplErr
	}
	return tmpl.ExecuteTemplate(w, "base.html", pc)
}
