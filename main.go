package main

import (
	"os"
	"net/http"
	"log"
	"time"
	"io/ioutil"
)

func main() {
	port := os.Getenv("PORT0")
	if port == "" {
		port = "8080"
	}
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        logHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

type logHandler struct {
}

func (logHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/__health" {
		w.WriteHeader(200)
	} else {
		body, _ := ioutil.ReadAll(req.Body)
		log.Print("Request: ", req)
		log.Print("Headers: ", req.Header)
		log.Print("Body: ", string(body))
	}
}
