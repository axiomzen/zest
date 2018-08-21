package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Running in: %s, revision: %s\n", os.Getenv("ENVIRONMENT"), os.Getenv("REVISION"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})

	fmt.Println("Listening on 2222")
	http.ListenAndServe(":2222", nil)
}
