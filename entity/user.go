package entity

//User represents users table in database
type User struct {
	ID       	uint64  `json:"id"`
	Fullname    string  `gorm:"type:varchar(255)" json:"fullname"`
	Username    string  `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Password 	string  `gorm:"type:varchar(255)" json:"password"`
	User_role		string	`json:"user_role"`
	AccessToken    	string  `json:"access_token,omitempty"`
	RefreshToken	string	`json:"refresh_token,omitempty"`
	Outlet_ID		string	`json:"outlet_id"`
}