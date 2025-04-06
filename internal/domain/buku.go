package domain

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Buku struct {
	gorm.Model                 // Menyertakan ID, CreatedAt, UpdatedAt, DeletedAt
	Judul       string         `json:"judul"`
	Penulis     string         `json:"penulis"`
	Penerbit    string         `json:"penerbit"`
	TahunTerbit int            `json:"tahun_terbit"`
	Stok        int            `json:"stok"`
	Tags        datatypes.JSON `gorm:"type:jsonb;index:,type:gin" json:"tags"`
}

// BukuSwagger is a struct for Swagger documentation
type BukuSwagger struct {
	ID          uint     `json:"id_buku"`
	Judul       string   `json:"judul"`
	Penulis     string   `json:"penulis"`
	Penerbit    string   `json:"penerbit"`
	TahunTerbit int      `json:"tahun_terbit"`
	Stok        int      `json:"stok"`
	Tags        []string `json:"tags"`
}
