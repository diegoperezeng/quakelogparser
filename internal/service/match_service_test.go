package service

import (
	"quakelogparser/internal/domain"
	"testing"
)

func TestGroupMatches(t *testing.T) {
	matches := []*domain.Match{
		{
			ID:         "game_1",
			TotalKills: 2,
			Players: map[string]*domain.Player{
				"HarryPotter": {Name: "HarryPotter"},
				"Voldemort":   {Name: "Voldemort"},
			},
			Kills: map[string]int{
				"HarryPotter": 1,
				"Voldemort":   0,
			},
		},
	}

	matchesDBM := []*domain.MatchDeathsByMeans{
		{
			ID: "game_1",
			DeathsByMeans: map[string]int{
				"MOD_TRIGGER_HURT": 1,
				"MOD_RAILGUN":      1,
			},
		},
	}

	matchService := NewMatchService()
	groupedMatches := matchService.GroupMatches(matches)
	groupedMatchesDBM := matchService.GroupMatchesDBM(matchesDBM)

	if len(groupedMatches) != 1 {
		t.Fatalf("Expected 1 grouped match, got %d", len(groupedMatches))
	}

	if len(groupedMatchesDBM) != 1 {
		t.Fatalf("Expected 1 grouped match, got %d", len(groupedMatchesDBM))
	}

	match, exists := groupedMatches["game_1"]
	if !exists {
		t.Fatalf("Expected match with ID game_1")
	}

	if match.TotalKills != 2 {
		t.Fatalf("Expected 2 total kills, got %d", match.TotalKills)
	}

	if len(match.Players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(match.Players))
	}

	if match.Kills["HarryPotter"] != 1 {
		t.Fatalf("Expected HarryPotter to have 1 kill, got %d", match.Kills["HarryPotter"])
	}

	if match.Kills["Voldemort"] != 0 {
		t.Fatalf("Expected Voldemort to have 0 kills, got %d", match.Kills["Voldemort"])
	}

	matchDBM, existsDBM := groupedMatchesDBM["game_1"]

	if !existsDBM {
		t.Fatalf("Expected matchDBM with ID game_1")
	}

	if matchDBM.DeathsByMeans["MOD_TRIGGER_HURT"] != 1 {
		t.Fatalf("Expected 1 death by MOD_TRIGGER_HURT, got %d", matchDBM.DeathsByMeans["MOD_TRIGGER_HURT"])
	}

	if matchDBM.DeathsByMeans["MOD_RAILGUN"] != 1 {
		t.Fatalf("Expected 1 death by MOD_RAILGUN, got %d", matchDBM.DeathsByMeans["MOD_RAILGUN"])
	}
}
