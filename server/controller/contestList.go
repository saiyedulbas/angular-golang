package controller

import (
	"github.com/gin-gonic/gin"
)
func ContestList(c *gin.Context){
	token := c.Request.Header["Token"]
	if token == nil{
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "token not set",
			
		}) 
	}else{
		if ExampleParse(token[0],MySigningKey){ 
			type User struct {
				Id int64 `json:"id"`
				Name string `json:"name"`
				Email  string `json:"email"`
				
				
			}
			type Organization struct {
				Id int `json:"id"`
				Name string `json:"name"`
				Contact_Email string `json:"contact_email"`
				Description string `json:"description"`
				User *User `json:"manager"`
				UserId int `json:"managerid"`
				  
			}
			type Contest struct {
				Title string `json:"title"`
				Duration int `json:"duration"`
				Ending string `json:"ending"`
				Organization *Organization
				OrganizationId int `json:"organizationonid"`
				
			}
			
			result := []Contest{}
			
			query := db.Preload("Organization.User").Find(&result)
			
			
		
			_ = query
			
			c.JSON(200, gin.H{
				"status":  "Ok",
				"message": "Contest List",
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