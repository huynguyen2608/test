package routes

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

type User struct {
	ID   int  `json:"id"`
	Name string `json:"name"`
	Age  int  `json:"age"`
}

var  users []User

func init() {
	users = []User{
		User{
			ID: 1,
			Name: "Huy",
			Age: 25,
		},
		User{
			ID: 1,
			Name: "Huy",
			Age: 25,
		},
		User{
			ID: 1,
			Name: "Huy",
			Age: 25,
		},
		User{
			ID: 1,
			Name: "Huy",
			Age: 25,
		}}
}


func GetUsers(rs http.ResponseWriter, rp *http.Request) {
	rs.Header().Set("Content-type","application/json")
	result, err := json.Marshal((users))

	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError)
		rs.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	rs.WriteHeader(http.StatusOK)
	rs.Write((result))
}

func AddUsers(rs http.ResponseWriter, rp *http.Request) {
	rs.Header().Set("Content-type","application/json")
	var user User

	err :=json.NewDecoder(rp.Body).Decode(&user)
	if err != nil {
		rs.WriteHeader(http.StatusInternalServerError)
		rs.Write([]byte(`{"error": "error unmarshalling"}`))
		return
	}

	user.ID = len(users) + 1

	users = append(users, user)
	rs.WriteHeader(http.StatusOK)

	result, err := json.Marshal(user)
	rs.Write(result)
}