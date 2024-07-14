package parser

import (
	"fmt"
	"quakelogparser/internal/domain"
	"strings"
)

func ParseLogFile(data string) ([]*domain.Match, []*domain.MatchDeathsByMeans, error) {
	var matches []*domain.Match
	var matchesDBM []*domain.MatchDeathsByMeans
	lines := strings.Split(data, "\n")
	var currentMatch *domain.Match
	var currentMatchDBM *domain.MatchDeathsByMeans

	for _, line := range lines {
		if strings.Contains(line, "InitGame") {
			currentMatch, currentMatchDBM = initializeNewGame(len(matches) + 1)
			matches = append(matches, currentMatch)
			matchesDBM = append(matchesDBM, currentMatchDBM)
		}

		if strings.Contains(line, "Kill") {
			processKillLine(line, currentMatch, currentMatchDBM)
		}
	}

	return matches, matchesDBM, nil
}

func initializeNewGame(matchNumber int) (*domain.Match, *domain.MatchDeathsByMeans) {
	matchID := fmt.Sprintf("game_%02d", matchNumber)
	currentMatch := &domain.Match{
		ID:      matchID,
		Players: make(map[string]*domain.Player),
		Kills:   make(map[string]int),
	}

	currentMatchDBM := &domain.MatchDeathsByMeans{
		ID:            matchID,
		DeathsByMeans: make(map[string]int),
	}

	return currentMatch, currentMatchDBM
}

func processKillLine(line string, currentMatch *domain.Match, currentMatchDBM *domain.MatchDeathsByMeans) {
	parts := strings.Split(line, ": ")
	if len(parts) < 3 {
		return
	}
	killInfo := parts[2]
	killData := strings.Fields(killInfo)
	if len(killData) != 5 {
		return
	}
	killerName := killData[0]
	killedName := killData[2]
	modID, err := domain.ParseDeathCause(killData[4])
	if err != nil {
		return
	}

	killer := resolvePlayer(killerName, currentMatch)
	killed := resolvePlayer(killedName, currentMatch)

	if killer.Name == "<world>" {
		currentMatch.Kills[killed.Name]--
	} else {
		currentMatch.Kills[killer.Name]++
	}

	currentMatchDBM.DeathsByMeans[domain.DeathCause(modID).String()]++
	currentMatch.TotalKills++
}

func resolvePlayer(playerName string, match *domain.Match) *domain.Player {
	if player, exists := match.Players[playerName]; exists {
		return player
	} else if playerName == "<world>" {
		return &domain.Player{Name: playerName}
	}

	player := &domain.Player{Name: playerName}
	match.Players[playerName] = player

	return player
}
