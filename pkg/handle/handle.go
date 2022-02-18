package handle

import (
	"encoding/json"
	"github.com/uptrace/bunrouter"
	"io/ioutil"
	"net/http"
	"src/pkg/dto"
	"src/pkg/service"
)

func NewHandle(s service.Service) Handle {
	return Handle{
		service: s,
	}
}

type Handle struct {
	service service.Service
}

func (h Handle) GetTwitters(w http.ResponseWriter, r bunrouter.Request) error {
	response := h.service.GetTweetsService()
	writeResponse(w, 200, response)
	return nil
}

func (h Handle) CreateTweet(w http.ResponseWriter, r bunrouter.Request) error {
	var tweetRequest dto.TweetDto

	err := ReadRequest(r, &tweetRequest)

	response := h.service.CreateTweet(tweetRequest)

	if err != nil {
		writeResponse(w, 400, "deu pau pra parsear")
	}

	writeResponse(w, 201, response)
	return nil
}

func writeResponse(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}

func ReadRequest(r bunrouter.Request, payload interface{}) error {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(body, &payload)

	return nil
}
