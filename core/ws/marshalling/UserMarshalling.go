package marshalling

import (
	"encoding/json"

	. "github.com/guidomantilla/bookstore_users-api/core/model"
)

type PublicUser struct {
	Id     int64  `json:"id"`
	Date   string `json:"date"`
	Status string `json:"status"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}

func MarshallUsers(users *[]User, isPublic bool) []interface{} {
	result := make([]interface{}, len(*users))
	for index, user := range *users {
		result[index] = MarshallUser(&user, isPublic)
	}
	return result
}

func MarshallUser(user *User, isPublic bool) interface{} {

	userJson, _ := json.Marshal(user)

	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	}

	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
