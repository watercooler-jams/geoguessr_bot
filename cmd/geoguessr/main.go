package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"

	geoguessr "watercooler/geoguessr/internal/client"
	"watercooler/geoguessr/internal/model"
	"watercooler/geoguessr/internal/slack"
)

var challengeUrl = "https://www.geoguessr.com/challenge/"

var maps = map[string]model.MapConfig{
	"GEODETECTIVE": {
		ForbidMoving:   true,
		ForbidRotating: false,
		ForbidZooming:  false,
		TimeLimit:      300,
		MapId:          "5d374dc141d2a43c1cd4527b",
		Type:           "MAP",
	},
	"URBAN_WORLD": {
		ForbidMoving:   false,
		ForbidRotating: false,
		ForbidZooming:  false,
		TimeLimit:      120,
		MapId:          "640d01bf1b14982128374759",
		Type:           "MAP",
	},
	"DIVERSE_WORLD": {
		ForbidMoving:   false,
		ForbidRotating: false,
		ForbidZooming:  false,
		TimeLimit:      120,
		MapId:          "603377b9db1db5000118d1d6",
		Type:           "MAP",
	},
	"COUNTRY_STREAK": {
		ForbidMoving:   false,
		ForbidRotating: false,
		ForbidZooming:  false,
		TimeLimit:      120,
		MapId:          "603377b9db1db5000118d1d6",
		Type:           "STREAK",
		StreakType:     "CountryStreak",
	},
}

var mapsPerDay = map[int]string{
	1: "DIVERSE_WORLD",
	2: "URBAN_WORLD",
	3: "COUNTRY_STREAK",
	5: "GEODETECTIVE",
}

func main() {
	args, err := readArgs()
	if err != nil {
		log.Fatal("Failed to parse input parameters: ", err)
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		// error handling
	}
	client := http.Client{
		Timeout: time.Duration(1) * time.Second,
		Jar:     jar,
	}

	_, err = geoguessr.Login(&client, *args.Email, *args.Password)
	if err != nil {
		log.Fatal("Failed to login, skipping creating challenge")
		return
	}

	mapConfig := mapConfigForToday()

	resp, err := geoguessr.CreateChallenge(&client, mapConfig)
	if err != nil {
		log.Fatal("Failed to create challenge", err)
		return
	}

	challengeLink := challengeUrl + resp.Token
	log.Printf("find challenge under: %s\n", challengeLink)

	result, err := slack.PostToSlack(challengeLink, args.SlackHook)
	if err != nil {
		log.Fatal("Failed to send request to slack", err)
	}
	log.Printf("Sent request to slack: %s", result)
}

func mapConfigForToday() *model.MapConfig {
	weekday := int(time.Now().Weekday())
	todaysMap := mapsPerDay[weekday]
	mapConfig := maps[todaysMap]
	return &mapConfig
}

func readArgs() (*model.CmdArgs, error) {
	args := model.CmdArgs{
		Email:     flag.String("ggemail", "", "Geoguessr username"),
		Password:  flag.String("ggpassword", "", "Geoguessr password"),
		SlackHook: flag.String("slackhook", "", "Slack Webhook ID"),
	}
	flag.Parse()
	fmt.Println(*args.SlackHook)
	if args.Email == nil {
		return nil, errors.New("missing GeoGuessr Email!")
	}
	if args.Password == nil {
		return nil, errors.New("missing GeoGuessr Password!")
	}
	if args.SlackHook == nil {
		return nil, errors.New("missing Slack Hook ID!")
	}

	// return &args, errors.New("Stop this")
	return &args, nil
}
