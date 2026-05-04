package app

// SiteEntry maps one rendered page: a city, the JSON variant filename it
// loads from data/<city>/<variant>.json, and the URL segment it serves at
// (/<city>/<route>). For most pages Route == Variant; the exception is the
// "general" variant which serves at /brooklyn/therapy.
type SiteEntry struct {
	City    string
	Variant string
	Route   string
}

// Sites is the single source of truth for which pages exist. Both the gin
// server (server/routes/api.go) and the static export (cmd/export) iterate
// it. To add a new page: drop a JSON file under data/<city>/, append here.
var Sites = []SiteEntry{
	{City: "brooklyn", Variant: "general", Route: "therapy"},
	{City: "brooklyn", Variant: "couples", Route: "couples"},
	{City: "brooklyn", Variant: "family", Route: "family"},
	{City: "brooklyn", Variant: "anxiety-depression", Route: "anxiety-depression"},
	{City: "brooklyn", Variant: "cbt", Route: "cbt"},
	{City: "brooklyn", Variant: "lgbtq", Route: "lgbtq"},
	{City: "brooklyn", Variant: "addiction", Route: "addiction"},
}
