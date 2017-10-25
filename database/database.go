package database

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/JamesHovious/echocrudpoc/models"
)

var GobDB models.Schema

// Encode via Gob to file
func SaveDB(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// Decode Gob file
func LoadDB(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

// QueryUser will return the record for a user given a username.
func QueryUser(user string, db models.Schema) string {
	for _, v := range db.User {
		if v.Username == user {
			ret := fmt.Sprintf("{\"Email\": \"%s\",\"Username\":\"%s\",\"Password\":\"%s\"}", v.Email, v.Username, v.Password)
			return ret
		}
	}
	return "{\"status\":\"no result\"}"
}

// UpdateUser will modify and the record for a user given a username.
func UpdateUser(user models.User, db models.Schema) string {
	for _, v := range db.User {
		if v.Username == user.Username {
			var newUser models.User // Create an empty struct
			// And populate it
			newUser.Password = user.Password
			newUser.Email = user.Email
			newUser.Username = user.Username
			// Noew change the pointer of the struct we want to modify
			*v = newUser
			ret := fmt.Sprintf("{\"Email\": \"%s\",\"Username\":\"%s\",\"Password\":\"%s\"}", v.Email, v.Username, v.Password)
			return ret
		}
	}
	return "{\"status\":\"no result\"}"
}

// DeleteUser will modify and the record for a user given a username.
func DeleteUser(user models.User, db models.Schema) string {
	for _, v := range db.User {
		if v.Username == user.Username {
			var newUser models.User // Create an empty struct
			// Noew change the pointer of the struct we want to modify
			*v = newUser // Note: This isn't the best way to remove an item from a slice.
			// The pointer is still there. It's just an empty struct now.
			ret := fmt.Sprintf("{\"Email\": \"%s\",\"Username\":\"%s\",\"Password\":\"%s\"}", v.Email, v.Username, v.Password)
			return ret
		}
	}
	return "{\"status\":\"no result\"}"
}
