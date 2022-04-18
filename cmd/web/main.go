package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/palchung/flare/internal/config"
	"github.com/palchung/flare/internal/handlers"
	"github.com/palchung/flare/internal/helpers"
	"github.com/palchung/flare/internal/render"
)

var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Fail to load Env file: %s", err)
	}
	port := os.Getenv("PORT")

	app = config.Setup(&app)

	app.TemplateCache, err = render.CreateTemplateCahce()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.Session = NewSession()

	repo := handlers.NewRepo(&app)
	handlers.NewController(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes(&app),
	}
	fmt.Println(fmt.Sprintf("Starting server on port %s", port))

	err = srv.ListenAndServe()
	log.Fatal(err)

}
