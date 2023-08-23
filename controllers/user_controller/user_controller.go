package user_controller

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Dizzy-nt/go-rest-api/initializers"
	"github.com/Dizzy-nt/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(c *gin.Context) {
	// Get the requested body for user
	//Get the email/pass off req body
	var body struct {
		Email       string
		Password    string
		NamaLengkap string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request, Failed to read body",
		})
		return
	}
	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error, Failed to hash password",
		})
		return
	}
	//Create the user
	user := models.User{
		Id:          0,
		Email:       body.Email,
		Password:    string(hash),
		NamaLengkap: body.NamaLengkap,
		Token:       "",
	}
	//Insert the user
	result := initializers.Db.Create(&user)
	initializers.Db.Model(&user).Where("id = ?", user.Id).Update("nama-lengkap", body.NamaLengkap)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error, Failed to create user",
		})
		return
	}
	//Return the user
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
func Login(c *gin.Context) {
	var body struct {
		Email       string
		Password    string
		NamaLengkap string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request, Failed to read body",
		})
		return
	}
	//look up requested user
	var user models.User
	result := initializers.Db.First(&user, "email =?", body.Email)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found, User not found",
		})
		return
	}
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized, Wrong password",
		})
		return
	}
	//Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token",
		})
		return
	}
	initializers.Db.Model(&user).Where("id = ?", user.Id).Update("token", tokenString)
	//Send it back
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm logged in",
	})
}

func Index(c *gin.Context) {
	var user []models.User
	initializers.Db.Find(&user)
	c.JSON(http.StatusOK, gin.H{"user list": user})
}

func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := initializers.Db.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data-nya ga ketemu nih :("})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": user})
}

func Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	initializers.Db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
func Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if initializers.Db.Model(&user).Where("id=?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data-nya ga ketemu nih :("})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui :3"})
}

func Delete(c *gin.Context) {
	var user models.User
	var input struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	if initializers.Db.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal menghapus data :("})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
