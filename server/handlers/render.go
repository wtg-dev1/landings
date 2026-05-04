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
func Page(city, variant string) gin.HandlerFunc {
	return func(c *gin.Context) {
		content, err := app.LoadPageContent(city, variant)
		if err != nil {
			c.String(http.StatusInternalServerError, "content load error")
			return
		}
		var buf bytes.Buffer
		if err := app.RenderPage(content, &buf); err != nil {
			c.String(http.StatusInternalServerError, "render error")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
	}
}
