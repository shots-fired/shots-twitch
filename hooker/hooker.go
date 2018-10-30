package hooker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shots-fired/shots-twitch/payloads"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type (
	Hooker interface {
		AddStreamer(name string) error
		AddStreamers(names []string) []error
		RemoveStreamer(name string) error
	}

	hooker struct {
		callbackURL           string
		streamerEncodings     map[string]string
		streamerSubscriptions map[string][]string
	}
)

func NewHooker(callbackURL string) Hooker {
	return hooker{
		callbackURL:           callbackURL,
		streamerEncodings:     make(map[string]string),
		streamerSubscriptions: make(map[string][]string),
	}
}

func (h hooker) AddStreamers(names []string) []error {
	var errors []error
	for _, name := range names {
		errors = append(errors, h.AddStreamer(name))
	}
	return errors
}

func (h hooker) AddStreamer(name string) error {
	client := &http.Client{}

	// Because webhooks require a user_id instead of a user_name
	// we must find these from the api ourselves
	req, _ := http.NewRequest("GET", payloads.UserInfoUrl+"?login="+name, nil)
	req.Header.Add("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))
	resp, _ := client.Do(req)

	out := name + " User Info :> " + resp.Status
	if resp.Status != "200 OK" {
		body, _ := ioutil.ReadAll(resp.Body)
		out += string(body)
	}
	fmt.Println(out)

	var userInfoPayload map[string][]payloads.UserInfo
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&userInfoPayload)
	h.streamerEncodings[name] = userInfoPayload["data"][0].Id

	// Now we need to construct the sub request to the webhook server
	callback := strings.Join([]string{h.callbackURL, "streams", name}, "/")

	out = name + " Callback URL :> " + callback
	fmt.Println(out)

	values := map[string]string{
		"hub.mode":          "subscribe",
		"hub.topic":         payloads.StreamStatusUrl + "?user_id=" + h.streamerEncodings[name],
		"hub.callback":      callback,
		"hub.lease_seconds": "4800",
	}
	valuesJSON, _ := json.Marshal(values)

	req, _ = http.NewRequest("POST", payloads.WebhookUrl, bytes.NewBuffer(valuesJSON))
	req.Header.Add("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))
	req.Header.Add("Content-Type", "application/json")
	resp, _ = client.Do(req)

	out = name + " Webhook Req :> " + resp.Status
	if resp.Status != "202 Accepted" {
		body, _ := ioutil.ReadAll(resp.Body)
		out += string(body)
	}
	fmt.Println(out)
	return nil
}

func (h hooker) RemoveStreamer(name string) error {
	panic("implement me")
	return nil
}
