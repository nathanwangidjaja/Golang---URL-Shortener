package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", Map)
	log.Fatal(http.ListenAndServe("localhost:8086", nil))
	}

func Map(w http.ResponseWriter, r *http.Request) {
	pathMap := make(map[string]string)
	pathMap["/git"] = "https://github.com/nathanwangidjaja"
	pathMap["/go"] = "https://golang.org"
	url := r.URL.Path

	for key, val := range pathMap {
		if url == key {
			http.Redirect(w, r, val, http.StatusSeeOther)
			return
		}
	}

}