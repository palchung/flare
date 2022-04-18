package config

import (
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}

func Setup(app *AppConfig) AppConfig {
	app.InProduction, _ = strconv.ParseBool(os.Getenv("InProduction"))
	app.UseCache, _ = strconv.ParseBool(os.Getenv("UseCache"))
	app.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return *app
}
