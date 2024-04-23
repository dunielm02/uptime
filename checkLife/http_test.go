package checklife

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUpTestingService() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/ping/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body map[string]any
			read, _ := io.ReadAll(r.Body)

			json.Unmarshal(read, &body)

			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	server := httptest.NewTLSServer(mux)

	return server
}

func TestHttpCheckLife(t *testing.T) {
	server := setUpTestingService()

	t.Run("proving life by http GET request", func(t *testing.T) {
		service := HttpService{
			url:                server.URL + "/ping",
			method:             "GET",
			client:             server.Client(),
			requestBody:        nil,
			expectedStatusCode: 200,
		}

		err := service.CheckLife()

		assert.Nil(t, err, err)
	})

	t.Run("proving life by http POST request", func(t *testing.T) {
		body, _ := json.Marshal(map[string]string{
			"hola": "hola",
		})
		service := HttpService{
			url:                server.URL + "/ping/post",
			method:             "POST",
			client:             server.Client(),
			requestBody:        body,
			expectedStatusCode: 202,
		}

		err := service.CheckLife()

		assert.Nil(t, err, err)
	})
}
