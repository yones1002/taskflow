package main

import (
	"log"
	"net/http"
)

const (
	serverPort          = ":8080"
	welcomeMsg          = "سلام! به TaskFlow خوش اومدی"
	healthStatus        = "OK"
	appName             = "TaskFlow"
	appVersion          = "0.1.0"
	appStatus           = "running"
	methodNotAllowedMsg = "فقط متد GET مجاز است"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)

	log.Printf("سرور روی پورت %s شروع شد", serverPort)

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatalf("خطا در اجرای سرور: %v", err)
	}
}
