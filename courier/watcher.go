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
		keepWatch chan bool
	}
)

func NewWatcher(userLogin string) Watcher {
	w := watcher{
		userLogin: userLogin,
		keepWatch: make(chan bool, 1),
	}
	go w.watch()
	return w
}

func (w watcher) Start() {
	select {
	case w.keepWatch <- true:
	default:
	}
}

func (w watcher) Stop() {
	select {
	case w.keepWatch <- false:
	default:
	}
}

func (w watcher) watch() {
	keepWatch := <- w.keepWatch
	ticker := time.NewTicker(15 * time.Second)
	for {
		select {
		case keepWatch = <-w.keepWatch:
			return
		default:
		}

		select {
		case <- ticker.C:
			if keepWatch {
				w.RelayStatus()
			}
		default:
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
