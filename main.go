package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func sendHTMLServerHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := os.ReadFile("index.html")
	fmt.Fprintf(w, string(page))
}

func consoleLog(msg string, source string) {
	time := fmt.Sprintf("%s", time.Now())
	fmt.Printf("[%s] [%s] %s\n", time[0:19], source, msg)
}

func main() {
	http.HandleFunc("/", sendHTMLServerHandler)
	http.ListenAndServe(":8000", nil)
}
