package entity

//User represents users table in database

type User struct {
	User_id       	uint64  `gorm:"primary_key" json:"user_id"`
	User_fullname   string  `gorm:"type:varchar(255)" json:"user_fullname"`
	User_name    	string  `gorm:"uniqueIndex;type:varchar(255)" json:"user_name"`
	User_password 	string  `gorm:"type:varchar(255)" json:"-"`
	User_role		string	`json:"user_role"`
	AccessToken    	string  `json:"access_token,omitempty"`
	RefreshToken	string	`json:"refresh_token,omitempty"`
	Outlet_id		string	`json:"outlet_id"`
}