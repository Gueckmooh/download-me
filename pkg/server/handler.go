package server

import (
	"download-me/pkg/utils/colors"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(article/[A-Za-z0-9-]*||articles/)$") // nolint:lll

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			log.Printf("%sGot request for page %s%s denied%s",
				colors.ColorOrange, string(r.URL.Path), colors.ColorRed, colors.ColorReset)
			http.NotFound(w, r)
			return
		}
		log.Printf("%sGot request for page %s%s accepted%s",
			colors.ColorOrange, string(r.URL.Path), colors.ColorGreen, colors.ColorReset)
		log.Printf("-- %sServe page %s%s\n", colors.ColorGreen, string(r.URL.Path), colors.ColorReset)
		fn(w, r)
	}
}
