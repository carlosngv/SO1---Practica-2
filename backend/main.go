package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"so1/practica2/controllers"
)

func main() {

	r := httprouter.New()

	r.GET("/api/systeminfo", controllers.GetOutput)

	handler := cors.Default().Handler(r)
	
	http.ListenAndServe(":9000", handler)
}
