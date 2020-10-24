package wargaming

import (
	"log"
	"testing"
)

func TestGetRealmTournaments(t *testing.T) {
	ts, err := GetRealmTournaments("NA")
	if err != nil {
		t.Fail()
	}

	log.Printf("%+v", ts)
}
