package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Running in:", os.Getenv("ENVIRONMENT"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})

	fmt.Println("Listening on 2222")
	http.ListenAndServe(":2222", nil)
}
