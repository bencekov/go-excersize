package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -build_flags=--mod=mod -package api -destination ./mock_logger.go -source=../../internal/logging/interface.go
//go:generate mockgen -build_flags=--mod=mod -package api -destination ./mock_api.go -source=./interface.go

func TestHandleMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockService := NewMockServiceInterface(ctrl)

	reader := strings.NewReader("test string")
	resultString := "tst strng"

	mockService.EXPECT().CounterAdd().Times(1).Return()
	mockService.EXPECT().RemoveVowels(gomock.Any()).Times(1).Return(resultString, nil)

	req := httptest.NewRequest(http.MethodPost, "/message", reader)
	w := httptest.NewRecorder()

	mux := chi.NewMux()
	NewAPI(mockService, mockLogger).RegisterEndpoints(mux)

	mux.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	resultMessage := new(Message)
	if err := json.Unmarshal(data, resultMessage); err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	assert.Equalf(t, resultString, resultMessage.Message, "Expected %s, got %s", resultString, resultMessage.Message)
}

func TestHandleCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	refCount := 3

	mockLogger := NewMockLoggerInterface(ctrl)
	mockService := NewMockServiceInterface(ctrl)
	mockService.EXPECT().GetCounter().Times(1).Return(refCount)

	req := httptest.NewRequest(http.MethodGet, "/count", nil)
	w := httptest.NewRecorder()

	mux := chi.NewMux()
	NewAPI(mockService, mockLogger).RegisterEndpoints(mux)

	mux.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	resultCount := new(Count)
	if err := json.Unmarshal(data, resultCount); err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}
	assert.Equalf(t, refCount, resultCount.Count, "Expected %v, got %v", refCount, resultCount.Count)
}
