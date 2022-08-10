package entity

import (
	"time"
)

type Customer struct {
	Customer_id string `gorm:"primary_key;type:varchar(64)" json:"customer_id"`
	Customer_name string `gorm:"type:varchar(64)" json:"customer_name"`
	Customer_type string `gorm:"type:varchar(64)" json:"customer_type"`
	Customer_address string `gorm:"type:varchar(512)" json:"customer_address"`
	City_code string  `gorm:"type:varchar(13)" json:"city_code"`
	Province_code string `gorm:"type:varchar(13)" json:"province_code"`
	Customer_phone string `gorm:"type:varchar(64)" json:"customer_phone"`
	Vehicle_type string `gorm:"type:varchar(32)" json:"vehicle_type"`
	Vehicle_brand string `gorm:"type:varchar(64)" json:"vehicle_brand"`
	Vehicle_number string	`gorm:"type:varchar(64)" json:"vehicle_number"`
	Customer_point int32 `gorm:"type:int(11)" json:"customer_point"`
	Customer_update time.Time `gorm:"type:datetime" json:"customer_update"`
	Outlet        Outlet  `gorm:"foreignKey:Outlet_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"outlet"`
	// Outlet_id []Outlet `gorm:"foreignKey:Outlet_id" json:"outlet_id"`
	// Outlet_refer *Outlet `gorm:"foreignKey:Outlet_id;" json:"outlet"`
	// // constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:Outlet_id
}
