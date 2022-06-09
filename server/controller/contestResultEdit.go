package controller

import (
	"github.com/gin-gonic/gin"
)
func ContestResultEdit(c *gin.Context){
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
				 Result string `json:"result"`
				 ContestId int `json:"contestid"`
				 UserId int `json:"userid"`
				 Submitted string `json:"submitted"`
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			  result := body.Result
			  contestid := body.ContestId
			  userid := body.UserId
			  submitted := body.Submitted
			  
			type ContestResult struct {
				Result string `json:"result"`
				ContestId int `json:"contestid"`
				UserId int `json:"userid"`
				Submitted string `json:"submitted"`
				
			}
			var message2 string
			contestResult := []ContestResult{}
			result2:= db.Model(&contestResult).Where("id = ?", id).Updates(ContestResult{Result: result, ContestId: contestid, UserId: userid, Submitted: submitted })
			if result2.RowsAffected > 0{
				message2 = "Contest Result Updated"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Contest Result not Updated"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": result2.Error,
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