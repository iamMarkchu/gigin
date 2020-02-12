package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// route  /
func Index(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "index/index.tmpl", gin.H{
		"title": "index",
	})
}

// route /api/categories [get]
func CategoryStore(ctx *gin.Context)  {

}

