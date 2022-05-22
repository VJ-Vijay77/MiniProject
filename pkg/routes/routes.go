package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	uuid "github.com/satori/go.uuid"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}

func PostLogin(c *gin.Context){
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	if username !="vijay" && password != "12345"{
		c.JSON(200,gin.H{"details":"Not Correct",})
		}
	cookie, err := c.Request.Cookie("session")
	if err != nil {
		uuid := uuid.NewV4()
		c.SetCookie("session", uuid.String(), 300, "/", "localhost", false, false)
	}
	_=cookie
		c.Redirect(http.StatusSeeOther,"/home")
}


func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}

func Home(c *gin.Context) {

	c.HTML(http.StatusOK, "welcomeuser.html",nil)
}

func Logout(c *gin.Context) {

	cookie, err := c.Request.Cookie("session")
	if err != nil {
		c.Redirect(303, "/login")
	}
	c.SetCookie("session", "", -1, "/", "localhost", false, false)
	_ = cookie
	c.Redirect(http.StatusSeeOther, "/login")
}

//cache memory clearing
func ClearCache(w gin.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
