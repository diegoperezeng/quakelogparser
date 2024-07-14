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
			matchID := fmt.Sprintf("game_%02d", len(matches)+1)
			currentMatch = &domain.Match{
				ID:      matchID,
				Players: make(map[string]*domain.Player),
				Kills:   make(map[string]int),
			}

			currentMatchDBM = &domain.MatchDeathsByMeans{
				ID:            matchID,
				DeathsByMeans: make(map[string]int),
			}

			matches = append(matches, currentMatch)
			matchesDBM = append(matchesDBM, currentMatchDBM)
		}

		if strings.Contains(line, "Kill") {
			parts := strings.Split(line, ": ")
			if len(parts) < 3 {
				continue
			}
			killInfo := parts[2]
			killData := strings.Fields(killInfo)
			if len(killData) != 5 {
				continue
			}
			killerName := killData[0]
			killedName := killData[2]
			modID, err := domain.ParseDeathCause(killData[4])
			if err != nil {
				continue
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
	}

	return matches, matchesDBM, nil
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
