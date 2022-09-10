package main

import (
	"net/http"
	"so1/practica2/controllers"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	r := httprouter.New()

	r.GET("/api/cpu/info", controllers.GetCPUOutput)
	r.GET("/api/cpu/usage", controllers.GetCPUUsage)
	r.GET("/api/ram/info", controllers.GetRAMOutput)

	handler := cors.Default().Handler(r)

	http.ListenAndServe(":9000", handler)
}
