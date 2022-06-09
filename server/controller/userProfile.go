package controller

import (
	"github.com/gin-gonic/gin"
)

func UserProfile(c *gin.Context){
	token := c.Request.Header["Token"]
	id := c.Param("id")
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
				Password string `json:"password"`
				
			}
			result := []User{}
			mrinal := db.Where("id = ?", id).Find(&result)
			_ = mrinal
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": &result,
				
			}) 
		 }else{
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Not validated",
				
			})
		}
	}
	
	
	 
}
 