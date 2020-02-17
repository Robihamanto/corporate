package model

// Car has owned by users
type Car struct {
	User  int    `json:"user"`
	Color string `json:"color"`
}
