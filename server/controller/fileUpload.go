package controller

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)
const (
    S3_REGION = "ap-southeast-1"
    S3_BUCKET = "dgraphy.pics/contest/eid-special-2022"
)
func UploadFile(c *gin.Context){
	token := c.Request.Header["Token"]
	if token == nil{
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "token not set",
			
		}) 
	}else{
		if ExampleParse(token[0],MySigningKey){ 
			file, _ := c.FormFile("file")
			dst := "uploads/"+file.Filename
	
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		// The session the S3 Uploader will use
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Credentials: credentials.NewSharedCredentials(".aws/credentials", "default"),
		})
		if err != nil{
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": err.Error(),
				
			}) 
		}
	
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	
	f, err  := os.Open(dst)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": err.Error(),
			
		}) 
	}
	
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("dgraphy.pics/contest/eid-special-2022/"),
		Key:    aws.String(file.Filename),
		Body:   f,
	})
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": err.Error(),
			
		}) 
	}
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": result,
		
	}) 
		}else{
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Not validated",
				
			})
		}
		
	}
	

}
