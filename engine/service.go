package engine

import (
	"encoding/json"
	"errors"
	"github.com/ameykpatil/football-engine/utils/helper"
	"io/ioutil"
	"net/http"
	"strconv"
)

var apiPrefix = "https://vintagemonster.onefootball.com/api/teams/en/"
var validTeams = []string{"Germany", "England", "France", "Spain", "Manchester Utd", "Arsenal", "Chelsea", "Barcelona", "Real Madrid", "FC Bayern Munich"}

// TeamAPIResponse response of the team API
type TeamAPIResponse struct {
	Status string  `json:"status"`
	Code   float64 `json:"code"`
	Data   struct {
		Team *Team `json:"team"`
	} `json:"data"`
}

// Team encapsulates properties of football team
type Team struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Players []Player `json:"players"`
}

// Player encapsulates properties of a player
// Teams array would be empty initially & would be filled gradually later
type Player struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Age   string   `json:"age"`
	Teams []string `json:"teams"`
}

// fetch teams & players from vintagemonster api
// check if the team is in the list of valid teams
// if yes return the players & team
func fetchTeamPlayers(teamID int) ([]Player, string, error) {
	// create api url & make a GET call
	apiURL := apiPrefix + strconv.Itoa(teamID) + ".json"
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	// read bytes from the response body
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	if resp == nil || resp.StatusCode != 200 {
		return nil, "", errors.New(string(respBytes))
	}

	// unmarshal bytes into TeamAPIResponse struct
	var teamAPIResponse TeamAPIResponse
	err = json.Unmarshal(respBytes, &teamAPIResponse)
	if err != nil {
		return nil, "", err
	}

	// return nil if the team is not in the list of valid teams
	if !helper.ContainsString(validTeams, teamAPIResponse.Data.Team.Name) {
		return nil, "", errors.New("invalid team")
	}

	return teamAPIResponse.Data.Team.Players, teamAPIResponse.Data.Team.Name, nil
}
