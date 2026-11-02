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
	badRequest          = "بدنه درخواست نامعتبر است"
	methodNotAllowedMsg = "فقط متد GET مجاز است"
	notFound            = "رکورد مورد نظر پیدا نشد"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/task", tasksHandler)
	http.HandleFunc("/task/done", doneHandler)
	http.HandleFunc("/task/count", countHandler)

	log.Printf("سرور روی پورت %s شروع شد", serverPort)

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatalf("خطا در اجرای سرور: %v", err)
	}
}
