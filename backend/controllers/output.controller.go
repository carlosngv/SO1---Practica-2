package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os/exec"
	"so1/practica2/models"
)

func GetOutput(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)

	newOutput := models.Output{}
	
	cmd := exec.Command("sh", "-c", "cat /proc/p2_module")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	newOutput.Output = output
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(newOutput)
	fmt.Fprintf(w, "%s\n", json)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
