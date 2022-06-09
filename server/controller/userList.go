package controller

import (
	"github.com/gin-gonic/gin"
)
func UserList(c *gin.Context){
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
			result := []User{}
			query := db.Find(&result)
			_ = query
			c.JSON(200, gin.H{
				"status":  "Ok",
				"message": "User List",
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