package server

import (
	"wtg/landings/server/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.Static("/static", "./static")

	routes.LoadRoutes(r)

	_ = r.Run(":8080")
}
