package models

type Addresses struct {
	Id          *uint32 `json:"id"`
	Village     *string `json:"village"`
	Subdistrict *string `json:"subdistrict"`
	City        *string `json:"city"`
	Province    *string `json:"province"`
	Country     *string `json:"country"`
	UserId      *uint32 `json:"userId"`
	Users       *Users  `gorm:"foreignKey:UserId"`
}

type PayloadAddress struct {
	Village     string `json:"village" binding:"required"`
	Subdistrict string `json:"subdistrict" binding:"required"`
	City        string `json:"city" binding:"required"`
	Province    string `json:"province" binding:"required"`
	Country     string `json:"country" binding:"required"`
	UserId      uint32 `json:"userId"`
}
