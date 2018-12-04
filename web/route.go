package web

import (
	"fmt"
	"restdoc-api/app"
	"restdoc-api/handlers"

	"github.com/gin-gonic/gin"
)

// Route 路由配置
func Route(r *gin.Engine) {
	// PING & PONG
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("pong from <%s>", app.Config.SiteName))
	})

	// NOT FOUND
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "not found"})
	})

	// USER APIS
	user(r)

	// ADMIN APIS
	admin(r)
}

func user(r *gin.Engine) {
	g := r.Group("/api/v1")

	g.POST("/project", handlers.NewProject)
	g.GET("/projects", handlers.GetProjects)
	g.GET("/projects/:id", handlers.GetProject)
}

func admin(r *gin.Engine) {
	g := r.Group("/admin-api")

	g.GET("", nil)
}
