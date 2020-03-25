package user

import (
	"fmt"
	"log"
	"net/http"
)

func (m *Module) AddArticlesUser(w http.ResponseWriter, r *http.Request) {
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

func (m *Module) EditArticlesUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := m.Template.ExecuteTemplate(w, "editarticle.html", nil)
		if err != nil {
			log.Println(`error execute template home, err : `, err)
			return
		}
	} else {
		getId := r.FormValue("id")
		getContent := r.FormValue("contents")

		query, err := m.Queries.EditArticles.Exec(getContent,getId)
		if err != nil {
			log.Println("Failed to edit data")
			return
		}

		log.Println("berhasil insert data")

		fmt.Println(query)

		http.Redirect(w, r, "http://localhost:9090/", 303)
	}
}



func (m *Module) RemoveArticlesUser(w http.ResponseWriter, r *http.Request) {
		remove := r.FormValue("id")
		fmt.Println("REMOVE: ", remove)
		query, err := m.Queries.RemoveArticles.Exec(remove)
		if err != nil {
			log.Println("Failed to delete data", err)
			return
		}

		fmt.Println("QUERYYY: ", query)

		log.Println("berhasil delete data")

		http.Redirect(w, r, "http://localhost:9090/", 303)
}


