package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/proxy", func(w http.ResponseWriter, req *http.Request) {
		URL := req.URL.Query().Get("url")

		r, err := http.Get(URL)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("failed to curl url (%s): %v", URL, err)))
			return
		}

		body := r.Body
		defer body.Close()

		content, err := io.ReadAll(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("failed to read response from url (%s): %v", URL, err)))
			return
		}

		w.Write(content)
	})

	log.Println("Ready to accept requests")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
