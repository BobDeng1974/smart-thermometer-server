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

func main() {
	cfg := utils.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", utils.BasicAuth(controllers.Index, cfg))
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(router)
	log.Fatal(http.ListenAndServe(":8080", n))
}
