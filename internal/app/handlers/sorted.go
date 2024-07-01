package handlers

import (
	service "arr-sorted/internal/app/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Sorted(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*") //Обход политики одинакового источника.

	params := r.URL.Query().Get("numbers")

	sortedArr := service.ParamsToArr(params)
	fmt.Println(sortedArr)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Ars", "400")

	data, err := json.Marshal(sortedArr)
	if err != nil {
		log.Print(err)
	}

	w.Write(data)

}
