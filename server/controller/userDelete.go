package controller

import (
	"github.com/gin-gonic/gin"
)
func UserDelete(c *gin.Context){
	token := c.Request.Header["Token"]
	if token == nil{
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "token not set",
			
		}) 
	}else{
		if ExampleParse(token[0],MySigningKey){ 
			id := c.Param("id")
			
			type User struct {
				Name string 
				Password string 
				
			}
			var message2 string
			user := []User{}
			result:= db.Where("id = ?", id).Delete(&user)
			if result.RowsAffected > 0{
				message2 = "Data Deleted"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Data not Deleted"
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