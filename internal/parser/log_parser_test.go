package parser

import (
	"testing"
)

func TestParseLogFile(t *testing.T) {
	data := `  0:00 InitGame:
                20:37 Kill: 1022 2 22: <world> killed HarryPotter by MOD_TRIGGER_HURT
                2:22 Kill: 3 2 10: HarryPotter killed Voldemort by MOD_RAILGUN`
	matches, err := ParseLogFile(data)
	if err != nil {
		t.Fatalf("Error parsing log file: %v", err)
	}

	if len(matches) != 1 {
		t.Fatalf("Expected 1 match, got %d", len(matches))
	}

	match := matches[0]
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

	if match.DeathsByMeans["MOD_TRIGGER_HURT"] != 1 {
		t.Fatalf("Expected 1 death by MOD_TRIGGER_HURT, got %d", match.DeathsByMeans["MOD_TRIGGER_HURT"])
	}

	if match.DeathsByMeans["MOD_RAILGUN"] != 1 {
		t.Fatalf("Expected 1 death by MOD_RAILGUN, got %d", match.DeathsByMeans["MOD_RAILGUN"])
	}
}
