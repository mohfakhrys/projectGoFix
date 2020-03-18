package user

import (
	"log"
	"net/http"
)

func (m *Module) HomeUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			log.Println(`error execute template login, err : `, err)
			return
		}
		//} else {
		//	username := r.FormValue("username")
		//	password := r.FormValue("password")
		//
		//	if username == "" {
		//		fmt.Println("missing username")
		//	}
		//	if password == "" {
		//		fmt.Println("missing password")
		//	}
		//
		//	_, err := m.Queries.InsertUser.Exec(username, password)
		//	if err != nil {
		//		log.Println("Failed to insert data")
		//		return
		//	}
		//
		//	log.Println("berhasil insert data")
	}
}
