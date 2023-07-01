package main

import (
	"fmt"
	"log"
	"net/http"
)

func abc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		if err := r.ParseForm(); r != nil {
			fmt.Fprintf(w, "ParseForm() err : %v", err)
			return
		}
		fmt.Fprintf(w, "post form webSite r.postfrom : %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)

	default:
		fmt.Fprintf(w, "Only Get and Post")
	}
}

func main() {
	http.HandleFunc("/", abc)
	fmt.Printf("Starting Server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
