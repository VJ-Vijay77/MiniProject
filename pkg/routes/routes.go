package routes

import (
	"net/http"

	"github.com/VJ-Vijay77/miniProject/pkg/database"
	"github.com/gin-gonic/gin"
	"tawesoft.co.uk/go/dialog"

	uuid "github.com/satori/go.uuid"
)

type Users struct {
	ID       int
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
	db := database.InitDB()
	//	db.AutoMigrate(&Users{})
	db.Find(&user)
	for _, i := range user {
		if i.Username == Fusername && i.Password == Fpassword {
			status = true
			break
		}
	}
	// if ok := db.Raw("SELECT username,password users WHERE username=? AND password=?", Fusername, Fpassword); ok.Error != nil {
	// 	status = false
	// }

	//checking end

	if !status {
		dialog.Alert("Wrong Username or Password\n\t\tTry Again")
		c.Redirect(http.StatusSeeOther, "/login")
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
	var user []Users
	var status bool = true
	Fname := c.Request.FormValue("name")
	FusernameN := c.Request.FormValue("username")
	Fpassword := c.Request.FormValue("password")

	//database things
	db := database.InitDB()
	db.AutoMigrate(&Users{})
	db.Find(&user)

	for _, i := range user {
		if i.Username == FusernameN {
			status = false
			break
		}
	}

	if !status {
		dialog.Alert("hello %s , The username is already taken", Fname)
		c.Redirect(303, "/signup")
		return

	}

	db.Create(&Users{Username: FusernameN, Password: Fpassword})
	dialog.Alert("Hey %s, Your account is successfully created. Click OK to LOGIN!", Fname)
	c.Redirect(http.StatusSeeOther, "/login")

}

//database things end

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", nil)
}

func PostAdmin(c *gin.Context) {
	Fusername := c.Request.FormValue("username")
	Fpassword := c.Request.FormValue("password")

	if Fusername != "adminvijay" || Fpassword != "12345" {
		dialog.Alert("Wrong Username or Password , Check Again!")
		c.Redirect(303, "/admin")
		return
	}

	// c.HTML(200,"welcomeadmin.html",nil)

	c.Redirect(303, "/wadmin")

}

func Wadmin(c *gin.Context) {
	var user []Users
	// //var usersnew

	db := database.InitDB()
	//db.AutoMigrate(&Users{})
	// db.Find(&user)
	//  var i Users
	//  var ind int
	var us = [11]string{}
	// for ind, i := range user {
	// 	us[ind]=i.Username
	// }
	var id = [11]int{}
	db.Raw("SELECT id,username FROM users").Scan(&user)
	for ind, i := range user {
		us[ind], id[ind] = i.Username, i.ID

	}

	c.HTML(http.StatusOK, "welcomeadmin.html", gin.H{

		"users": us,
		"id":    id,
	})
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

func DeleteUser(c *gin.Context) {
	var user Users
	name := c.Param("name")
	db := database.InitDB()
	db.Where("username=?", name).Delete(&user)
	c.Redirect(303, "/wadmin")

}

func UpdateUser(c *gin.Context) {

	updateData := c.Request.FormValue("updatedata")
	var user Users
	name := c.Param("name")
	db := database.InitDB()
	db.Model(&user).Where("username=?", name).Update("username", updateData)
	c.Redirect(303, "/wadmin")
}


func CreateUser(c *gin.Context) {
	var user []Users
	var status bool = true
	
	FusernameN := c.Request.FormValue("username")
	Fpassword := c.Request.FormValue("password")

	//database things
	db := database.InitDB()
	db.AutoMigrate(&Users{})
	db.Find(&user)

	for _, i := range user {
		if i.Username == FusernameN {
			status = false
			break
		}
	}

	if !status {
		dialog.Alert("hello Admin , The username is already in Use")
		c.Redirect(303, "/wadmin")
		return

	}

	db.Create(&Users{Username: FusernameN, Password: Fpassword})
	dialog.Alert("Hey Admin, Account is successfully created.")
	c.Redirect(http.StatusSeeOther, "/wadmin")

}



//cache memory clearing
func ClearCache(w gin.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
