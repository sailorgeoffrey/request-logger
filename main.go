package main

import (
	"os"
	"net/http"
	"log"
	"time"
	"fmt"
	"strings"
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
		fmt.Println(formatRequest(req))
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))
	}
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}
