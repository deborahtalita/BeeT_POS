package entity

type Product_discount struct {
	Discount_id		uint64		`gorm:"primaryKey;type:int" json:"discount_id"`
	Discount_name	string		`gorm:"type:varchar(64)" json:"discount_name"`
	Discount_type	string		`gorm:"type:varchar(64)" json:"discount_type"`
	Discount_value	float64		`gorm:"type:int" json:"discount_value"`
	Product_id		string		`gorm:"type:varchar(64)" json:"product_id"`
}