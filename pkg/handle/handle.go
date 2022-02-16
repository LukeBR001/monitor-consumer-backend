package handle

import (
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"net/http"
	"src/pkg/service"
)

func GetTwitters(w http.ResponseWriter, r bunrouter.Request) error {
	response := service.GetTweetsService()
	writeResponse(w, 200, response)
	return nil
}

func writeResponse(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)

}
