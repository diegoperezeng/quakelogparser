package model

type Match struct {
	ID         string             `json:"id"`
	TotalKills int                `json:"total_kills"`
	Players    map[string]*Player `json:"players"`
	Kills      map[string]int     `json:"kills"`
}

type MatchDeathsByMeans struct {
	ID            string         `json:"id"`
	DeathsByMeans map[string]int `json:"kills_by_means"`
}
