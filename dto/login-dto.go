package dto

type LoginDTO struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	OutletID	string	`json:"outlet_id"`
}