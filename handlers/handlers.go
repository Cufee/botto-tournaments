package handlers

import (
	wg "botto-tournaments/wargaming"
	"fmt"
	"time"

	"github.com/Necroforger/dgrouter/exrouter"
)

// TournamentsAll - get all tounraments on realm
func TournamentsAll(ctx *exrouter.Context) {
	realm := ctx.Args.Get(1)

	if realm == "" {
		ctx.Reply("Please include a server you want to check.")
		return
	}

	tourneys, err := wg.GetRealmTournaments(realm)
	if err != nil {
		ctx.Reply(fmt.Sprintf("Something did not work:\n```%v```", err))
	}

	var msg string
	for i, t := range tourneys {
		startTime := time.Unix(int64(t.StartTime), 0)
		startStr := startTime.Format("Jan _2 3:04PM")
		msg = msg + fmt.Sprintf("**%v** - %v\n*Start Time: __%v__*", t.Title, t.Status, startStr)
		if i < len(tourneys) {
			msg = msg + "\n\n"
		}
	}

	ctx.Reply(msg)

	return
}
