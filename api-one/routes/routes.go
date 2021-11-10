package routes

import (
	"encoding/json" // for encoding/decoding JSON
	"net/http"
)

// defining Post entity
type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Creating dummy database
var (
	posts []Post // slice (dynamically sized array)
)

// initializing the database with a post
func init() {
	posts = []Post{Post{
		Id:    1,
		Title: "Title 1",
		Text:  "Text 1",
	}}
}
func GetPosts(res http.ResponseWriter, req *http.Request) {
	
	// We generally interact with api's in JSON
	res.Header().Set("Content-type", "application/json")
	
	// returns the json encoding of posts
	result, err := json.Marshal(posts)
	
	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError) // status: 500
		res.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	
	res.WriteHeader(http.StatusOK) // status: 200
	res.Write(result)
}
func AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	
	// declaring new post of type Post
	var post Post
	
	// reads the JSON value and decodes it into a Go value
	err := json.NewDecoder(req.Body).Decode(&post) 
	
	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	
	// fake ID for the post
	post.Id = len(posts) + 1
	
	// appending the post at the end of dummy array
	posts = append(posts, post)
	res.WriteHeader(http.StatusOK)
	
	// returns the json encoding of post
	result, err := json.Marshal(post)
	res.Write(result)
}
