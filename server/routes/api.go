package routes

import (
	"wtg/landings/app"
	"wtg/landings/server/handlers"

	"github.com/gin-gonic/gin"
)

func LoadRoutes(r *gin.Engine) {
	for _, s := range app.Sites {
		r.GET("/"+s.City+"/"+s.Route, handlers.Page(s.City, s.Variant, s.Route))
	}
}
