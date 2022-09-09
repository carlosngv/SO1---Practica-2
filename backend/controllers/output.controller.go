package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"so1/practica2/models"

	"github.com/julienschmidt/httprouter"
)

func GetCPUOutput(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)

	newOutput := models.CPUOutput{}

	cmd := exec.Command("sh", "-c", "cat /proc/cpu_201801434")
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

func GetRAMOutput(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)

	newOutput := models.RAMOutput{}

	cmd := exec.Command("sh", "-c", "cat /proc/ram_201801434")
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
