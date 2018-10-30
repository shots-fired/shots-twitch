package courier

import (
	"github.com/shots-fired/shots-twitch/payloads"
	"net/http"
	"os"
	"time"
)

type (
	Watcher interface {
		Start()
		Stop()
		watch()
		RelayStatus()
	}

	watcher struct {
		userLogin string
		keepWatch bool
	}
)

func NewWatcher(userLogin string) Watcher {
	w := watcher{
		userLogin: userLogin,
		keepWatch: true,
	}
	go w.watch()
	return w
}

func (w watcher) Start() {
	w.keepWatch = true

}

func (w watcher) Stop() {
	w.keepWatch = false
}

func (w watcher) watch() {
	for w.keepWatch {
		for range time.Tick(15 * time.Second) {
			w.RelayStatus()
		}
	}
}

func (w watcher) RelayStatus() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", payloads.StreamStatusUrl+"?user_login="+w.userLogin, nil)
	req.Header.Add("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))
	resp, _ := client.Do(req)

	http.Post("http://shotsfired.xyz:3375/webhooks/streams/" + w.userLogin, "application/json", resp.Body)
}
