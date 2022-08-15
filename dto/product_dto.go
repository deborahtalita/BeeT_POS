package dto

import "time"

type AddProductDTO struct {
	Product_id     string    `json:"product_id"`
	Product_name   string    `json:"product_name"`
	Product_desc   string    `json:"product_desc"`
	Product_type   string    `json:"product_type"`
	//Product_status bool		 `json:"product_status"`
	Product_update time.Time `json:"customer_update"`
	Outlet_id		string	 `json:"outlet_id"`	
	Category_id		string	 `json:"category_id"`	
	Subcategory_id	string 	 `json:"subcategory_id"`	
}

type UpdateProductDTO struct {
	Product_id     string    `json:"product_id"`
	Product_name   string    `json:"product_name"`
	Product_desc   string    `json:"product_desc"`
	Product_update time.Time `json:"customer_update"`
}