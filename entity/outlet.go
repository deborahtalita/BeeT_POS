package entity

type Outlet struct {
	Outlet_id	 	string `gorm:"primary_key;type:varchar(64)" json:"outlet_id"`
	Outlet_number 	string `gorm:"type:varchar(8)" json:"outlet_number"`
	Outlet_name 	string `gorm:"type:varchar(64)" json:"outlet_name"`
	Outlet_phone 	string `gorm:"type:varchar(16)" json:"outlet_phone"`
	Outlet_address 	string `gorm:"type:varchar(64)" json:"outlet_address"`
	City_code 		string `gorm:"type:varchar(13)" json:"city_code"`
	City_name 		string `gorm:"type:varchar(64)" json:"city_name"`
	// Province_code 	string `gorm:"type:varchar(13)" json:"province_code"`
	Province_name 	string `gorm:"type:varchar(64)" json:"province_name"`
	User_id			uint64 `gorm:"type:int" json:"user_id"`
	Outlet_link 	string `gorm:"type:varchar(1024)" json:"outlet_link"`
	Outlet_ig 		string `gorm:"type:varchar(64)" json:"outlet_ig"`
	Outlet_type 	string `gorm:"type:varchar(8)" json:"outlet_type"`
	Outlet_status 	bool `gorm:"type:bool;default:true" json:"outlet_status"`
	Outlet_created 	string `gorm:"type:varchar(55)" json:"outlet_created"`
	Outlet_update 	string `gorm:"type:varchar(55)" json:"outlet_update"`
	// User_id			string `gorm:"type:varchar(64)" json:"user_id"`	
	User	   		User  `json:"user"`
}
