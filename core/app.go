package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iamMarkchu/gigin/controllers"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type App struct {
	ConfigHandler *viper.Viper
	Engine        *gin.Engine
}

func NewApp() *App {
	configHandler := viper.New()
	configHandler.SetConfigName("config")
	configHandler.AddConfigPath(".")
	err := configHandler.ReadInConfig() // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	engine := gin.Default()
	return &App{ConfigHandler: configHandler, Engine: engine}
}

func (app *App) Run() {
	app.SetRouters()
	if err := app.Engine.Run(app.ConfigHandler.GetString("web.port")); err != nil {
		log.Println("app.Run()出现错误:", err.Error())
	}
}

func (app *App) SetRouters() {
	indexController := controllers.NewIndexController(app)
	app.Engine.GET("/", indexController.Index)
	app.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
