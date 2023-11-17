package models

type Product struct {
	Id          *uint32 `json:"id"`
	NameProduct *string `json:"name_product"`
	Stock       *uint16 `json:"stock"`
	Price       *uint32 `json:"price"`
	Images      *string `json:"images"`
}

type CreateProduct struct {
	NameProduct string `json:"name_product" binding:"required"`
	Stock       uint16 `json:"stock" binding:"required"`
	Price       uint32 `json:"price" binding:"required"`
	Images      string `json:"images"`
}
