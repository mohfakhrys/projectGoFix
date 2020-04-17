package user

import (
	"fmt"
	"log"
	"net/http"
)

func (m *Module) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			log.Println(`error execute template login, err : `, err)
			return
		}
	} else {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" {
			fmt.Println("missing username")
		}
		if password == "" {
			fmt.Println("missing password")
		}

		_, err := m.Queries.LoginUser.Query(username, password)
		if err != nil {
			log.Println("Failed to login")
			return
		}

		log.Println("berhasil login")

		http.Redirect(w, r, "http://localhost:9090/", 303)
	}
}
