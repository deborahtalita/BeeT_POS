package entity


type Product_picture struct {
	Picture_id   string 	`gorm:"primaryKey;type:varchar(64)" json:"picture_id"`
	Filename     string 	`gorm:"type:varchar(64)" json:"filename"`
	Filepath     string 	`gorm:"type:varchar(64)" json:"filepath"`
	Product_id		string		`gorm:"type:varchar(64)" json:"-"`
	//Created_time time.Time	`gorm:"type:datetime" json:"created_time"`
}