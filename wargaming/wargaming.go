package wargaming

import (
	"botto-tournaments/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// GetRealmTournaments - Get all tournaments on realm
func GetRealmTournaments(realm string) (tourneys []Tournament, err error) {
	// Get api domain
	domain, err := getAPIDomain(realm)
	if err != nil {
		return tourneys, err
	}
	url := domain + config.WGtournamentsAPI

	// Get tournaments
	var res tournamentsRes
	err = getJSON(url, &res)
	if err != nil {
		return tourneys, err
	}

	for _, t := range res.Data {
		t.Realm = realm
		tourneys = append(tourneys, t)
	}

	return tourneys, nil
}

// AddTournamentTeams - Get all tournaments teams from tournament ID
func AddTournamentTeams(tournament Tournament) (Tournament, error) {
	domain, err := getAPIDomain(tournament.Realm)
	if err != nil {
		return tournament, err
	}

	// Get teams
	url := domain + config.WGtournamentTeamsAPI + "&tournament_id=" + strconv.Itoa(tournament.ID)
	var res teamsRes
	err = getJSON(url, &res)
	if err != nil {
		return tournament, err
	}
	teams := res.Data

	pages := res.Meta.Total / res.Meta.Count
	// Get all other pages
	if pages > 1 {
		for p := 2; p <= pages; p++ {
			url := domain + config.WGtournamentTeamsAPI + "&tournament_id=" + strconv.Itoa(tournament.ID) + "&page=" + strconv.Itoa(p)
			var res teamsRes
			err = getJSON(url, &res)
			if err != nil {
				return tournament, err
			}
			teams = append(teams, res.Data...)
		}
	}

	tournament.Teams = teams
	return tournament, nil
}

// HTTP client
var clientHTTP = &http.Client{Timeout: 10 * time.Second}

// Mutex lock for rps counter
var waitGroup sync.WaitGroup
var limiterChan chan int = make(chan int, config.WGAPIrateLimit)

// getFlatJSON -
func getJSON(url string, target interface{}) error {
	// Outgoing rate limiter
	start := time.Now()
	limiterChan <- 1
	defer func() {
		timer := time.Now().Sub(start)
		if timer < (time.Second * 1) {
			time.Sleep((time.Second * 1) - timer)
		}
		<-limiterChan
	}()

	res, err := clientHTTP.Get(url)
	if err != nil || res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %v. error: %s", res.StatusCode, err)
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

// getAPIDomain - Get WG API domain using realm
func getAPIDomain(realm string) (string, error) {
	realm = strings.ToUpper(realm)
	if realm == "NA" {
		return "http://api.wotblitz.com", nil

	} else if realm == "EU" {
		return "http://api.wotblitz.eu", nil

	} else if realm == "RU" {
		return "http://api.wotblitz.ru", nil

	} else if realm == "ASIA" || realm == "AS" {
		return "http://api.wotblitz.asia", nil

	} else {
		message := fmt.Sprintf("realm %s not found", realm)
		return "", errors.New(message)
	}
}
