package user

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"

	_ "github.com/lib/pq"
)

type Module struct {
	Template 	*template.Template
	DB 			*sql.DB
	Queries 	*Queries
}

const (
	host		= "localhost"
	port		= 5432
	user		= "postgres"
	password	= "admin"
	dbname		= "blog"
)

func New(template *template.Template) *Module  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("Failed to connect database, Error", err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error PING", err)

	}

	queries := NewQueries(db)

	return &Module{
		Template: template,
		DB:	db,
		Queries: queries,
	}
}

type Queries struct {
	InsertUser		*sql.Stmt
	LoginUser 		*sql.Stmt
	InsertContents	*sql.Stmt
	SendMessage 	*sql.Stmt
}

func NewQueries(db *sql.DB) *Queries {
	queryInsertUser 	:= `insert into public.users (username, pass) values ($1, $2)`
	queryLoginUser		:= `select * from public.users where username = $1 and pass = $2`
	queryInsertContent 	:= `insert into public.articles (user_id, contents) values (1, $1)`
	querySendMessage	:= `insert into public.pesan (email, message) values ($1, $2)`
	//querySelectData := `select * from userinfo`

	return &Queries{
		InsertUser: prepare(queryInsertUser, db),
		LoginUser: prepare(queryLoginUser, db),
		InsertContents: prepare(queryInsertContent, db),
		SendMessage: prepare(querySendMessage, db),
		//UpdateUser: prepare(queryUpdateUser, db),
	}
}

func prepare(query string, db *sql.DB) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("failed to prepare query : ", err)
	}
	return  stmt
}