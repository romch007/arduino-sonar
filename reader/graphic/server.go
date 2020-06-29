package graphic

import (
	"log"
	"net/http"
)

func serveHTML() {

	http.Handle("/", http.FileServer(http.Dir("./graphic/public")))
	http.HandleFunc("/ws", wsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
