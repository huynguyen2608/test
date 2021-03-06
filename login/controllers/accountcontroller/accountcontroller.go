package accountcontroller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/accountcontroller/index.html")
	tmp.Execute(response, nil)
}

func Login(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	sessions.Save(request, response)
	if username == "abc" && password == "123" {
		data := map[string]interface{}{
			"err": "Invalid",
		}
		tmp, _ := template.ParseFiles("view/accountcontroller/index.html")
		tmp.Execute(response, data)

	} else {
		session, _ := store.Get(request, "mysession")
		session.Values["username"] = username
		http.Redirect(response, request, "/account/welcome", http.StatusSeeOther)
	}

}

func Welcome(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	tmp, _ := template.ParseFiles("views/accountcontroller/welcom.html")
	tmp.Execute(response, data)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "mysession")
	session.Options.MaxAge = -1
	session.Save(request, response)
	http.Redirect(response, request, "/account/index", http.StatusSeeOther)
}
