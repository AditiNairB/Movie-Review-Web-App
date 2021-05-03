package authController

import (
	"GoLang-WebServer/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// only names that begin with caps are exported!

// LoginAPI /api/auth/login
// SignUpAPI /api/auth/signup
// LogoutAPI /api/auth/logout

type LoginData struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignUpData struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
}

func Login(c *gin.Context) {
	// uncomment the following two lines to work with form data and comment the lines between 'JSON INPUT'
	// email := c.Request.FormValue("email")
	// password := c.Request.FormValue("password")

	// JSON input.
	var login LoginData
	c.BindJSON(&login)

	email := login.Email
	password := login.Password

	// JSON input

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill all fields: email and password"})
		return
	}

	userId, err := models.VerifyLogin(email, []byte(password))

	// log.Println(err)
	if err == nil {
		tokenString, expirationTime := SignJWT(userId.UserId, userId.FirstName, userId.LastName)
		// c.SetCookie("token", tokenString, expirationTime, "/", "localhost", true, false)
		c.JSON(http.StatusOK, gin.H{"token": tokenString, "expirationTime": expirationTime})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
}

func SignUp(c *gin.Context) {
	// uncomment the following four lines to work with form data and comment the lines between 'JSON INPUT'
	// firstName := c.Request.FormValue("firstName")
	// lastName := c.Request.FormValue("lastName")
	// password := c.Request.FormValue("password")
	// email := c.Request.FormValue("email")

	// JSON INPUT
	var signup SignUpData
	c.BindJSON(&signup)

	email := signup.Email
	password := signup.Password
	firstName := signup.FirstName
	lastName := signup.LastName

	// JSON INPUT

	if firstName == "" || lastName == "" || email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please fill all fields: first name, last name, email and password"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with password"})
		return
	}

	err = models.CreateUser(email, string(hashedPassword), firstName, lastName)
	if err != nil {

		log.Println(err)
		if strings.Contains(err.Error(), "E11000") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Please check the input and try again"})
		}

	} else {
		userId, err := models.VerifyLogin(email, []byte(password))

		// log.Println(err)
		if err == nil {
			tokenString, expirationTime := SignJWT(userId.UserId, userId.FirstName, userId.LastName)
			// c.SetCookie("token", tokenString, expirationTime, "/", "localhost", true, false)
			c.JSON(http.StatusOK, gin.H{"token": tokenString, "expirationTime": expirationTime})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			return
		}
		// c.JSON(http.StatusOK, gin.H{})
	}

}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
