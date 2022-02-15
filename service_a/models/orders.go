package models

type Order struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Phone          string `json:"phone_number"`
	Weight         string `json:"parcel_weight"`
	Country        string `json:"country"`
	ServiceBStatus string `json:"service_b_status"`
}
