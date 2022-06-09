package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/saiyedulbas/second/dbconn"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
var (
	db             *gorm.DB                  = dbconn.SetupConnection()
)
func Register(c *gin.Context) {
	type Body struct {
		// json tag to de-serialize json body
		 Name string `json:"name"`
		 Email string `json:"email"`
		 Password string `json:"password"`
	  }
	  body:=Body{}
	jsonParsedBody := c.BindJSON(&body)
	_ = jsonParsedBody
	
	type User struct {
		Name string 
		Email  string 
		Password string
		
	}
	name := body.Name
	email := body.Email
	password_prior := body.Password
	password := []byte(password_prior)
	hashedPassword, err9 := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	stringedHashedPassword := string(hashedPassword)
	result := []User{}
	mrinal := db.Where("name = ?", name).Find(&result)
	var message3 string
	if mrinal.RowsAffected > 0{
		message3 = "Username already exists"
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": message3,
		})
		
	}else{
		user := User{Name: name, Email: email, Password: stringedHashedPassword }
	result := db.Create(&user)
	
	var message2 string
	if result.RowsAffected > 0{
		message2 = "Data Inserted"
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": message2,
		})
	}else{
		message2 = "Data not Inserted"
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": message2,
		})
	}
	}
	
	
	
	_ = result
	_ = err9
	

	
}