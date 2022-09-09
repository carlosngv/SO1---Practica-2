package main

import (
	"net/http"
	"so1/practica2/controllers"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	r := httprouter.New()

	r.GET("/api/cpuinfo", controllers.GetCPUOutput)
	r.GET("/api/raminfo", controllers.GetRAMOutput)

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":9000", handler)
}
