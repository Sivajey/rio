package main


import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(w, "Hi there, this is a Rio demo from dev using webhook")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
