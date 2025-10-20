package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello bot 🌍")
	})

	fmt.Println("Running bot on port 8080")
	http.ListenAndServe(":8080", nil)
}
