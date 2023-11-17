package models

type CartsTabel struct {
	Id         *uint32 `json:"id"`
	IdProduct  *uint32 `json:"idProduct"`
	Quantity   *uint16 `json:"quantity"`
	PriceTotal *uint64 `json:"priceTotal"`
}

type PayloadCart struct {
	IdProduct uint32 `json:"idProduct" binding:"required"`
	Quantity  uint16 `json:"quantity" binding:"required"`
}
