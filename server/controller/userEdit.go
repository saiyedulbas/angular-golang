package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
func UserEdit(c *gin.Context){
	token := c.Request.Header["Token"]
	if token == nil{
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "token not set",
			
		}) 
	}else{
		if ExampleParse(token[0],MySigningKey){ 
			id := c.Param("id")
			type Body struct {
				// json tag to de-serialize json body
				 Name string `json:"name"`
				 Password string `json:"password"`
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			  name := body.Name
			  password_prior := body.Password
			  password := []byte(password_prior)
			  hashedPassword, err9 := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
			  stringedHashedPassword := string(hashedPassword)
			
			_ = err9
			type User struct {
				Name string 
				Password string 
				
			}
			var message2 string
			user := []User{}
			result:= db.Model(&user).Where("id = ?", id).Updates(User{Name: name, Password: stringedHashedPassword})
			if result.RowsAffected > 0{
				message2 = "Data Updated"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Data not Updated"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}
		}else{
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Not validated",
				
			})
		}
        
	}
	
	
}