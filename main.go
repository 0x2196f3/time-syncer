package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func currentTimeHandler(w http.ResponseWriter, _ *http.Request) {
	timestamp := time.Now().Unix()
	_, err := fmt.Fprintf(w, "%d", timestamp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	portStr := os.Getenv("PORT")

	var port int
	if portStr == "" {
		port = 30080
		fmt.Println("PORT environment variable is not set, use default port " + strconv.Itoa(port))
	} else {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Error: invalid PORT environment variable:", err)
			return
		}
	}

	http.HandleFunc("/", currentTimeHandler)

	fmt.Printf("Listening on port %d...\n", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error: failed to start server:", err)
		return
	}
}
