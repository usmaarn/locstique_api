package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData, err := json.Marshal(payload)

	if err != nil {
		log.Println("Error converting to json:", err)
		w.Write([]byte("{'message': 'an error occurred!'}"))
		return
	}
	w.Write([]byte(jsonData))
}
