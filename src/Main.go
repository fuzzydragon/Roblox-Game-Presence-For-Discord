package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"time"
	"regexp"
	"net/http"
	"encoding/json"
	"github.com/hugolgst/rich-go/client"
	//"strconv"
)

var (
	placeId string
	reset = false
)

type MarketPlaceInfo struct { // https://mholt.github.io/json-to-go/
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Creator     struct {
		ID              int    `json:"Id"`
		Name            string `json:"Name"`
		CreatorType     string `json:"CreatorType"`
		CreatorTargetID int    `json:"CreatorTargetId"`
	} `json:"Creator"`
	IconImageAssetID       int64       `json:"IconImageAssetId"`
}

func GetProcessByName(targetProcessName string) *process.Process {
	processes, _ := process.Processes()

	for _, proc := range processes {
		name, _ := proc.Name()
		
		if (name == targetProcessName) {
			return proc
		}
	}

	return nil
}

func GetPlaceInfoByPlaceId(placeId string) *MarketPlaceInfo {
	url := "https://api.roblox.com/marketplace/productinfo?assetId=" + placeId
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	var info *MarketPlaceInfo

	json.NewDecoder(resp.Body).Decode(&info)

	return info
}

func UpdateRobloxPresence() {
	roblox := GetProcessByName("RobloxPlayerBeta.exe")

	for (roblox == nil) {
		roblox = GetProcessByName("RobloxPlayerBeta.exe")

		if (reset == false) {
			reset = true

			client.Logout()
			fmt.Println("reset client activity")
		}
	}

	err := client.Login("823294557155754005")

	if (err != nil) {
		fmt.Println(err)
	}

	reset = false

	args, _ := roblox.Cmdline()

	placePattern := regexp.MustCompile(`placeId=(\d+)`)
	placeMatch := placePattern.FindStringSubmatch(args)[1]

	// timePattern := regexp.MustCompile(`launchtime=(\d+)`)
	// timeMatch := timePattern.FindStringSubmatch(args)[1]

	// startTime, _ := strconv.ParseInt(timeMatch, 10, 64)

	now := time.Now()

	if (placeMatch != placeId) {
		placeId = placeMatch
		place := GetPlaceInfoByPlaceId(placeId)

		client.SetActivity(client.Activity {
			State: "by " + place.Creator.Name,
			Details: place.Name,
			LargeImage: "roblox_logo",
			LargeText: "Playing Roblox!",
			Buttons: []*client.Button {
				&client.Button {
					Label: "Open Game Page",
					Url: "https://www.roblox.com/games/" + placeId + "/-",
				},
			},
			Timestamps: &client.Timestamps {
				Start: &now,
			},
		})

		fmt.Println("set activity: " + place.Name)
		fmt.Println("by: " + place.Creator.Name)
	}
}

func main() {
	for (true) {
		UpdateRobloxPresence()

		time.Sleep(time.Second * 5)
	}
}
