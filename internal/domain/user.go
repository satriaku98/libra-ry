package domain

import "gorm.io/datatypes"

type User struct {
	ID          uint           `gorm:"primaryKey"`
	Username    string         `gorm:"unique;not null"`
	Password    string         `gorm:"not null"`
	Role        string         `gorm:"not null"`                   // Bisa "admin" atau "user"
	Permissions datatypes.JSON `gorm:"type:jsonb;index:,type:gin"` // JSON array of permissions
}
