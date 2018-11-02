package server

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"testing"
)

type testResponseWriter struct {
	headers http.Header
	Body    io.Reader
	Code    int
}

func newResponseWriter() *testResponseWriter {
	return &testResponseWriter{
		headers: make(http.Header),
	}
}

func (w *testResponseWriter) Header() http.Header {
	return w.headers
}

func (w *testResponseWriter) Write(body []byte) (int, error) {
	w.Body = bytes.NewReader(body)
	return len(body), nil
}

func (w *testResponseWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
}

var testPlayersResponse = []string{
	"Alexis Sánchez; 29; Manchester Utd",
	"Cesc Fàbregas; 31; Chelsea",
	"Dani Ceballos; 22; Real Madrid, Spain",
}

func TestSendResponse(t *testing.T) {
	t.Run("Receive json response successfully", func(t *testing.T) {
		w := newResponseWriter()

		SendResponse(w, testPlayersResponse, http.StatusOK)

		if w.Code != 200 {
			t.Errorf("Expected returned code to be 200 but got %d", w.Code)
		}

		var playerResponse []string
		err := json.NewDecoder(w.Body).Decode(&playerResponse)
		if err != nil {
			t.Errorf("Expected to return err nil but got %v", err)
		}
		if playerResponse[0] != "Alexis Sánchez; 29; Manchester Utd" {
			t.Errorf("Expected string to be Alex... but got %v", playerResponse[0])
		}
		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected header to be aaplication/json but got %v", w.Header().Get("Content-Type"))
		}
	})

	t.Run("Receive internal error in case of unsuccessful marshaling", func(t *testing.T) {
		w := newResponseWriter()
		invalidData := math.Inf(1)

		SendResponse(w, invalidData, http.StatusOK)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected returned code to be 500 but got %d", w.Code)
		}

		bodyBytes, _ := ioutil.ReadAll(w.Body)
		resp := string(bodyBytes)
		if resp != `{"error": "Internal server error"}` {
			t.Errorf("Expected to return error message but got %v", resp)
		}

		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("Expected header to be aplication/json but got %v", w.Header().Get("Content-Type"))
		}
	})
}
