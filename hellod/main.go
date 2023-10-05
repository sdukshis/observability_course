package main

import (
	"os"
	"fmt"
	"github.com/gorilla/handlers"
	"log/slog"
	"net/http"
	"strconv"
)

type Token string

// LogValue implements slog.LogValuer and avoid to revealing the token
func (Token) LogValue() slog.Value {
	return slog.StringValue("TOKEN")
}


func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		t := Token(r.URL.Query().Get("token"))
		if err != nil || id < 1 {
			slog.Error("Wrong id passed", "id", id, "token", t)
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Printf("Server running (port=8080), route: http://localhost:8080/hello\n")
	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)); err != nil {
		slog.Error("Fatal", err)
	}
}
