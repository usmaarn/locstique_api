package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/usmaarn/locstique_api/internal/config"
	"net/http"
)

func ParseBody(r *http.Request, target any) error {
	err := json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		fmt.Println("Error parsing json: ", err)
		return errors.New("invalid request body")
	}
	return config.Validate.Struct(target)
}
