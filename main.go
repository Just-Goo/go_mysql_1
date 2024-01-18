package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Just-Goo/Go-MySql-1/cmd/config" 
	"github.com/Just-Goo/Go-MySql-1/cmd/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	templ := initTemplates()
	db := initDB()

	config.MyApp = &config.App{DB: db, Tpl: templ}

	defer db.Close()

	r := routes.Router()

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func initTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func initDB() (*sql.DB) {
	var err error
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/apidb")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
