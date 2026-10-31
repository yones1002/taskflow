package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	serverPort = ":8080"
	welcomeMsg = "سلام! به TaskFlow خوش اومدی"
)

func main() {
	http.HandleFunc("/", homeHandler)

	log.Printf("سرور روی پورت %s شروع شد", serverPort)

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatalf("خطا در اجرای سرور: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, welcomeMsg)
}
