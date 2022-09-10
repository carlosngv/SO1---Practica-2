package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"so1/practica2/models"
	"strings"
	"strconv"
	"reflect"

	"github.com/julienschmidt/httprouter"
)

var totalJiffies int
var totalWorkJiffies int
var totalOverPeriod int
var totalOverWork int
var CPUPercentage float64

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


func GetCPUUsage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)

	newOutput := models.CPUUsage{}

	cmd := exec.Command("sh", "-c", "cat /proc/stat")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	newOutput.CPUUsage = output

	stringList := strings.Split(output, "\n")
	cpuInfo := stringList[0]

	fmt.Printf("CPU Info: %v\n\n", cpuInfo)

	cpuValues := strings.Split(cpuInfo, " ")[2:]
	totalJiffiesAux := 0

	// Global totalJiffies will contain previouse total

	for _, item := range cpuValues {
		// totalCPU += strconv.Atoi(item)
		
		intItem, _ := strconv.Atoi(item)
		fmt.Printf("Value: %v - ", intItem)
		fmt.Printf("Type: %v, ", reflect.TypeOf(intItem))
		totalJiffiesAux += intItem
		
	}

	
	if totalJiffies != 0 {
		// previous totalJiffies - current totalJiffies
		totalOverPeriod = totalJiffiesAux - totalJiffies
	}

	totalJiffies = totalJiffiesAux

	totalWorkJiffiesAux := 0

	for index, item := range cpuValues {
		if index == 2 {
			break
		}
		intItem, _ := strconv.Atoi(item)
		fmt.Printf("Value: %v - ", intItem)
		fmt.Printf("Type: %v, ", reflect.TypeOf(intItem))
		totalWorkJiffiesAux += intItem
		
	}
	if totalWorkJiffies != 0 {
		totalOverWork = totalWorkJiffiesAux - totalWorkJiffies
	}

	totalWorkJiffies = totalWorkJiffiesAux

	
	if totalOverPeriod != 0 {
		CPUPercentage = float64((float64(totalOverWork)/float64(totalOverPeriod))*100)
		s := fmt.Sprintf("%.2f", CPUPercentage)
		newOutput.CPUUsage = s
	}else {
		newOutput.CPUUsage = "0.00"
	}
	

	fmt.Printf("\ntotalJiffies: %v", totalOverPeriod)
	fmt.Printf("\ntotalWorkJiffies: %v", totalOverWork)
	fmt.Printf("\nCurrentPercentage: %v", newOutput.CPUUsage)

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
