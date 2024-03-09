package main

import "net/http"

func handlerError(w http.ResponseWriter, req *http.Request) {
	respondWithError(w, http.StatusNotFound, "Error")
}