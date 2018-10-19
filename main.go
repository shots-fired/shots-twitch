package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/shots-fired/shots-twitch/hooker"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	h := hooker.NewHooker("8a787t0q4tuqdgyd2fz6cgoug7q853")

	var streamerNames = []string{"Odega"}


	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":3375", nil))
}
