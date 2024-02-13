package models

//represents items in the db
type Item struct {
	ID    string  `json:"itemID"`
	Name  string  `json:"itemName"`
	Price float32 `json:"itemPrice"`
}

//represents users in the db
type User struct {
	ID       string `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//represents post request data for items
type ItemSearch struct {
	Item string `json:"itemSearch"`
}

//represents post request data for logging in
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
