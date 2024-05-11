package mocks

import (
	"encoding/json"
	"io"
	"net/http"
)

func StartHttpServerMock() {
	mux := http.NewServeMux()

	mux.HandleFunc("/get/alive", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/get/dead", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	mux.HandleFunc("/post/with_data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body map[string]any
			read, _ := io.ReadAll(r.Body)

			json.Unmarshal(read, &body)

			if _, ok := body["data"]; !ok {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})
	mux.HandleFunc("/post/with_data/with_authorization", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body map[string]any
			read, _ := io.ReadAll(r.Body)

			header := r.Header.Get("authorization")

			if header == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			json.Unmarshal(read, &body)

			if _, ok := body["data"]; !ok {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	http.ListenAndServe("localhost:3200", mux)
}
