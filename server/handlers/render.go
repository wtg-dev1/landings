package handlers

import (
	"bytes"
	"net/http"

	"wtg/landings/app"

	"github.com/gin-gonic/gin"
)

// Page returns a gin handler that renders (city, variant) through the
// shared app.RenderPage path. Same renderer the static export uses, so
// any divergence between dev and prod fails identically in both.
//
// route is the URL segment served at /<city>/<route> — needed so the
// FormName matches the JS getFormName() helper, which derives from the
// path, not the variant filename (e.g. variant "general" serves at
// route "therapy" → "Brooklyn Therapy Form").
func Page(city, variant, route string) gin.HandlerFunc {
	formName := app.GetFormName("/" + city + "/" + route)
	return func(c *gin.Context) {
		content, err := app.LoadPageContent(city, variant)
		if err != nil {
			c.String(http.StatusInternalServerError, "content load error")
			return
		}
		content.FormName = formName
		var buf bytes.Buffer
		if err := app.RenderPage(content, &buf); err != nil {
			c.String(http.StatusInternalServerError, "render error")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
	}
}
