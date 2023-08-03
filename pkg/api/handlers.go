package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bencekov/go-exercise/internal/logging"
	"github.com/go-chi/chi/v5"
)

type Message struct {
	Message string `json:"message"`
}

type Count struct {
	Count int `json:"count"`
}

type API struct {
	service ServiceInterface
	logger  logging.LoggerInterface
}

func NewAPI(service ServiceInterface, logger logging.LoggerInterface) *API {
	a := new(API)

	a.service = service
	a.logger = logger
	return a
}

func (a *API) RegisterEndpoints(mux *chi.Mux) {
	mux.Post("/message", a.handleMessage)
	mux.Get("/count", a.handleCount)
}

func (a *API) handleMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	input, err := io.ReadAll(r.Body)
	if err != nil {
		a.logger.Errorf("Error: %s", err.Error())
	}
	messageString, err := a.service.RemoveVowels(string(input))
	if err != nil {
		a.logger.Errorf("Error: %s", err.Error())
	}
	message := Message{
		Message: messageString,
	}
	json.NewEncoder(w).Encode(message)
}

func (a *API) handleCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	count := Count{
		Count: a.service.GetCounter(),
	}

	json.NewEncoder(w).Encode(count)
}
