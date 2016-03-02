package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const steamHost = "http://api.steampowered.com/"
const steamAPIInterface = "ISteamUser/GetPlayerSummaries/v0002/"

type Player struct {
	pr *PlayerResponse
}

func (p *Player) showHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("response in handler: %v\n", p.pr)
}

type (
	// SteamUser is a single player from the Steam API
	SteamUser struct {
		ID       string `json:"steamid"`
		Name     string `json:"personaname"`
		State    int    `json:"personastate"`
		GameID   string `json:"gameid"`
		GameName string `json:"gameextrainfo"`
	}
	// PlayerResponse is a list of players from the Steam API
	PlayerResponse struct {
		Response struct {
			Players []SteamUser `json:"players"`
		} `json:"response"`
	}
)

func main() {
	SetEnvVars()
	steamKey := os.Getenv("STEAM_KEY")
	userIDs := os.Getenv("USER_IDS")
	getUsersURL := steamHost + steamAPIInterface + "?key=" + steamKey + "&steamids=" + userIDs

	r, err := http.Get(getUsersURL)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer r.Body.Close()

	var p PlayerResponse
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Printf("response is: %v\n", &p)

	player := &Player{pr: &p}

	http.HandleFunc("/", player.showHandler)
	http.ListenAndServe(":8001", nil)
}
