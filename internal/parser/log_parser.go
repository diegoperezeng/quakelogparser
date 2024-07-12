package parser

import (
	"quakelogparser/internal/domain"
	"strconv"
	"strings"
)

func ParseLogFile(data string) ([]*domain.Match, error) {
	matchMap := make(map[string]*domain.Match)
	var matches []*domain.Match
	lines := strings.Split(data, "\n")
	var currentMatch *domain.Match

	for _, line := range lines {
		if strings.Contains(line, "InitGame") {
			currentMatch = &domain.Match{
				ID:            "game_" + strconv.Itoa(len(matches)+1),
				Players:       make(map[string]*domain.Player),
				Kills:         make(map[string]int),
				DeathsByMeans: make(map[string]int),
			}
			matches = append(matches, currentMatch)
			matchMap[currentMatch.ID] = currentMatch
		}

		if strings.Contains(line, "Kill") {
			parts := strings.Split(line, ": ")
			killInfo := parts[1]
			killData := strings.Fields(killInfo)
			killerID := killData[1]
			killedID := killData[2]
			modID, _ := strconv.Atoi(killData[3])

			killer := resolvePlayer(killerID, currentMatch)
			killed := resolvePlayer(killedID, currentMatch)

			if killer.Name == "<world>" {
				currentMatch.Kills[killed.Name]--
			} else {
				currentMatch.Kills[killer.Name]++
			}

			currentMatch.DeathsByMeans[domain.DeathCause(modID).String()]++
			currentMatch.TotalKills++
		}
	}

	return matches, nil
}

func resolvePlayer(playerID string, match *domain.Match) *domain.Player {
	if playerID == "1022" {
		return &domain.Player{Name: "<world>"}
	}

	playerName := "Player_" + playerID
	if player, exists := match.Players[playerName]; exists {
		return player
	}

	player := &domain.Player{Name: playerName}
	match.Players[playerName] = player
	return player
}
