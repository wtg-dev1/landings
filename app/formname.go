package app

import "strings"

// GetFormName mirrors the marketing site's JS getFormName() helper:
//
//	function getFormName() {
//	  const path = location.pathname;
//	  if (path.includes('/brooklyn/therapy')) return 'Brooklyn Therapy Form';
//	  const slug = path.split('/').filter(Boolean).pop() || 'home';
//	  return slug.replace(/-/g, ' ').replace(/\b\w/g, c => c.toUpperCase()) + ' Form';
//	}
//
// Computing this server-side and injecting the value into the rendered
// HTML keeps the dataLayer push deterministic per page — the browser
// doesn't have to derive it from the URL, and the value matches across
// SSR and CDN-cached responses.
func GetFormName(path string) string {
	if strings.Contains(path, "/brooklyn/therapy") {
		return "Brooklyn Therapy Form"
	}
	slug := lastNonEmptySegment(path)
	if slug == "" {
		slug = "home"
	}
	return titleCaseWords(strings.ReplaceAll(slug, "-", " ")) + " Form"
}

func lastNonEmptySegment(path string) string {
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] != "" {
			return parts[i]
		}
	}
	return ""
}

// titleCaseWords upper-cases the first byte of each space-separated word.
// Slugs are ASCII lowercase, so a byte-level upcase is correct here — no
// need to pull in golang.org/x/text/cases for this scope.
func titleCaseWords(s string) string {
	words := strings.Split(s, " ")
	for i, w := range words {
		if w == "" {
			continue
		}
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, " ")
}
