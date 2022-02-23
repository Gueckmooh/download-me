package main

import (
	"fmt"
	"net/http"

	_ "download-me/pkg/server"
)

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("main: %s\n", err.Error())
		return
	}
}
