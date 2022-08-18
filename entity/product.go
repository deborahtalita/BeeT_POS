package entity

// import "time"

type Product struct {
	Product_id		string		`gorm:"primaryKey;type:varchar(64)" json:"product_id"`
	Product_name   	string 	 	`gorm:"type:varchar(64)" json:"product_name"`
	Product_desc   	string 	 	`gorm:"type:varchar(512)" json:"product_desc"`
	Product_type   	string 	 	`gorm:"type:varchar(64)" json:"product_type"`
	Product_status	bool		`gorm:"type:bool;default:true" json:"product_status"`
	//Product_update 	time.Time 	`gorm:"type:datetime" json:"product_update"`
	Outlet_id		string		`gorm:"not null" json:"-"`	
	Category_id		string		`gorm:"type:varchar(64)" json:"category_id"`	
	Subcategory_id	string		`gorm:"type:varchar(64)" json:"subcategory_id"`	
	Outlet			Outlet		`json:"outlet"`
}