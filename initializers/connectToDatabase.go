package initializers

import (
	//"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var Db *gorm.DB
func ConnectToDatabase() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("DB")
	db, err := gorm.Open(
		mysql.Open("root@tcp(localhost:3306)/go_restapi_gin"))
	if err != nil {
		panic(err)
	}
	Db=db
}
