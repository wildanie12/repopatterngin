package model

type Product struct {
	Id       int    `gorm:"type:int;primaryKey,not null" json:"id" form:"id" binding:"required"`
	Name     string `gorm:"type:varchar(255);not null" json:"name" form:"name" binding:"required"`
	Kategori string `gorm:"type:varchar(255);not null" json:"kategori" form:"kategori" binding:"required"`
}

type ProductUpdate struct {
	Id       int    `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Kategori string `json:"kategori" form:"kategori" binding:"required"`
}

func (p *ProductUpdate) ToModel() *Product {
	return &Product{
		Id:       p.Id,
		Name:     p.Name,
		Kategori: p.Kategori,
	}
}
