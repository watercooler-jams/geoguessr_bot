package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"watercooler/geoguessr/internal/model"
	"watercooler/geoguessr/internal/util"
)

var slackUrl = "https://hooks.slack.com/services/"

func PostToSlack(challengeUrl string, slackHook *string) (string, error) {
	payload, err := buildPayload(challengeUrl)
	if err != nil {
		return "", err
	}

	url := slackUrl + *slackHook
	fmt.Printf("Posting to slack webhook: %s\n", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(*payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "HTTPie/3.2.1")
	req.Header.Set("Content-Type", "application/json")

	util.PrintRequest(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	util.PrintResponse(resp)

	return "Success", nil
}

func buildPayload(challengeUrl string) (*[]byte, error) {
	payload, err := json.Marshal(model.SlackMessageRequest{Text: challengeUrl})
	return &payload, err
}
