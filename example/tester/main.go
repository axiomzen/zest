package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting integration tests")

	resp, err := http.Get("http://" + os.Getenv("TARGET") + ":2222")
	if err != nil {
		fmt.Println("GET errored")
		fmt.Println(err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Println("GET returend code:", resp.Status)
		os.Exit(1)
	}
}
