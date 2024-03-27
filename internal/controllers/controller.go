package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/book hitted")
	res, _ := json.Marshal(`{status: "success", data: "Success"}`)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
