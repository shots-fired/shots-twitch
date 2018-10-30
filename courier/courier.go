package courier

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shots-fired/shots-twitch/payloads"
	"net/http"
)

var watchers = make(map[string]Watcher)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/webhooks/streams/{login}", handleStreamEvent).Methods("POST")
	return r
}

func handleStreamEvent(writer http.ResponseWriter, req *http.Request) {
	writer.WriteHeader(http.StatusOK)

	userLogin := mux.Vars(req)["login"]

	out := userLogin + " :> Stream Event Received - "
	_, ok := watchers[userLogin]

	if !ok {
		watchers[userLogin] = NewWatcher(userLogin)
	}
	w := watchers[userLogin]

	var streamStatusPayload map[string][]payloads.StreamStatus
	defer req.Body.Close()
	json.NewDecoder(req.Body).Decode(&streamStatusPayload)
	data := streamStatusPayload["data"]

	if len(data) == 0 {
		w.Stop()
		out += "Stream Stopped"
	} else if len(data) == 1{
		w.Start()
		out += data[0].Type
	} else {
		panic( "Single streamer " + userLogin + " returned multiple data")
	}

	fmt.Println(out)

}
