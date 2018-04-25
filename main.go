package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Simple Webhook</h1>")
}

func webhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL.String())
	fmt.Println("Header:", r.Header)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("body:", string(b))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/webhook", webhook)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
