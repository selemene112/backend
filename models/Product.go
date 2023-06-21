package models



type Product struct {
	Id         int    `gorm:"primaryKey" json:"id"`
	Name_produk string `gorm:"column:Name_produk;type:varchar(25)" json:"name_produk"`
	Keterangan  string `gorm:"column:Keterangan;type:varchar(25)" json:"Keterangan"`
	Harga       int    `gorm:"type:int" json:"harga"`
	Jumlah      int    `gorm:"type:int" json:"jumlah"`
	
}