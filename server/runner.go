package server

import (
	"strings"
	"wtg/landings/server/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	// Long-lived cache on /static/* — PSI flags the gin default (1d) as
	// "inefficient cache lifetimes." 1 year is safe for the images we
	// hash-rename or replace under a new path; if you ever overwrite a
	// file in place without changing its URL, drop `immutable` or add
	// a version query string (e.g. main.css?v=BUILD_ID) before this.
	r.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/static/") {
			c.Header("Cache-Control", "public, max-age=31536000, immutable")
		}
		c.Next()
	})

	r.Static("/static", "./static")

	routes.LoadRoutes(r)

	_ = r.Run(":8080")
}
