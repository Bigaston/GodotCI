package main

type Webhook struct {
	Content   string         `json:"content"`
	Username  string         `json:"username"`
	AvatarUrl string         `json:"avatar_url"`
	Embeds    []WebhookEmbed `json:"embeds"`
}

type WebhookEmbed struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Url         string `json:"url"`
	// Timestamp   time.Time `json:"timestamp"`
	Color int `json:"color"`
}
