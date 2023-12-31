package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/random_note/conf"
	"github.com/random_note/controller"
	"github.com/random_note/dao"
	"github.com/random_note/models"
	"github.com/spf13/viper"
)

func main() {
	// connect db
	//err := dao.InitMySQLByFile()
	conf.LoadConfig()
	err := dao.InitMySQLByViper()
	if err != nil {
		panic("failed to connect database")
	}

	// migrate
	dao.DB.AutoMigrate(&models.Todo{})
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.AddTodo)
		v1Group.GET("/todo", controller.ViewTodo)
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	serverPort := fmt.Sprintf("0.0.0.0:%s", viper.GetString("GINS_PORT"))
	r.Run(serverPort)
}
