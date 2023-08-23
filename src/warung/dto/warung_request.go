package warung
type CreateWarungRequest struct {
	NamaWarung string `gorm:"type:varchar(255)" json:"nama_warung"`
	Gambar     string `gorm:"type:varchar(255)" json:"gambar"`
	Rating     int64  `gorm:"type:int" json:"rating"`
	Lokasi     string `gorm:"type:varchar(255)" json:"lokasi"`
	Deskripsi  string `gorm:"type:text" json:"deskripsi"`
}