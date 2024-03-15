package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter,code int ,msg string){
	if code>499{
		log.Println("Responding with 5xx error:",msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	respondWthJson(w,code,errResponse{
		Error: msg,
	})
}

func respondWthJson(w http.ResponseWriter, code int ,payload interface{}){
	dat,err:=json.Marshal(payload)
	if err!=nil{
		log.Printf("Failed to marshal JSON response: %v",payload)
		w.WriteHeader(code)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
