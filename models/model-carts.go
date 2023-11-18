package models

type Carts struct {
	Id         *uint32   `json:"id"`
	IdProduct  *uint32   `json:"idProduct"`
	Quantity   *uint16   `json:"quantity"`
	PriceTotal *uint64   `json:"priceTotal"`
	UserId     *uint32   `json:"userId"`
	Products   *Products `gorm:"foreignKey:IdProduct"`
	Users      *Users    `gorm:"foreignKey:UserId"`
}

type PayloadCart struct {
	IdProduct uint32 `json:"idProduct" binding:"required"`
	Quantity  uint16 `json:"quantity" binding:"required"`
	UserId    uint32 `json:"userId"`
}
