package config

import (
	"database/sql"
	"html/template"
)

var MyApp *App

type App struct {
	DB *sql.DB
	Tpl *template.Template
}