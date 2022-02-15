package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID      string `json:"id"`
	Email   string `json:"email"`
	Phone   string `json:"phone_number"`
	Weight  string `json:"parcel_weight"`
	Country string `json:"country"`
}
