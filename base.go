package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Container *App

type App struct {
	ConfigHandler *viper.Viper
	Engine *gin.Engine
}

func (app *App) SetRouter()  {
	app.Engine.GET("/", Index)
	api := app.Engine.Group("/api")
	{
		api.POST("/categories", CategoryStore)
	}
}

func (app *App) Run()  {
	if err := app.Engine.Run(app.ConfigHandler.GetString("web.port")); err != nil {
		fmt.Errorf("app.Run()出现错误%s", err.Error())
	}
}

func init()  {
	configHandler := viper.New()
	configHandler.SetConfigName("config")
	configHandler.AddConfigPath(".")
	err := configHandler.ReadInConfig() // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")
	Container = &App{ConfigHandler: configHandler, Engine: engine}
	Container.SetRouter()
}
