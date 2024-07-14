package domain

import "encoding/json"

type Match struct {
	ID         string
	TotalKills int
	Players    map[string]*Player
	Kills      map[string]int
}

type MatchDeathsByMeans struct {
	ID            string
	DeathsByMeans map[string]int
}

func (m *Match) MarshalJSON() ([]byte, error) {
	type Alias Match
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}

func (m *MatchDeathsByMeans) MarshalJSON() ([]byte, error) {
	type Alias MatchDeathsByMeans
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	})
}
