package handlers

import (
	"github.com/usmaarn/locstique_api/packages/response"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response.JsonResponse(w, http.StatusOK, map[string]string{
		"message": "app running fine.",
	})
}
