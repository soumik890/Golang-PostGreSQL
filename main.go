package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golang-crud/Helpers"
	"net/http"
)

func main() {

	fmt.Printf("Server Start")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "welcome home")
	})

	server := http.Server{Addr: "localhost:8080", Handler: routes}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)

}
