package dto

type CreateOutlet struct {
	Outlet_id      string `json:"outlet_id" form:"outlet_id" binding:"required"`
	Outlet_number  string `json:"outlet_number" form:"outlet_number" binding:"required"`
	Outlet_name    string `json:"outlet_name" form:"outlet_name" binding:"required"`
	Outlet_phone   string `json:"outlet_phone" form:"outlet_phone" binidng:"required"`
	Outlet_address string `json:"outlet_address" form:"outlet_address" binding:"required"`
	City_code      string `json:"city_code" form:"city_code" binding:"required"`
	City_name      string `json:"city_name" form:"city_name" binding:"required"`
	// Province_code  string `json:"province_code" form:"province_code" binding:"required"`
	Province_name  string `json:"province_name" form:"province_name" binding:"required"`
	// Outlet_manager string `json:"outlet_manager" form:"outlet_manager" binding:"required"`
	Outlet_link    string `json:"outlet_link" form:"outlet_link"`
	Outlet_ig      string `json:"outlet_ig" form:"outlet_ig"`
	Outlet_type    string `json:"outlet_type" form:"outlet_type"`
	Outlet_status  bool   `json:"outlet_status" form:"outlet_status"`
	Outlet_created string `json:"outlet_created"`
	Outlet_update  string `json:"outlet_update"` 
	User_id		   uint64 `json:"user_id" form:"user_id"`	

}


type UpdateOutlet struct {
	Outlet_id      string `json:"outlet_id" form:"outlet_id" binding:"required"`
	Outlet_number  string `json:"outlet_number" form:"outlet_number" binding:"required"`
	Outlet_name    string `json:"outlet_name" form:"outlet_name" binding:"required"`
	Outlet_phone   string `json:"outlet_phone" form:"outlet_phone" binidng:"required"`
	Outlet_address string `json:"outlet_address" form:"outlet_address" binding:"required"`
	City_code      string `json:"city_code" form:"city_code" binding:"required"`
	City_name      string `json:"city_name" form:"city_name" binding:"required"`
	// Province_code  string `json:"province_code" form:"province_code" binding:"required"`
	Province_name  string `json:"province_name" form:"province_name" binding:"required"`
	// Outlet_manager string `json:"outlet_manager" form:"outlet_manager" binding:"required"`
	Outlet_link    string `json:"outlet_link" form:"outlet_link"`
	Outlet_ig      string `json:"outlet_ig" form:"outlet_ig"`
	Outlet_type    string `json:"outlet_type" form:"outlet_type"`
	Outlet_status  bool `json:"outlet_status" form:"outlet_status"`
	Outlet_created string `json:"outlet_created"`
	Outlet_update  string `json:"outlet_update"` 
	User_id 	   uint64 `json:"user_id" form:"user_id"`
}

