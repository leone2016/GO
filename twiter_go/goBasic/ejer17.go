package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":593", nil)
}

func home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./index.html")
}
