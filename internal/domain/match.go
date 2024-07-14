package domain

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
