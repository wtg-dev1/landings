package app

// SiteEntry maps one rendered page: a city, the JSON variant filename it
// loads from data/<city>/<variant>.json, and the URL segment it serves at
// (/<city>/<route>). Route mirrors the PPC copy deck's "Variant URL slug"
// (e.g. /lp/couples-therapy) so ad copy and quality-score signals line up.
// The "general" master variant is the exception: served at /brooklyn/therapy.
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
	{City: "brooklyn", Variant: "couples", Route: "couples-therapy"},
	{City: "brooklyn", Variant: "family", Route: "family-therapy"},
	{City: "brooklyn", Variant: "child", Route: "child-therapy"},
	{City: "brooklyn", Variant: "marriage-counseling", Route: "marriage-counseling"},
	{City: "brooklyn", Variant: "anxiety-depression", Route: "anxiety-depression-therapy"},
	{City: "brooklyn", Variant: "cbt", Route: "cbt-therapy"},
	{City: "brooklyn", Variant: "lgbtq", Route: "lgbtq-therapy"},
	{City: "brooklyn", Variant: "addiction", Route: "addiction-therapy"},

	// {City: "miami", Variant: "general", Route: "therapy"},
	{City: "miami", Variant: "couples", Route: "couples-therapy"},
	{City: "miami", Variant: "family", Route: "family-therapy"},
	{City: "miami", Variant: "child", Route: "child-therapy"},
	{City: "miami", Variant: "marriage-counseling", Route: "marriage-counseling"},
	{City: "miami", Variant: "anxiety-depression", Route: "anxiety-depression-therapy"},
	{City: "miami", Variant: "cbt", Route: "cbt-therapy"},
	{City: "miami", Variant: "lgbtq", Route: "lgbtq-therapy"},
	{City: "miami", Variant: "addiction", Route: "addiction-therapy"},

	// {City: "austin", Variant: "general", Route: "therapy"},
	{City: "austin", Variant: "couples", Route: "couples-therapy"},
	{City: "austin", Variant: "family", Route: "family-therapy"},
	{City: "austin", Variant: "child", Route: "child-therapy"},
	{City: "austin", Variant: "marriage-counseling", Route: "marriage-counseling"},
	{City: "austin", Variant: "anxiety-depression", Route: "anxiety-depression-therapy"},
	{City: "austin", Variant: "cbt", Route: "cbt-therapy"},
	{City: "austin", Variant: "lgbtq", Route: "lgbtq-therapy"},
	{City: "austin", Variant: "addiction", Route: "addiction-therapy"},
}
