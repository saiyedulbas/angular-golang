package controller

import (
	"github.com/gin-gonic/gin"
)
func OrganizationList(c *gin.Context){
	token := c.Request.Header["Token"]
	if token == nil{
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "token not set",
			
		}) 
	}else{
		if ExampleParse(token[0],MySigningKey){ 
			type User struct {
		
				Id int `json:"id"`
				Name string `json:"name"`
				Email  string `json:"email"`
				Password string `json:"password"`
				
			}
			type Organization struct {
				Name string `json:"name"`
				Contact_Email string `json:"contact_email"`
				Description string `json:"description"`
				User *User `json:"manager"`
				UserId int `json:"managerid"`
				  
			}
			
			result := []Organization{}
			
			query := db.Preload("User").Find(&result)
			
			
		
			_ = query
			
			c.JSON(200, gin.H{
				"status":  "Ok",
				"message": "Organization List",
				"data": &result,
			})
		}else{
            c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Not validated",
				
			})
		}
		
	}
	
}