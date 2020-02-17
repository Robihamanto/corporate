package model

// Address of users
type Address struct {
	User       int    `json:"user"`
	Street     string `json:"street"`
	PostalCode int    `json:"postalCode"`
}
