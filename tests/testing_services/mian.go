package main

import (
	"net/http"
	httpserver "testing_service/http_server"
)

func main() {
	http.ListenAndServe(":8000", httpserver.GetHandler())
}
