package entity

import "time"

type Product_variant struct {
	Variant_id     uint64 		`gorm:"primary_key;type:int" json:"variant_id"`
	Variant_name   string 		`gorm:"type:varchar(64)" json:"variant_name"`
	Sell_price     uint64 		`gorm:"type:int" json:"sell_price"`
	Purchase_price uint64 		`gorm:"type:int" json:"purchase_price"`
	Stock          uint64 		`gorm:"type:int" json:"stock"`
	Created_time   time.Time	`gorm:"type:datetime" json:"created_time"`
	Updated_time   time.Time	`gorm:"type:datetime" json:"updated_time"`
	Product_id		string		`gorm:"type:varchar(64)" json:"product_id"`
}