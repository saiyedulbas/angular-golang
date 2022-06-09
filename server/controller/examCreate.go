package controller

import (
	"github.com/gin-gonic/gin"
)
func ExamCreate(c *gin.Context){
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
				 Name string `json:"name"`
				 Duration int `json:"duration"`
				 Ending string `json:"ending"`
				 OrganizationId int `json:"organizationid"`
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			  name := body.Name
			  duration := body.Duration
			  ending := body.Ending
			  organizationonId := body.OrganizationId
			  
			type Exam struct {
				Name string `json:"name"`
				Duration int `json:"duration"`
				Ending string `json:"ending"`
				OrganizationId int `json:"organizationonid"`
				
			}
			var message2 string
			exam := Exam{Name: name, Duration: duration, Ending: ending, OrganizationId: organizationonId }
			result := db.Create(&exam)
			if result.RowsAffected > 0{
				message2 = "Exam Created"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Exam not Created"
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