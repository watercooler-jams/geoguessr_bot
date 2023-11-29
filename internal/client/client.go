package geoguessr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"watercooler/geoguessr/internal/model"
	"watercooler/geoguessr/internal/util"
)

var geoguessrUrl = "https://www.geoguessr.com/api/v3"
var signinPath = "/accounts/signin"
var challengePath = "/challenges"
var streakPath = "/challenges/streak"

func Login(client *http.Client, email string, password string) (*model.LoginResponse, error) {
	payload, err := loginRequest(email, password)
	if err != nil {
		return nil, err
	}

	url := geoguessrUrl + signinPath
	body, err := send(client, url, payload)

	response := model.LoginResponse{}
	json.Unmarshal(body, &response)
	return &response, nil
}

func loginRequest(email string, password string) (*[]byte, error) {
	fmt.Printf("data: %s | %s\n", email, password)
	payload, err := json.Marshal(model.LoginRequest{Email: email, Password: password})

	if err != nil {
		fmt.Printf("Error marshalling payload: %s\n", err)
		return nil, err
	}
	return &payload, nil
}

func CreateChallenge(client *http.Client, mapConfig *model.MapConfig) (*model.CreateChallengeResponse, error) {
	request, err := createChallengeRequest(mapConfig)
	if err != nil {
		fmt.Printf("Error creating payload: %s\n", err)
		return nil, err
	}

	// override path for streak challenges
	url := geoguessrUrl + challengePath
	if mapConfig.Type == "STREAK" {
		url = geoguessrUrl + streakPath
	}

	resp, err := send(client, url, request)

	response := model.CreateChallengeResponse{}
	json.Unmarshal(resp, &response)

	return &response, err
}

func createChallengeRequest(mapConfig *model.MapConfig) (*[]byte, error) {
	var mapId string
	if mapConfig.Type == "MAP" {
		mapId = mapConfig.MapId
	}
	var streakType string
	if mapConfig.Type == "STREAK" {
		streakType = mapConfig.StreakType
	}

	payload, err := json.Marshal(model.CreateChallengeRequest{
		ForbidMoving:   mapConfig.ForbidMoving,
		ForbidRotating: mapConfig.ForbidRotating,
		ForbidZooming:  mapConfig.ForbidZooming,
		TimeLimit:      mapConfig.TimeLimit,
		Map:            mapId,
		StreakType:     streakType,
	})

	return &payload, err
}

func send(client *http.Client, url string, payload *[]byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(*payload))
	req.Header.Set("User-Agent", "HTTPie/3.2.1")
	req.Header.Set("Content-Type", "application/json")

	util.PrintRequest(req)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error POSTing data: %s\n", err)
		return nil, err
	}

	util.PrintResponse(resp)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return nil, err
	}
	return body, nil
}
