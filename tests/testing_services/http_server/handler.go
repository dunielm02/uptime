package httpserver

import "net/http"

func GetHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World from a container"))
	})

	return mux
}
