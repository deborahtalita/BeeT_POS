package entity

type Outlet struct {
	Outlet_id      string    `gorm:"primary_key;type:varchar(64)" json:"outlet_id"`
	// Outlet_number  string    `gorm:"type:varchar(8)" json:"outlet_number"`
	Outlet_name    string    `gorm:"type:varchar(64)" json:"outlet_name"`
	// Outlet_phone   string    `gorm:"type:varchar(16)" json:"outlet_phone"`
	// Outlet_address string    `gorm:"type:varchar(64)" json:"outlet_address"`
	// City_code      string    `gorm:"type:varchar(13)" json:"city_code"`
	// Province_code  string    `gorm:"type:varchar(13)" json:"province_code"`
	// Outlet_manager string    `gorm:"type:varchar(64)" json:"outlet_manager"`
	// Outlet_link    string    `gorm:"type:varchar(1024)" json:"outlet_link"`
	// Outlet_ig      string    `gorm:"type:varchar(64)" json:"outlet_ig"`
	// Outlet_type    string    `gorm:"type:varchar(4)" json:"outlet_type"`
	// Outlet_status  string    `gorm:"type:varchar(32)" json:"outlet_status"`
	// Outlet_created time.Time `gorm:"type:datetime" json:"outlet_created"`
	// Outlet_update  time.Time `gorm:"type:datetime" json:"outlet_update"`
	//Customers	   []Customer `json:"customers,omitempty"`
	//Customers	   []Customer `gorm:"foreignKey:Outlet_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"outlet"`
	//Products	   []Product  `json:"-"`
}