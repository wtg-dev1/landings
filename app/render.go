package app

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	viewsDir  = "views"
	dataDir   = "data"
	staticDir = "static"

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

// SetStaticDir overrides the default static directory. Used by the
// inlineCSS template helper so the base layout can inline main.css
// into the rendered HTML (eliminates a render-blocking request).
func SetStaticDir(s string) {
	staticDir = s
}

func loadTemplates() {
	funcMap := template.FuncMap{
		// inlineCSS reads a file from staticDir and returns it as
		// template.CSS so html/template skips escaping. Lets the
		// base layout drop /static/css/main.css inline and shave the
		// CSS roundtrip off LCP.
		"inlineCSS": func(relPath string) (template.CSS, error) {
			b, err := os.ReadFile(filepath.Join(staticDir, relPath))
			if err != nil {
				return "", err
			}
			return template.CSS(b), nil
		},
	}
	t := template.New("").Funcs(funcMap)
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
