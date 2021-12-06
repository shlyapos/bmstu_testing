package util

import (
	"encoding/json"
	"net/http"
)

func SetResponse(res http.ResponseWriter, message string, httpStatusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(httpStatusCode)

	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)

	res.Write(jsonResp)
}
