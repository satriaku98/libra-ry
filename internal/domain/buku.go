package domain

type Buku struct {
	ID          uint   `json:"id_buku" gorm:"primaryKey"`
	Judul       string `json:"judul"`
	Penulis     string `json:"penulis"`
	Penerbit    string `json:"penerbit"`
	TahunTerbit int    `json:"tahun_terbit"`
	Stok        int    `json:"stok"`
}
