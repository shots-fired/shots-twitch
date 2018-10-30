package payloads

const WebhookUrl = "https://api.twitch.tv/helix/webhooks/hub"
const StreamStatusUrl = "https://api.twitch.tv/helix/streams"
const UserInfoUrl = "https://api.twitch.tv/helix/users"

type UserInfo struct {
	Id              string
	Login           string
	DisplayName     string
	Type            string
	BroadcasterType string
	Description     string
	ProfileImageURL string
	OfflineImageURL string
	ViewCount       string
	Email           string
}

type StreamStatus struct {
	Id           string
	UserID       string
	UserName     string
	GameId       string
	CommunityIds []string
	Type         string
	Title        string
	ViewerCount  int
	StartedAt    string
	Language     string
	ThumbnailURL string
}
