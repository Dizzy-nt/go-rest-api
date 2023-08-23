package models

type Warung struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	NamaWarung string `gorm:"type:varchar(255)" json:"nama_warung"`
	Gambar     string `gorm:"type:varchar(255)" json:"gambar"`
	Rating     int64  `gorm:"type:int" json:"rating"`
	Lokasi     string `gorm:"type:varchar(255)" json:"lokasi"`
	Deskripsi  string `gorm:"type:text" json:"deskripsi"`
}
