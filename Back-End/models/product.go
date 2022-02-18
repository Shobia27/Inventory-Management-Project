package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Id            uint
	Product_Name  string  `json:"productname"`
	Quantity      float64 `json:"quantity"`
	SupplierId    string  `json:"supplierid"`
	Selling_Price float64 `json:"sp"`
}
