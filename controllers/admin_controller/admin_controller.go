package admin_controller

import (
	"encoding/json"
	"net/http"

	"github.com/Dizzy-nt/go-rest-api/initializers"
	"github.com/Dizzy-nt/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var admin []models.Admin
	initializers.Db.Find(&admin)
	c.JSON(http.StatusOK, gin.H{"admin list": admin})
}

func Show(c *gin.Context) {
	var admin models.Admin
	id := c.Param("id")
	if err := initializers.Db.First(&admin, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data-nya ga ketemu nih :("})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": admin})
}

func Create(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	initializers.Db.Create(&admin)
	c.JSON(http.StatusOK, gin.H{"admin": admin})
}

func Update(c *gin.Context) {
	var admin models.Admin
	id := c.Param("id")

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if initializers.Db.Model(&admin).Where("id=?", id).Updates(&admin).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data-nya ga ketemu nih :("})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui :3"})
}

func Delete(c *gin.Context) {
	var admin models.Admin
	//input := map[string]string{"id":"0"}
	var input struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	if initializers.Db.Delete(&admin, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal menghapus data :("})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
