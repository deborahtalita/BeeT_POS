package dto

import "time"

type AddProductDTO struct {
	Product_id     string    `json:"product_id" binding:"required"`
	Product_name   string    `json:"product_name" binding:"required"`
	Product_desc   string    `json:"product_desc" binding:"required"`
	Product_type   string    `json:"product_type" binding:"required"`
	//Product_status bool		 `json:"product_status"`
	Product_update time.Time `json:"product_update"`
	Outlet_id		string	 `json:"outlet_id"`	
	Category_id		string	 `json:"category_id" binding:"required"`	
	Subcategory_id	string 	 `json:"subcategory_id"`	
}

type UpdateProductDTO struct {
	// Product_id     string    `json:"product_id"`
	Product_name   string    `json:"product_name"`
	Product_desc   string    `json:"product_desc"`
	Product_update time.Time `json:"product_update"`
}

type AddVariantDTO struct {
	Variant_id     uint64 		`gorm:"autoIncrement;primary_key;type:int" json:"variant_id"`
	Variant_name   string 		`gorm:"type:varchar(64)" json:"variant_name" binding:"required"`
	Sell_price     uint64 		`gorm:"type:int" json:"sell_price" binding:"required"`
	Purchase_price uint64 		`gorm:"type:int" json:"purchase_price" binding:"required"`
	Stock          uint64 		`gorm:"type:int" json:"stock" binding:"required"`
	Created_time   time.Time	`gorm:"type:datetime" json:"created_time"`
	Updated_time   time.Time	`gorm:"type:datetime" json:"updated_time"`
	Product_id		string		`gorm:"not null" json:"-"`
}