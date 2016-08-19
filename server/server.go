package main

import (
	//	"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	//	"github.com/gorilla/handlers"
	"net/http"

	rhandlers "github.com/remotejob/real_estate_redux_go/server/handlers"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api", rhandlers.AllRecords)
	gorillaRoute.HandleFunc("/api/{id}", rhandlers.OneRecords)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", c.Handler(gorillaRoute))
}
