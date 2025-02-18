package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 30080
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := fmt.Fprintf(w, "%d", time.Now().Unix())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("listening on port %d...\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("failed to start server: %s\n", err.Error())
	}
}
