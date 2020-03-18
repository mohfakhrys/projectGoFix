package user

import (
	"fmt"
	"log"
	"net/http"
)

func (m *Module) ArticlesUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "addarticle.html", nil)
		if err != nil {
			log.Println(`error execute template home, err : `, err)
			return
		}
	} else {
		contents := r.FormValue("contents")

		if contents == "" {
			fmt.Println("missing contents")
		}

		_, err := m.Queries.InsertContents.Exec(contents)
		if err != nil {
			log.Println("Failed to insert data")
			return
		}

		log.Println("berhasil insert data")

		http.Redirect(w, r, "http://localhost:9090/", 303)
	}
}
