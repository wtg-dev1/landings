// Command export pre-renders every page in app.Sites to a directory tree
// suitable for `aws s3 sync` to a bucket fronted by CloudFront. Output:
//
//	<out>/<city>/<route>/index.html   — one per SiteEntry
//	<out>/404.html                    — for CloudFront custom_error_response
//	<out>/static/...                  — copied verbatim from --static-dir
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"wtg/landings/app"
)

func main() {
	dataDir := flag.String("data-dir", "data", "Path to data/ (JSON content per city/variant).")
	viewsDir := flag.String("views-dir", "views", "Path to views/ (layouts + partials).")
	staticDir := flag.String("static-dir", "static", "Path to static/ (css, images, fonts) — copied verbatim into <out>/static/.")
	outDir := flag.String("out-dir", "dist", "Output directory. Wiped of stale files only via the deployer; this CLI just writes into it.")
	flag.Parse()

	absViews, err := filepath.Abs(*viewsDir)
	if err != nil {
		log.Fatalf("resolve views-dir: %v", err)
	}
	absData, err := filepath.Abs(*dataDir)
	if err != nil {
		log.Fatalf("resolve data-dir: %v", err)
	}
	app.SetPaths(absViews, absData)

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		log.Fatalf("mkdir out: %v", err)
	}

	for _, s := range app.Sites {
		if err := renderOne(s, *outDir); err != nil {
			log.Fatalf("render %s/%s: %v", s.City, s.Route, err)
		}
	}

	if err := write404(*outDir); err != nil {
		log.Fatalf("write 404: %v", err)
	}

	if err := copyTree(*staticDir, filepath.Join(*outDir, "static")); err != nil {
		log.Fatalf("copy static: %v", err)
	}

	fmt.Printf("rendered %d pages + 404 + static into %s\n", len(app.Sites), *outDir)
}

func renderOne(s app.SiteEntry, outDir string) error {
	content, err := app.LoadPageContent(s.City, s.Variant)
	if err != nil {
		return err
	}
	content.FormName = app.GetFormName("/" + s.City + "/" + s.Route)
	dir := filepath.Join(outDir, s.City, s.Route)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, "index.html"))
	if err != nil {
		return err
	}
	defer f.Close()
	return app.RenderPage(content, f)
}

// write404 emits a minimal 404 page so CloudFront's custom_error_response
// has something to serve. Intentionally not template-driven — keeping it
// independent of PageContent shape avoids coupling the error page to
// per-variant data.
func write404(outDir string) error {
	const body = `<!DOCTYPE html>
<html lang="en"><head><meta charset="UTF-8"><title>Not found</title>
<link rel="stylesheet" href="/static/css/main.css"></head>
<body><main style="padding:4rem 1rem;text-align:center">
<h1>Page not found</h1><p><a href="/brooklyn/therapy">Williamsburg Therapy Group</a></p>
</main></body></html>`
	return os.WriteFile(filepath.Join(outDir, "404.html"), []byte(body), 0o644)
}

func copyTree(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		return copyFile(path, target)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
