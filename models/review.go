package models

type Review struct {
	IdReview   int    `gorm:"type:int(10) primaryKey" json:"id-review"`
	IdWarung   int    `gorm:"type:int(10)" json:"id-warung"`
	NamaWarung string `gorm:"type:varchar(255)" json:"nama-warung"`
	IdUser     int    `gorm:"type:int(10)" json:"id-user"`
	Email      string `grom:"type:varchar(255)" json:"email"`
	Rating     int    `gorm:"type:int(1)" json:"rating"`
	Deskripsi  string `gorm:"type:text" json:"deskripsi"`
}
