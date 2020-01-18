package main

import (
	"net/http"
	"strings"
)

var Map = make(map[string]string)

func Redirect(w http.ResponseWriter, r *http.Request) {
	mainPath := r.URL.Path
	mainPath = strings.TrimPrefix(mainPath, "/")
	w.Write([]byte("URL Shortener for Golang, type a path to be redirected"))	
	if mainPath ==""{
	} else if  Map[mainPath]== "" {
		w.Write([]byte("Error, there path cannot be redirected. Please try another path"))
	} else {
		http.Redirect(w, r, Map[mainPath], 301)
	}
	print(mainPath)
}

func main() {
	pathMap["go"] = "https://golang.org"
	pathMap["git"] = "https://github.com/nathanwangidjaja"

	http.HandleFunc("/", Redirect)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
