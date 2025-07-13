package handlers

import (
	"github.com/usmaarn/locstique_api/packages/response"
	"net/http"
)

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "app running fine.")
}
