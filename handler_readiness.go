package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}
