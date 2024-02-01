package models

type Item struct {
	ID    string  `json:"itemID"`
	Name  string  `json:"itemName"`
	Price float32 `json:"itemPrice"`
}

type ItemSearch struct {
	Item string `json:"itemSearch"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
