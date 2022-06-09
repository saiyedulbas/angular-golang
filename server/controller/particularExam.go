package controller

import (
	"github.com/gin-gonic/gin"
)

func ParticularExam(c *gin.Context){
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
				
				
			}
			type Mcq struct {
				Id int `json:"id"`
				Choice string `json:"choice"`
				QuestionId int `json:"questionid"`
				
			}
			type Question struct {
				Id int `json:"id"`
				QuestionTitle string `json:"question_title"`
				ExamId int `json:"examid"`
				Mcq []Mcq
				
			}
			type Organization struct {
				Id int `json:"id"`
				Name string `json:"name"`
				Contact_Email string `json:"contact_email"`
				Description string `json:"description"`
				User *User `json:"manager"`
				UserId int `json:"managerid"`
			}
			type Exam struct {
				Id int `json:"id"`
				Name string `json:"name"`
				Duration int `json:"duration"`
				Ending string `json:"ending"`
				Organization *Organization 
				OrganizationId int `json:"organizationid"`
				Question []Question

			}
			result := []Exam{}
			query := db.Where("id = ?", id).Preload("Organization.User").Preload("Question.Mcq").Find(&result)
			
			
		
			_ = query
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "exam info",
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
 