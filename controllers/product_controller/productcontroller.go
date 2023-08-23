package productcontroller

import (
	"net/http"

	"github.com/Dizzy-nt/go-rest-api/initializers"
	"github.com/Dizzy-nt/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var prodak []models.Productt
	initializers.Db.Find(&prodak)
	c.JSON(http.StatusOK, gin.H{"product": prodak})
}
