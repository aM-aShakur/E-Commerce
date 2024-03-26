package handlers

import (
	"crypto/sha512"
	"db"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"models"
	"net/http"
	"strconv"
	"time"

	s "strings"

	"github.com/lib/pq"
)

// send json response by passing a status code, responseWriter and a struct
func SendJSONResponse(statusCode int, w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

// used to create a base64 id
func CreateID(data string) string {
	/*
	* encode the data with base64 encoding
	* and if the length of the encoded ID
	* is higher than 64, then just use the
	* first 64 characters
	 */
	id := b64.StdEncoding.EncodeToString([]byte(data))
	if len(id) > 64 {
		id = id[:64]
	}

	return id
}

// used to decode a base64 id
func DecodeID(data string) string {
	//decoding base64 encoded string
	decodedID, _ := b64.StdEncoding.DecodeString(data)
	return string(decodedID)
}

// convert each byte of the hash into hexadecimal format using sha512
func GetHash(data string) string {
	/*
	* []byte(data) converts the string variable
	* into a byte array so it's usable in the
	* Sum256 function

	* strconv.FormatInt(int64(hash[i]), 16) takes
	* the current byte of the hash and first
	* converts it to int64 so then it can be
	* converted into a hexadecimal formatted number
	 */
	hash := sha512.Sum512([]byte(data))
	var newHash string = ""
	for i := 0; i < len(hash); i++ {
		newHash += strconv.FormatInt(int64(hash[i]), 16)
	}
	return "0x" + newHash
}

// route: /search
func SearchItem(w http.ResponseWriter, r *http.Request) {
	//read response body (post request data) bytes
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Can't read request body:", err)
	}

	//create post request struct instance
	var itemSearch models.ItemSearch
	//get post request data into loginData struct
	err = json.Unmarshal(body, &itemSearch)

	if err != nil {
		fmt.Println("Unmarshal error:", err)
	} else {
		//use post request data
		fmt.Printf("Item: %s\n", itemSearch.Item)
	}

	//item struct for json response
	//later will come from database query
	item := models.Item{ID: "1", Name: "laptop", Price: 49999}

	SendJSONResponse(200, w, item)
}

// route: /register
func RegisterAccount(w http.ResponseWriter, r *http.Request) {
	//read response body (post request data) bytes
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "Can't read request body")
	}

	//create post request struct instance
	var loginData models.UserLogin
	//get post request data into loginData struct
	err = json.Unmarshal(body, &loginData)

	if err != nil {
		fmt.Println("Unmarshal error:", err)
	} else {
		//connect to database
		db := db.GetDBConnection()

		//get username and password from post request
		username := loginData.Username
		password := loginData.Password

		//hash the password
		password = GetHash(password)

		//format query
		query := `insert into users (id, username, password) values ($1, $2, $3)`

		//data used to create a base64 id
		data := username + fmt.Sprintln(time.Now().UnixMilli())
		data = CreateID(data)

		//attempt query
		_, err := db.Exec(query, data, username, password)

		if err != nil {
			fmt.Println("Failed to execute query", err)
			SendJSONResponse(500, w, nil)
		} else {
			fmt.Println("Created user account successfully")
		}
	}
	loginData.Password = ""
	SendJSONResponse(200, w, loginData)
}

// route: /login
func Login(w http.ResponseWriter, r *http.Request) {
	//read response body (post request data) bytes
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "Can't read request body")
	}

	//create post request struct instance
	var loginData models.UserLogin
	//get post request data into loginData struct
	err = json.Unmarshal(body, &loginData)

	//create user struct variable
	var user models.User

	//variables to store post request data for db query
	var username, password string

	if err != nil {
		fmt.Println("Unmarshal error:", err)
	} else {
		//connect to database
		db := db.GetDBConnection()

		/*
		* get username and password from post request
		* and convert them into valid postgres strings
		* for database queries
		 */

		username = pq.QuoteLiteral(loginData.Username)
		password = pq.QuoteLiteral(GetHash(loginData.Password))

		//format query
		query := fmt.Sprintf(`select id, username from users where username = %s and password = %s`, username, password)

		//attempt db query
		err := db.QueryRow(query).Scan(&user.ID, &user.Username)
		if err != nil {
			//send 500 status code to frontend for query error
			fmt.Println("Failed to execute query", err)
			SendJSONResponse(500, w, nil)
		} else {
			fmt.Println("Logged in user account successfully")
		}
	}
	SendJSONResponse(200, w, user)
}

// route: /item?id=ITEM_ID
func GetItemFromID(w http.ResponseWriter, r *http.Request) {
	//read the item id from url
	id := r.URL.Query().Get("id")

	//using item model
	var item models.Item

	//get db connection
	db := db.GetDBConnection()

	//convert string id to sql compatible string
	id = pq.QuoteLiteral(id)

	//format query
	query := fmt.Sprintf(`select * from items where id = %s`, id)

	//attempt db query (using db.QueryRow)
	err := db.QueryRow(query).Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.AverageRating, &item.Condition, &item.AmountInStock, &item.URLValue)
	if err != nil {
		fmt.Println("Failed to execute query", err)
	}

	//send item to frontend
	SendJSONResponse(200, w, item)
}

// route: /item?name=ITEM_NAME
func GetItemFromName(w http.ResponseWriter, r *http.Request) {
	//read the item name from url
	itemName := r.URL.Query().Get("name")

	//using item model
	var item models.Item

	//get db connection
	db := db.GetDBConnection()

	var urlItemName []rune

	/*
	* creates new formated string incase
	* there are whitespaces, in which those
	* are changed to "-" instead as well
	* as appeneding the rest of the characters
	* lower cased
	 */

	for i := 0; i < len(itemName); i++ {
		if itemName[i] == ' ' {
			urlItemName = append(urlItemName, '-')
		} else {
			//add rest of characters but lowercased

			/*
			* s.ToLower(string(itemName[i])) turns the current character
			* into a string to lowercase it

			* Then, turn the first character of that string back into
			* a rune (character) to append it to rune array (character array)
			* to create new formatted string for item name
			 */
			urlItemName = append(urlItemName, rune(s.ToLower(string(itemName[i]))[0]))
		}
	}

	//convert urlItemName to string
	itemName = string(urlItemName)

	//convert newly formatted string newItem to sql compatible string
	itemName = pq.QuoteLiteral(itemName)

	//format query
	query := fmt.Sprintf(`select * from items where urlvalue = %s`, itemName)

	//attempt db query (using db.QueryRow)
	err := db.QueryRow(query).Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.AverageRating, &item.Condition, &item.AmountInStock, &item.URLValue)
	if err != nil {
		fmt.Println("Failed to execute query", err)
	}

	//send item to frontend
	SendJSONResponse(200, w, item)
}

// route: /items
func GetItems(w http.ResponseWriter, r *http.Request) {
	//using item model
	var item models.Item

	//array of items
	var items = make([]models.Item, 0)

	//get db connection
	db := db.GetDBConnection()

	//attempt db query
	rows, err := db.Query(`select * from items`)
	if err != nil {
		fmt.Println("Failed to execute query", err)
	}

	//get the values from queried rows
	for rows.Next() {
		rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description, &item.AverageRating, &item.Condition, &item.AmountInStock, &item.URLValue)
		items = append(items, item)
	}

	//return array of items to frontend
	SendJSONResponse(200, w, items)
}
