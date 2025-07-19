package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// setupTestEnv("success")

	if _, err := os.Stat(".woodpecker/output.env"); err == nil {
		godotenv.Load(".woodpecker/output.env")
	}

	discordWebhook, isPresent := os.LookupEnv("PLUGIN_WEBHOOK")

	if !isPresent {
		panic("webhook setting is not present")
	}

	result, isPresent := os.LookupEnv("PLUGIN_RESULT")

	if !isPresent {
		panic("result setting is not present")
	}

	matrix, isMatrixPresent := os.LookupEnv("PLUGIN_MATRIX")

	embed := WebhookEmbed{
		Title:       "",
		Description: formatMessage(),
		Type:        "rich",
		Url:         os.Getenv("CI_PIPELINE_URL"),
		Color:       0x000000,
		// Timestamp:   getDateFromEnv("CI_PIPELINE_CREATED"),
	}

	switch result {
	case "success":
		if isMatrixPresent {
			embed.Title = fmt.Sprintf(":white_check_mark: CI for %s#%s (%s) success", os.Getenv("CI_REPO"), os.Getenv("CI_PIPELINE_NUMBER"), matrix)
		} else {
			embed.Title = fmt.Sprintf(":white_check_mark: CI for %s#%s success", os.Getenv("CI_REPO"), os.Getenv("CI_PIPELINE_NUMBER"))
		}

		embed.Color = 0x2ecc71

	case "failure":
		if isMatrixPresent {
			embed.Title = fmt.Sprintf(":x: CI for %s#%s (%s) failed", os.Getenv("CI_REPO"), os.Getenv("CI_PIPELINE_NUMBER"), matrix)
		} else {
			embed.Title = fmt.Sprintf(":x: CI for %s#%s failed", os.Getenv("CI_REPO"), os.Getenv("CI_PIPELINE_NUMBER"))
		}

		embed.Color = 0xe74c3c
	}

	webhookMessage := Webhook{
		Embeds:    []WebhookEmbed{embed},
		Username:  os.Getenv("CI_REPO"),
		AvatarUrl: "https://avatars.githubusercontent.com/u/84780935?s=200&v=4",
	}

	body, err := json.Marshal(webhookMessage)

	if err != nil {
		panic(err)
	}

	_, err = http.Post(discordWebhook, "application/json", bytes.NewBuffer(body))

	if err != nil {
		panic(err)
	}

	fmt.Println("Webhook Sended!")
}

func getDateFromEnv(env string) time.Time {
	var timestamp int64
	timestamp, err := strconv.ParseInt(os.Getenv(env), 10, 64)

	if err != nil {
		panic(err)
	}

	return time.Unix(timestamp, 0)
}

func formatMessage() string {
	buildUrl, isBuildUrlPresent := os.LookupEnv("PLUGIN_BUILD_URL")

	templateData := MessageTemplate{
		StartedAt:  getDateFromEnv("CI_PIPELINE_CREATED").Format("15:04 02/01/2006"),
		Commit:     os.Getenv("CI_COMMIT_SHA"),
		IsBuildURL: isBuildUrlPresent,
		BuildURL:   buildUrl,
	}

	tmpl, err := template.New("message.tmpl").ParseFS(templates, "templates/message.tmpl")

	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, templateData)

	if err != nil {
		panic(err)
	}

	return buf.String()
}

// func setupTestEnv(status string) {
// 	os.Setenv("PLUGIN_WEBHOOK", "")
// 	os.Setenv("PLUGIN_RESULT", status)

// 	os.Setenv("CI_REPO", "bigaston/superrepo")
// 	os.Setenv("CI_PIPELINE_CREATED", "1722617519")
// 	os.Setenv("CI_COMMIT_SHA", "eba09b46064473a1d345da7abf28b477468e8dbd")
// 	os.Setenv("CI_PIPELINE_URL", "https://ci.example/something")
// }
