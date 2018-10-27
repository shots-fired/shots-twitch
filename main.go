package main

import (
	"github.com/shots-fired/shots-twitch/hooker"
)

func main() {
	var h hooker.Hooker = hooker.NewHooker("8a787t0q4tuqdgyd2fz6cgoug7q853", "http://shotsfired.xyz")

	var streamerNames = []string{"Odega"}
	h.AddStreamers(streamerNames)
}
