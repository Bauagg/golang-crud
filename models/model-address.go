package models

type Address struct {
	Id          *uint32 `json:"id"`
	Village     *string `json:"village"`
	Subdistrict *string `json:"subdistrict"`
	City        *string `json:"city"`
	Province    *string `json:"province"`
	Country     *string `json:"country"`
}

type PayloadAddress struct {
	Village     string `json:"village" binding:"required"`
	Subdistrict string `json:"subdistrict" binding:"required"`
	City        string `json:"city" binding:"required"`
	Province    string `json:"province" binding:"required"`
	Country     string `json:"country" binding:"required"`
}
