package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func JsonResponse(w http.ResponseWriter, status int, message string, data any, errorMsg any) {
	w.Header().Set("Content-Type", "application/json")

	responseData := Response{status < 400, message, data, errorMsg}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		log.Println("Error converting to json:", err)
		w.WriteHeader(500)
		w.Write([]byte("{'error': 'an error occurred!'}"))
		return
	}

	w.WriteHeader(status)
	w.Write([]byte(jsonData))
}

func Error(w http.ResponseWriter, status int, data any) {
	if message, ok := data.(string); ok {
		JsonResponse(w, status, message, nil, nil)
		return
	}
	JsonResponse(w, status, "", nil, data)
}

func Success(w http.ResponseWriter, data any) {
	if message, ok := data.(string); ok {
		JsonResponse(w, 200, message, nil, nil)
		return
	}
	JsonResponse(w, 200, "", data, nil)
}
