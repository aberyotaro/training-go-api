package routes

import (
	"github.com/aberyotaro/go-api-sandbox/internal/di"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine, h *di.Handlers) {
	v1 := g.Group("/v1")
	v1.GET("/users/:id", h.UserHandler.GetUserByID)
}
