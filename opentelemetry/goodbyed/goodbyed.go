package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		// Print all headers
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(name, value)
			}
		}
		fmt.Fprintf(w, "Goodbye!")
	})
	fmt.Printf("Server running (port=8081), route: http://localhost:8081/goodbye\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		slog.Error("Fatal", err)
	}
}
