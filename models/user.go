package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Email       string `gorm:"primaryKey" json:"email"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	NamaLengkap string `gorm:"type:varchar(255)" json:"nama-lengkap"`
	Token       string `gorm:"type:varchar(255)" json:"token"`
}
