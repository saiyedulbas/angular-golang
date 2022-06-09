package controller

import (
	"github.com/gin-gonic/gin"
)
func OrganizationEdit(c *gin.Context){
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
				 Name string `json:"name"`
				 Contact_Email string `json:"contact_email"`
				 Description string `json:"description"`
				 
			  }
			  body:=Body{}
			  jsonParsedBody := c.BindJSON(&body)
			_ = jsonParsedBody
			name := body.Name
			contact_email := body.Contact_Email
			description := body.Description
			
			
			type Organization struct {
				Name string `json:"name"`
				Contact_Email string `json:"contact_email"`
				Description string `json:"description"`
				
				
			}
			var message2 string
			organization := []Organization{}
			result:= db.Model(&organization).Where("id = ?", id).Updates(Organization{Name: name, Contact_Email: contact_email,Description: description})
			if result.RowsAffected > 0{
				message2 = "Organization Updated"
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": message2,
				})
			}else{
				message2 = "Organization not Updated or nothing to update"
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