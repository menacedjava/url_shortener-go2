package main

import (
    "fmt"
    "net/http"
)

var urls = map[string]string{}

func main() {
    http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
        long := r.URL.Query().Get("url")
        key := fmt.Sprintf("u%d", len(urls)+1)
        urls[key] = long
        fmt.Fprintf(w, "Short URL: /r/%s", key)
    })

    http.HandleFunc("/r/", func(w http.ResponseWriter, r *http.Request) {
        key := r.URL.Path[len("/r/"):]
        if long, ok := urls[key]; ok {
            http.Redirect(w, r, long, http.StatusFound)
        } else {
            http.NotFound(w, r)
        }
    })

    http.ListenAndServe(":8080", nil)
}
