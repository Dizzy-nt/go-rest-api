package models

type Admin struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Username    string `gorm:"type:varchar(127)" json:"username"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	Token       string `gorm:"type:varchar(255)" json:"token"`
}