package mocks

import "net/http"

func StartWebhooksMock() {
	mux := http.NewServeMux()

	mux.HandleFunc("/slack", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/discord", func(w http.ResponseWriter, r *http.Request) {
	})
	mux.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
	})
	mux.HandleFunc("/gws", func(w http.ResponseWriter, r *http.Request) {
	})

	go http.ListenAndServe(":3202", mux)
}
