package main

import (
	_ "net/http"

	"github.com/VJ-Vijay77/miniProject/pkg/routes"
	"github.com/gin-gonic/gin"
)

const Dsn = `postgres://vijay:zmxmcmvbn@localhost/databasevj`

func main() {
	route := gin.Default()

	//route.Static("public/", "./public/css")
	route.LoadHTMLGlob("templates/*.html")
	route.GET("/login",routes.Login)
	route.POST("/login",routes.PostLogin)
	route.GET("/logout",routes.Logout)
	route.GET("/signup",routes.Signup)
	route.GET("/admin",routes.Admin)
	route.GET("/home",routes.Home)
	route.Run()

}
