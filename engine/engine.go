package engine

import (
	"fmt"
	"github.com/ameykpatil/football-engine/utils/helper"
	"sort"
	"strings"
	"sync"
)

const (
	maxConcurrency = 10
)

// Engine is a struct which encapsulate all the properties related to crawling
type Engine struct {
	teamsCount      int               // number of teams processed till date
	playersMap      map[string]Player // playersMap to store all the players
	playersMapMutex *sync.Mutex       // mutex to avoid concurrent access to players map
}

// NewEngine is a constructor for creating engine instance
func NewEngine() *Engine {
	return &Engine{
		teamsCount:      0,
		playersMap:      map[string]Player{},
		playersMapMutex: &sync.Mutex{},
	}
}

// Start method start the process of fetching the players for the teams
// concurrent fetching of teams happens which is controlled by maxConcurrency value
func (engine *Engine) Start() []string {
	resp := make([]string, 0)
	validTeamCount := len(validTeams)
	teamCounter := 1
	for engine.teamsCount < validTeamCount {
		engine.fetch(teamCounter)
		teamCounter = teamCounter + maxConcurrency
		fmt.Println("completed fetching players ", teamCounter)
	}

	for _, player := range engine.playersMap {
		teamStr := ""
		for _, team := range player.Teams {
			teamStr = teamStr + team + ", "
		}
		teamStr = strings.TrimSuffix(teamStr, ", ")
		str := player.Name + "; " + player.Age + "; " + teamStr
		resp = append(resp, str)
	}

	sort.Strings(resp)
	for _, playerStr := range resp {
		fmt.Println(playerStr)
	}
	return resp
}

// fetch fetches players concurrently through go routines
// go routines are synchronized with main routine using waitGroup
func (engine *Engine) fetch(teamCounter int) {

	// waitGroup to synchronize the spawned routines
	var wg sync.WaitGroup
	wg.Add(maxConcurrency)

	for i := 0; i < maxConcurrency; i++ {

		// start a go routine to fetch the players & add into a playersMap
		go func(teamNumber int) {

			players, teamName, err := fetchTeamPlayers(teamNumber)
			if err != nil {
				fmt.Println("error occurred for fetching players with teamID ", teamNumber, err)
			} else {
				for _, player := range players {
					// mutex to avoid concurrent access to playersMap
					engine.playersMapMutex.Lock()
					if engine.playersMap[player.ID].ID == "" {
						player.Teams = []string{teamName}
						engine.playersMap[player.ID] = player
					} else {
						existingPlayer := engine.playersMap[player.ID]
						existingPlayer.Teams = helper.AppendIfMissingString(existingPlayer.Teams, teamName)
						engine.playersMap[existingPlayer.ID] = existingPlayer
					}
					engine.playersMapMutex.Unlock()
				}
				engine.teamsCount++
			}

			// notify parent routine that work is done
			wg.Done()

		}(teamCounter + i)
	}

	// parent routine is waiting for spawned routines
	// to complete their work & notify as done
	wg.Wait()
}
