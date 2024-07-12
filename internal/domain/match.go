package domain

type Match struct {
	ID            string
	TotalKills    int
	Players       map[string]*Player
	Kills         map[string]int
	DeathsByMeans map[string]int
}
