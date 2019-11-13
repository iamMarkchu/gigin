package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	app := gin.Default()
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	app.GET("/find-me", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"appName":     viper.GetString("appName"),
			"mysqlConfig": viper.GetStringMapString("mysql"),
		})
	})
	if err := app.Run(":8888"); err != nil {
		log.Println("app.Run()出现错误:", err.Error())
	}
}
