package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWthJson(w, 200, struct{}{})
}