// You can edit this code!
// Click here and start typing.
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saiyedulbas/second/controller"
	"github.com/saiyedulbas/second/dbconn"
	"gorm.io/gorm"
)
var (
	db             *gorm.DB                  = dbconn.SetupConnection()
)

func main() {
	
	defer dbconn.CloseDatabaseConnection(db)
 


		router := gin.Default()
		config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  config.AllowHeaders = []string{"Token","Content-Type"}
 router.Use(cors.New(config)) 
        authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/register",controller.Register)
		authRoutes.POST("/login",controller.Login)
	}
	   userRoutes := router.Group("api/user")
	{
		userRoutes.GET("/profile/:id",controller.UserProfile)
		userRoutes.POST("/fileupload",controller.UploadFile)
		userRoutes.GET("/userlist",controller.UserList)
		userRoutes.POST("/useredit/:id",controller.UserEdit)
		userRoutes.POST("/userdelete/:id",controller.UserDelete)
		
	}
	organizationRoutes := router.Group("api/organization")
	{
		organizationRoutes.POST("/create",controller.OrganizationCreate)
		organizationRoutes.POST("/edit/:id",controller.OrganizationEdit)
		organizationRoutes.GET("/list",controller.OrganizationList)
	}
	contestRoutes := router.Group("api/contest")
	{
		contestRoutes.POST("/",controller.ContestCreate)
		contestRoutes.POST("/edit/:id",controller.ContestEdit)
		contestRoutes.GET("/",controller.ContestList)
	}
	contestResultRoutes := router.Group("api/contest-result")
	{
		contestResultRoutes.POST("/",controller.ContestResultCreate)
		contestResultRoutes.POST("/edit/:id",controller.ContestResultEdit)
		contestResultRoutes.GET("/",controller.ContestResultList)
	}
	examRoutes := router.Group("api/exam")
	{
		examRoutes.POST("/",controller.ExamCreate)
		examRoutes.GET("/",controller.ExamList)
		examRoutes.GET("/:id",controller.ParticularExam)
	}
	questionRoutes := router.Group("api/question")
	{
		questionRoutes.POST("/",controller.QuestionCreate)
		
	}
	mcqRoutes := router.Group("api/mcq")
	{
		mcqRoutes.POST("/",controller.McqCreate)
		
		
	}
		router.Run(":8091")
	

	
}
