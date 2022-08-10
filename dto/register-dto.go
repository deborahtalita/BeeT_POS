package dto

import (
	"time"
)

type RegisterUserDTO struct {
	User_id     string `json:"user_id" form:"user_id" binding:"required"`
	Username    string `json:"user_name" form:"user_name" binding:"required,email"`
	User_fullname string `json:"user_fullname" form:"user_fullname" binding:"required"`
	User_email string `json:"user_email" form:"user_email" binding:"required"`
	User_password string `json:"user_password" form:"user_password" binding:"required"`
	User_role string `json:"user_role" form:"user_role" binding:"required"`
	User_status string `json:"user_status" form:"user_status" binding:"required"`
	User_update time.Time `json:"user_update"`
}

type RegisterCustomerDTO struct{
	Customer_id     	string `json:"customer_id" form:"customer_id" binding:"required"`
	Customer_name 		string `json:"customer_name" form:"customer_name" binding:"required"`
	Customer_type 		string `json:"customer_type" form:"customer_type" binding:"required"`
	Customer_address 	string `json:"customer_address" form:"customer_address" binding:"required"`
	// City_code 			string `json:"city_code" form:"city_code" binding:"required"`
	// Province_code		string `json:"province_code" form:"province_code" binding:"required"`
	Customer_phone 		string `json:"customer_phone" form:"customer_phone" binding:"required"`
	// Vehicle_type		string `json:"vehicle_type" form:"vehicle_type" binding:"required"`
	// Vehicle_brand 		string `json:"vehicle_brand" form:"vehicle_brand" binding:"required"`
	// Vehicle_number		string `json:"vehicle_number" form:"vehicle_number" binding:"required"`
	// Customer_point 		int32 `json:"customer_point" form:"customer_point" binding:"required"`
	Customer_update		time.Time `json:"customer_update"`
	Outlet_id			string `json:"outlet_id" form:"outlet_id,omitempty"`
}