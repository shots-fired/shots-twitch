package main

import (
	"github.com/shots-fired/shots-twitch/courier"
	"github.com/shots-fired/shots-twitch/hooker"
	"log"
	"net/http"
)

func main() {
	var callbackURL = "http://shotsfired.xyz:3375/webhooks"

	h := hooker.NewHooker(
		callbackURL,
	)

	var streamerNames = []string{"Odega"}
	go h.AddStreamers(streamerNames)

	log.Fatal(http.ListenAndServe(":3375", courier.NewRouter()))
}
