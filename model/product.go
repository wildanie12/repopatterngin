package model

type Product struct {
	Id       int    `gorm:"type:int;primaryKey,not null" json:"id" binding:"required"`
	Name     string `gorm:"type:varchar(255);not null" json:"name" binding:"required"`
	Kategori string `gorm:"type:varchar(255);not null" json:"kategori" binding:"required"`
}
