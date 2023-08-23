package initializers

import (
	"fmt"

	"github.com/Dizzy-nt/go-rest-api/models"
)

func SyncDatabase() {
	fmt.Println("Syncing database")
	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Admin{})
	Db.AutoMigrate(&models.Warung{})
	Db.AutoMigrate(&models.Review{})
}
