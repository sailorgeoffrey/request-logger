package main

import (
	"os"
	"fmt"
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
		Addr:           ":8080",
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
	body,_ := ioutil.ReadAll(req.Body)
	fmt.Println(req)
	fmt.Println(string(body))
}
