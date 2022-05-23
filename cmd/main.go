package main

import (
	_ "net/http"

	"github.com/VJ-Vijay77/miniProject/pkg/database"
	"github.com/VJ-Vijay77/miniProject/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	database.InitDB()

	//route.Static("public/", "./public/css")
	route.LoadHTMLGlob("templates/*.html")
	route.GET("/login", routes.Login)
	route.POST("/login", routes.PostLogin)
	route.GET("/logout", routes.Logout)
	route.GET("/signup", routes.Signup)
	route.POST("/signup", routes.PostSignup)
	route.GET("/admin", routes.Admin)
	route.GET("/home", routes.Home)
	route.Run(":8080")

}
