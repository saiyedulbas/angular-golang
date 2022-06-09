package controller

import (
	"github.com/gin-gonic/gin"
)
func McqCreate(c *gin.Context){
	token := c.Request.Header["Token"]
	if token == nil{
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "token not set",
			
		}) 
	}else{
		if ExampleParse(token[0],MySigningKey){ 
			type Body struct {
				// json tag to de-serialize json body
				 Choice string `json:"choice"`
				 QuestionId int `json:"questionid"`
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			choice := body.Choice
			questionid := body.QuestionId
			  
			  
			type Mcq struct {
				Choice string `json:"choice"`
				QuestionId int `json:"questionid"`
				
			}
			var message2 string
			mcq := Mcq{Choice: choice, QuestionId: questionid }
			result := db.Create(&mcq)
			if result.RowsAffected > 0{
				message2 = "Mcq Created"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Mcq not Created"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": result.Error,
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