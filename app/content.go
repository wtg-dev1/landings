package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadPageContent reads <dataDir>/<city>/<variant>.json and decodes it
// into PageContent. dataDir defaults to "data" (relative to cwd) and can
// be overridden via SetPaths.
func LoadPageContent(city, variant string) (*PageContent, error) {
	path := filepath.Join(dataDir, city, variant+".json")

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var pc PageContent
	if err := json.Unmarshal(raw, &pc); err != nil {
		return nil, fmt.Errorf("decode %s: %w", path, err)
	}
	return &pc, nil
}
