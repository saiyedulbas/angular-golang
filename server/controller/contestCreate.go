package controller

import (
	"github.com/gin-gonic/gin"
)
func ContestCreate(c *gin.Context){
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
				 Title string `json:"title"`
				 Duration int `json:"duration"`
				 Ending string `json:"ending"`
				 OrganizationId int64 `json:"organizationonid"`
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			  title := body.Title
			  duration := body.Duration
			  ending := body.Ending
			  organizationonId := int(body.OrganizationId)
			  
			type Contest struct {
				Title string `json:"title"`
				Duration int `json:"duration"`
				Ending string `json:"ending"`
				OrganizationId int `json:"organizationonid"`
				
			}
			var message2 string
			contest := Contest{Title: title, Duration: duration, Ending: ending, OrganizationId: organizationonId }
			result := db.Create(&contest)
			if result.RowsAffected > 0{
				message2 = "Contest Created"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Contest not Created"
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