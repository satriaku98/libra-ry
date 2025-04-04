package domain

type Anggota struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `gorm:"uniqueIndex;not null"` // FK ke users.id
	Nama   string `gorm:"not null"`
	Email  string `gorm:"unique;not null"`
	Alamat string `json:"alamat"`
	NoHP   string `json:"no_hp"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type AnggotaSwagger struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
	NoHP   string `json:"no_hp"`
}
