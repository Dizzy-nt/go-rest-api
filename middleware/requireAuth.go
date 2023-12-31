package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Dizzy-nt/go-rest-api/initializers"
	"github.com/Dizzy-nt/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("In middleware")
	//Get the cookie off req
	tokenString,err:=c.Cookie("Authorization")
	if err!=nil{
		c.JSON(401,gin.H{"error":"No token provided"})
		return
	}
	//Decode / Validate it
	
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return  []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.JSON(401,gin.H{"error":"Token has expired"})
			return
		}
		//Find the user with token sub
		var user models.User
		initializers.Db.First(&user,claims["sub"])

		if user.Id == 0{
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to request
		c.Next()
	} else {
		c.JSON(401,gin.H{"error":"No token provided"})
		return
	}

	
}