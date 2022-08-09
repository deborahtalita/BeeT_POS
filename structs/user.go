package structs

import "time"

//User struct
type User struct {
	User_id 			string `gorm:"primary_key" json:"user_id"`
	User_name 			string `gorm:"type:varchar(64)" json:"user_name"`
	User_fullname 		string `gorm:"type:varchar(64)" json:"user_fullname"`
	User_email 			string `gorm:"uniqueIndex;type:varchar(255)" json:"user_email"`
	User_password 		string `gorm:"->;<-;not null" json:"-"`
	User_role			string `gorm:"type:varchar(32)" json:"user_role"`
	User_status 		string `gorm:"type:varchar(32)" json:"user_status"`
	User_update 		time.Time `gorm:"type:datetime" json:"user_update"`
	// Token 		string `gorm:"-" json:"token,omitempty"`
}
