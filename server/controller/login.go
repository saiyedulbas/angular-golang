package controller

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
const (
    MySigningKey = "WOW,MuchShibe,ToDogge"
)

func Login(c *gin.Context) {
	type Body struct {
		// json tag to de-serialize json body
		 Name string `json:"name"`
		 Password string `json:"password"`
	  }
	  body:=Body{}
	jsonParsedBody := c.BindJSON(&body)
	_ = jsonParsedBody
	name := body.Name
	
	
        type User struct {
			Id int64 `json:"id"`
			Name string `json:"name"`
			Email  string `json:"email"`
			Password string `json:"password"`
			
		}
		
		
		
		password := body.Password
		result := []User{}
		mrinal := db.Where("name = ?", name).Find(&result)
		
	
		var message3 string
	
		if mrinal.RowsAffected > 0{
			storedPassword := result[0].Password
		bytePassword := []byte(storedPassword)
		bytePassword2 := []byte(password)
		err := bcrypt.CompareHashAndPassword(bytePassword, bytePassword2)
		if err != nil{
			c.JSON(200, gin.H{
				"status":  "Ok",
				"message": "Wrong Password",
			})
			
		}else{
			// Create the token
			createdToken, err := ExampleNew([]byte(MySigningKey),result[0].Id, result[0].Name, result[0].Email)
			if err != nil {
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": "Logged In",
					"token": err,
				})
			}else{
				c.JSON(200, gin.H{
					"status":  "ok",
					"message": "Logged In",
					"token": createdToken,
				})
			}
			
			
		}
			
			
			
		}else{
			message3 = "Username Not Found"
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": message3,
			})
		}
	
	
	
	

	
}
func ExampleNew(mySigningKey []byte,id int64, name string, email string) (string, error) {
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)
    // Set some claims
    claims := make(jwt.MapClaims)
claims["user_id"] = id
claims["name"] = name
claims["email"] = email
claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
token.Claims = claims
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)
    return tokenString, err
}

func ExampleParse(myToken string, myKey string) bool{
    token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(myKey), nil
    })
    var validationMessage bool
    if err == nil && token.Valid {
        validationMessage = true
    } else {
        validationMessage = false
    }
   return validationMessage
}