package main

import (
	adminController "github.com/Dizzy-nt/go-rest-api/controllers/admin_controller"
	productController "github.com/Dizzy-nt/go-rest-api/controllers/product_controller"
	reviewController "github.com/Dizzy-nt/go-rest-api/controllers/review_controller"
	userController "github.com/Dizzy-nt/go-rest-api/controllers/user_controller"
	warungController "github.com/Dizzy-nt/go-rest-api/controllers/warung_controller"
	"github.com/Dizzy-nt/go-rest-api/initializers"
	"github.com/Dizzy-nt/go-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func init()  {
	initializers.ConnectToDatabase()
	initializers.LoadEvnVariables()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Coba CRUD
	r.GET("/api/product", productController.Index)

	//signup, login, logout
	r.POST("/api/user/signup", userController.SignUp)
    r.POST("/api/user/login", userController.Login)
	r.GET("/api/user/validate",middleware.RequireAuth ,userController.Validate)
    //r.POST("/api/user/logout", userController.Logout)

	//Find all list Admin
	r.GET("/api/admin", adminController.Index)
	//Find Admin by Id
	r.GET("/api/admin/:id", adminController.Show)
	//Create new Admin
	r.POST("/api/admin", adminController.Create)
	//Edit Admin by Id
	r.PUT("/api/admin/:id", adminController.Update)
	//Delete Admin by Id
	r.DELETE("/api/admin", adminController.Delete)

	//Find all list User
	r.GET("/api/user", userController.Index)
	//Find User by Id
	r.GET("/api/user/:id", userController.Show)
	//Create new User
	r.POST("/api/user", userController.Create)
	//Edit User by Id
	r.PUT("/api/user/:id", userController.Update)
	//Delete User by Id
	r.DELETE("/api/user", userController.Delete)

	//Find all list Warung
	r.GET("/api/warung", warungController.Index)
	//Find Warung by Id
	r.GET("/api/warung/:id", warungController.Show)
	//Create new Warung
	r.POST("/api/warung", warungController.Create)
	//Edit warung by Id
	r.PUT("/api/warung/:id", warungController.Update)
	//Delete warung by Id
	r.DELETE("/api/warung", warungController.Delete)

	//Find all list Review
	r.GET("/api/review", reviewController.Index)
	//Find Review by Id
	r.GET("/api/review/:id", reviewController.Show)
	//Create new Review
	r.POST("/api/review", reviewController.Create)
	//Edit Review by Id
	r.PUT("/api/review/:id", reviewController.Update)
	//Delete Review by Id
	r.DELETE("/api/review", reviewController.Delete)
	r.Run()
}
