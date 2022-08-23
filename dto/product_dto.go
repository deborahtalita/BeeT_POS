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
	Category_id		string	 `json:"category_id"`	
	Subcategory_id	string 	 `json:"subcategory_id"`	
}

type UpdateProductDTO struct {
	// Product_id     string    `json:"product_id"`
	Product_name   string    `json:"product_name"`
	Product_desc   string    `json:"product_desc"`
	Product_update time.Time `json:"product_update"`
}

type AddVariantDTO struct {
	Variant_id     uint64 		`json:"variant_id" binding:"required"`
	Variant_name   string 		`json:"variant_name" binding:"required"`
	Sell_price     uint64 		`json:"sell_price" binding:"required"`
	Purchase_price uint64 		`json:"purchase_price" binding:"required"`
	Stock          uint64 		`json:"stock" binding:"required"`
	Created_time   time.Time	`json:"created_time"`
	Updated_time   time.Time	`json:"updated_time"`
	Product_id		string		`json:"-"`
}

type AddDiscountDTO struct {
	Discount_id		uint64		`json:"discount_id" binding:"required"`
	Discount_name	string		`json:"discount_name" binding:"required"`
	Discount_type	string		`json:"discount_type"`
	Discount_value	float64		`json:"discount_value" binding:"required"`
	Product_id		string		`json:"-"`
}

type AddPictureDTO struct {
	Picture_id   string 	`json:"picture_id" binding:"required"`
	Filename     string 	`json:"filename" binding:"required"`
	Filepath     string 	`json:"filepath" binding:"required"`
	Product_id		string		`json:"-"`
	//Created_time time.Time	`gorm:"type:datetime" json:"created_time"`
}