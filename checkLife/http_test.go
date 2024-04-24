package checklife

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpCheckLife(t *testing.T) {
	server := setUpTestingService()

	t.Run("proving life by http GET request", func(t *testing.T) {
		service := HttpService{
			url:                server.URL + "/ping",
			method:             "GET",
			client:             server.Client(),
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

	t.Run("proving life by http POST request With headers", func(t *testing.T) {
		body, _ := json.Marshal(map[string]string{
			"hola": "hola",
		})
		headers := map[string]string{
			"with_headers": "true",
		}
		service := HttpService{
			url:                server.URL + "/ping/post/with_headers",
			method:             "POST",
			client:             server.Client(),
			requestBody:        body,
			requestHeaders:     headers,
			expectedStatusCode: 202,
		}

		err := service.CheckLife()

		assert.Nil(t, err, err)

		service.requestHeaders["with_headers"] = "false"

		err = service.CheckLife()

		assert.NotNil(t, err, err)
	})
}

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

	mux.HandleFunc("/ping/post/with_headers", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body map[string]any
			read, _ := io.ReadAll(r.Body)

			header := r.Header.Get("with_headers")

			if header != "true" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			json.Unmarshal(read, &body)

			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	server := httptest.NewTLSServer(mux)

	return server
}
