package models

//represents items in the db
type Item struct {
	ID            string  `json:"itemID"`
	Name          string  `json:"itemName"`
	Price         int     `json:"itemPrice"`
	Description   string  `json:"itemDescription"`
	AverageRating float32 `json:"itemAverageRating"`
	Condition     string  `json:"itemCondition"`
	AmountInStock int     `json:"itemAmountInStock"`
	URLValue      string  `json:"itemURLValue"`
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
