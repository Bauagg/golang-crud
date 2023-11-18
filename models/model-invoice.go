package models

type Invoice struct {
	Id                *uint32    `json:"id"`
	IdCart            *uint32    `json:"idCart"`
	Diskon            *uint64    `json:"diskon"`
	IdAddress         *uint32    `json:"idAddress"`
	UserId            *uint32    `json:"userId"`
	PriceTotalInvoice *uint64    `json:"priceTotalInvoice"`
	Carts             Carts      `gorm:"foreignKey:IdCart"`
	Addresses         *Addresses `gorm:"foreignKey:IdAddress"`
	Users             Users      `gorm:"foreignKey:UserId"`
}

type PayloadInvoice struct {
	IdCart    *uint32 `json:"idCart" binding:"required"`
	Diskon    *uint64 `json:"diskon"`
	IdAddress *uint32 `json:"idAddress" binding:"required"`
}
