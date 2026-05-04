package app

// PageContent is the full payload backing a single landing page render.
// Each top-level section maps to a partial in views/partials.
type PageContent struct {
	City    string `json:"city"`
	Variant string `json:"variant"`

	Phone string `json:"phone"`

	Meta         Meta         `json:"meta"`
	Header       Header       `json:"header"`
	Anchors      []Anchor     `json:"anchors"`
	Hero         Hero         `json:"hero"`
	Stats        []Stat       `json:"stats"`
	Treat        Treat        `json:"treat"`
	Comparison   Comparison   `json:"comparison"`
	HowItWorks   HowItWorks   `json:"howItWorks"`
	WhyWTG       WhyWTG       `json:"whyWTG"`
	Clinicians   Clinicians   `json:"clinicians"`
	MidCTA       MidCTA       `json:"midCTA"`
	Pricing      Pricing      `json:"pricing"`
	Locations    Locations    `json:"locations"`
	Testimonials Testimonials `json:"testimonials"`
	FAQ          FAQ          `json:"faq"`
	Form         Form         `json:"form"`
	StickyBar    StickyBar    `json:"stickyBar"`
	Footer       Footer       `json:"footer"`
}

type Meta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Header struct {
	LogoEyebrow  string `json:"logoEyebrow"`
	LogoWordmark string `json:"logoWordmark"`
	BookLabel    string `json:"bookLabel"`
}

type Anchor struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

type Hero struct {
	Chips     []string `json:"chips"`
	H1        string   `json:"h1"`
	Subhead   string   `json:"subhead"`
	TrustLine string   `json:"trustLine"`
	PrimaryCTA   string `json:"primaryCTA"`
	SecondaryCTA string `json:"secondaryCTA"`
	Image     string   `json:"image"`
	ImageAlt  string   `json:"imageAlt"`
}

type Stat struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Treat struct {
	Eyebrow string        `json:"eyebrow"`
	H2      string        `json:"h2"`
	Subhead string        `json:"subhead"`
	Cards   []ServiceCard `json:"cards"`
	CTA     string        `json:"cta"`
}

type ServiceCard struct {
	Number string `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Image  string `json:"image"`
}

type Comparison struct {
	Eyebrow   string          `json:"eyebrow"`
	H2        string          `json:"h2"`
	Subhead   string          `json:"subhead"`
	ThemLabel string          `json:"themLabel"`
	UsLabel   string          `json:"usLabel"`
	Rows      []ComparisonRow `json:"rows"`
	CTA       string          `json:"cta"`
}

type ComparisonRow struct {
	Them string `json:"them"`
	Us   string `json:"us"`
}

type HowItWorks struct {
	Eyebrow string `json:"eyebrow"`
	H2      string `json:"h2"`
	Subhead string `json:"subhead"`
	Steps   []Step `json:"steps"`
	CTA     string `json:"cta"`
}

type Step struct {
	Number string `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type WhyWTG struct {
	Eyebrow  string   `json:"eyebrow"`
	H2       string   `json:"h2"`
	Subhead  string   `json:"subhead"`
	Bullets  []Bullet `json:"bullets"`
	CTA      string   `json:"cta"`
	Image    string   `json:"image"`
	ImageAlt string   `json:"imageAlt"`
}

type Bullet struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Clinicians struct {
	Eyebrow  string      `json:"eyebrow"`
	H2       string      `json:"h2"`
	Subhead  string      `json:"subhead"`
	Featured []Clinician `json:"featured"`
	CTA      string      `json:"cta"`
}

type Clinician struct {
	Name       string `json:"name"`
	Credential string `json:"credential"`
	City       string `json:"city"`
	Treats     string `json:"treats"`
	Photo      string `json:"photo"`
}

type MidCTA struct {
	H2   string `json:"h2"`
	Body string `json:"body"`
	CTA  string `json:"cta"`
}

type Pricing struct {
	Eyebrow       string   `json:"eyebrow"`
	H2            string   `json:"h2"`
	Subhead       string   `json:"subhead"`
	PriceBlock    string   `json:"priceBlock"`
	PriceUnit     string   `json:"priceUnit"`
	Reimbursement string   `json:"reimbursement"`
	Steps         []string `json:"steps"`
	CTA           string   `json:"cta"`
}

type Locations struct {
	Eyebrow string   `json:"eyebrow"`
	H2      string   `json:"h2"`
	Subhead string   `json:"subhead"`
	Offices []Office `json:"offices"`
	CTA     string   `json:"cta"`
}

type Office struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Hours   string `json:"hours"`
	Online  bool   `json:"online"`
}

type Testimonials struct {
	Eyebrow string        `json:"eyebrow"`
	H2      string        `json:"h2"`
	Reviews []Testimonial `json:"reviews"`
	CTA     string        `json:"cta"`
}

type Testimonial struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
	City   string `json:"city"`
}

type FAQ struct {
	Eyebrow  string    `json:"eyebrow"`
	H2       string    `json:"h2"`
	Items    []FAQItem `json:"items"`
	CTA      string    `json:"cta"`
	Image    string    `json:"image"`
	ImageAlt string    `json:"imageAlt"`
}

type FAQItem struct {
	Q string `json:"q"`
	A string `json:"a"`
}

type Form struct {
	Eyebrow     string         `json:"eyebrow"`
	H2          string         `json:"h2"`
	Subhead     string         `json:"subhead"`
	Locations   []string       `json:"locations"`
	Consent     string         `json:"consent"`
	Submit      string         `json:"submit"`
	PhoneFallback string       `json:"phoneFallback"`
	Reassurance string         `json:"reassurance"`
	Image       string         `json:"image"`
	ImageAlt    string         `json:"imageAlt"`
}

type StickyBar struct {
	Text         string `json:"text"`
	PrimaryCTA   string `json:"primaryCTA"`
	SecondaryCTA string `json:"secondaryCTA"`
}

type Footer struct {
	Offices  []Office `json:"offices"`
	Legal    string   `json:"legal"`
}
