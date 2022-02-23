package server

import "net/http"

func init() {
	http.HandleFunc("/", makeHandler(homeHandler))

	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("download/"))))
}
