package main

import "net/http"

// function signature for http handler that go standard library expects
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type healthMessageResponse struct {
		Message string `json:"message"`
	}
	respondWithJson(w, 200, healthMessageResponse{
		Message: "Server is ready!",
	})
}
