package routes

import (
	"net/http"

	"github.com/VJ-Vijay77/miniProject/pkg/database"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"

	uuid "github.com/satori/go.uuid"
)

type Users struct {
	Username string 
	Password string 
}



func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)
}



func PostLogin(c *gin.Context) {
	var user []Users
	var status bool
	Fusername := c.Request.FormValue("username")
	Fpassword := c.Request.FormValue("password")


	// //checking in database
	db:= database.InitDB()
	//db.AutoMigrate(&Users{})
	db.Find(&user)
	for _,i := range user{
		if i.Username == Fusername && i.Password==Fpassword{
			status = true
			break
		}
	}
	//checking end


	if !status {
		dialog.Alert("Wrong Username or Password\n\t\tTry Again")
		c.Redirect(http.StatusSeeOther,"/login")
		return
	}

	cookie, err := c.Request.Cookie("session")
	if err != nil {
		uuid := uuid.NewV4()
		c.SetCookie("session", uuid.String(), 300, "/", "localhost", false, false)
	}
	_ = cookie
	c.Redirect(http.StatusSeeOther, "/home")

}




func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func PostSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}




func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}





func Home(c *gin.Context) {

	c.HTML(http.StatusOK, "welcomeuser.html", nil)
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
