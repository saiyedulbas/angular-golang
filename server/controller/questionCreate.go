package controller

import (
	"github.com/gin-gonic/gin"
)
func QuestionCreate(c *gin.Context){
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
				 QuestionTitle string `json:"question_title"`
				 ExamId int `json:"examid"`
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			question_title := body.QuestionTitle
			examid := body.ExamId
			  
			  
			type Question struct {
				QuestionTitle string `json:"question_title"`
				ExamId int `json:"examid"`
				
			}
			var message2 string
			question := Question{QuestionTitle: question_title, ExamId: examid }
			result := db.Create(&question)
			if result.RowsAffected > 0{
				message2 = "Question Created"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Question not Created"
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