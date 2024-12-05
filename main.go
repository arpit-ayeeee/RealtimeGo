package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":3000", "Address")

func main() {
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{
		"this", "is", "a", "stream"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s \n\n", string(token))

		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		time.Sleep(time.Millisecond * 420)
	}
}
