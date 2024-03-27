package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/book hitted")
	data := map[string]string{"status": "success", "data": "Success"}
	res, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
