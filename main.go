package main

import (
	"fmt"
	"log"
	"net/http"

	"controllers"
	"utils"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func createRouter(cfg utils.Config) (rtr *mux.Router) {
	router := mux.NewRouter().StrictSlash(true)
	router.
		HandleFunc("/", utils.BasicAuth(controllers.Index, cfg))
	router.
		HandleFunc("/addRecord", utils.BasicAuth(controllers.AddRecord, cfg)).
		Methods("POST")
	return router
}

func main() {
	cfg := utils.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	router := createRouter(cfg)
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":8080", n))
}
