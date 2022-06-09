package controller

import (
	"github.com/gin-gonic/gin"
)
func OrganizationCreate(c *gin.Context){
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
				 Contact_Email string `json:"contact_email"`
				 Description string `json:"description"`
				 Manager int64 `json:"manager"`
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			  name := body.Name
			  contact_email := body.Contact_Email
			  description := body.Description
			  manager := int(body.Manager)
			  
			type Organization struct {
				Name string `json:"name"`
				Contact_Email string `json:"contact_email"`
				Description string `json:"description"`
				UserId int `json:"userid"`
				
			}
			var message2 string
			organization := Organization{Name: name, Contact_Email: contact_email, Description: description, UserId: manager }
			result := db.Create(&organization)
			if result.RowsAffected > 0{
				message2 = "Organization Created"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Organization not Created"
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