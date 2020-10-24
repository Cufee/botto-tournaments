package wargaming

type tournamentsRes struct {
	Status string       `json:"status"`
	Error  string       `json:"error"`
	Data   []Tournament `json:"data"`
}

type teamsRes struct {
	Status string `json:"status"`
	Meta   struct {
		Count int `json:"count"`
		Total int `json:"total"`
		Page  int `json:"page"`
	} `json:"meta"`
	Error string `json:"error"`
	Data  []team `json:"data"`
}

// Tournament - tournament data
type Tournament struct {
	Realm        string `json:"-"`
	ID           int    `json:"tournament_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	StartTime    int    `json:"start_at"`
	EndTime      int    `json:"end_at"`
	RegistrStart int    `json:"registration_start_at"`
	RegistrEnd   int    `json:"registration_end_at"`
	Teams        []team
}

type team struct {
	ID           int      `json:"team_id"`
	Status       string   `json:"status"`
	Name         string   `json:"title"`
	ClanID       int      `json:"clan_id"`
	TournamentID int      `json:"tournament_id"`
	Players      []player `json:"players"`
}

type player struct {
	ID   int    `json:"account_id"`
	Name string `json:"name"`
}
