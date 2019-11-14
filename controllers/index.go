package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/gigin/core"
	"net/http"
)

type IndexController struct {
	App *core.App
}

func NewIndexController(app *core.App) *IndexController {
	return &IndexController{App: app}
}

func (c *IndexController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "i am homepage",
		"mysql":   c.App.ConfigHandler.Get("mysql"),
	})
}
