package hooker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const WebhookUrl = "https://api.twitch.tv/helix/webhooks/hub"
const StreamStatusUrl = "https://api.twitch.tv/helix/streams"
const UserInfoUrl = "https://api.twitch.tv/helix/users"

type (
	Hooker interface {
		AddStreamer(name string) error
		AddStreamers(names []string) []error
		RemoveStreamer(name string) error
	}

	hooker struct {
		clientID              string
		callbackURL           string
		streamerEncodings     map[string]string
		streamerSubscriptions map[string][]string
	}
)

func NewHooker(clientID string, callbackURL string) Hooker {
	return hooker{
		clientID:              clientID,
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
	// Because webhooks require a user_id instead of a user_name we must find these from the api ourselves
	getUserInfoResp, getUserInfoErr := http.Get(UserInfoUrl + "?login=" + name)
	if getUserInfoErr != nil {
		return getUserInfoErr
	}
	fmt.Println(getUserInfoResp)

	// Now we need to construct the sub request to the webhook server
	values := map[string]string{
		"hub.topic":         StreamStatusUrl + "?user_id=" + h.streamerEncodings[name],
		"hub.callback":      h.callbackURL,
		"hub.lease_seconds": "600",
	}
	valuesJSON, _ := json.Marshal(values)
	postWebhookResp, postWebhookErr:= http.Post(WebhookUrl, "application/json", bytes.NewBuffer(valuesJSON))
	if postWebhookErr != nil {
		return postWebhookErr
	}
	fmt.Println(name + ":> " + string(postWebhookResp.StatusCode))
	return nil
}

func (h hooker) RemoveStreamer(name string) error {
	panic("implement me")
	return nil
}
