package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model

	Id                    uint
	SupplierId            string  `json:"supplierid"`
	Supplier_Name         string  `json:"suppliername"`
	Supplier_Address      string  `json:"supplieraddress"`
	Contact_Num           string  `json:"contactnum"`
	Num_Inventory_Product float64 `json:"numproduct"`
}
