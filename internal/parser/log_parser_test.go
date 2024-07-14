package parser

import (
	"testing"
)

func TestParseLogFile(t *testing.T) {

	data := `  19:30 InitGame:
				19:34 ClientConnect: 2
 				19:34 ClientUserinfoChanged: 2 n\HarryPotter\t\0\model\xian/default\hmodel\xian/default\g_redteam\\g_blueteam\\c1\4\c2\5\hc\100\w\0\l\0\tt\0\tl\0
 				19:35 ClientUserinfoChanged: 2 n\HarryPotter\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0
 				19:36 ClientBegin: 2
                19:37 Kill: 1022 2 22: <world> killed Voldemort by MOD_TRIGGER_HURT
                19:37 Kill: 3 2 10: HarryPotter killed Voldemort by MOD_RAILGUN
				19:45 ShutdownGame:
				19:45 ------------------------------------------------------------
				19:45 ------------------------------------------------------------
				20:37 InitGame:
				20:38 ClientConnect: 2
				20:38 ClientUserinfoChanged: 2 n\HarryPotter\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0
				20:38 ClientBegin: 2				
				20:54 Kill: 1022 2 22: <world> killed HarryPotter by MOD_TRIGGER_HURT
				21:07 Kill: 1022 2 22: <world> killed Voldemort by MOD_TRIGGER_HURT
				21:10 ClientDisconnect: 2
				21:37 ShutdownGame:
				21:37 ------------------------------------------------------------
				21:37 ------------------------------------------------------------`

	matches, matchesDBM, err := ParseLogFile(data)
	if err != nil {
		t.Fatalf("Error parsing log file: %v", err)
	}

	if len(matches) != 2 {
		t.Fatalf("Expected 2 match, got %d", len(matches))
	}

	if len(matchesDBM) != 2 {
		t.Fatalf("Expected 2 match in DeathByMeans, got %d", len(matchesDBM))
	}

	match := matches[0]
	if match.TotalKills != 2 {
		t.Fatalf("Expected 2 total kills, got %d", match.TotalKills)
	}

	if len(match.Players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(match.Players))
	}

	if match.Kills["HarryPotter"] != 1 {
		t.Fatalf("Expected HarryPotter to have 1 kill, got %d", match.Kills["Player_3"])
	}

	if match.Kills["Voldemort"] != -1 {
		t.Fatalf("Expected Voldemort to have -1 kill, got %d", match.Kills["Player_2"])
	}

	match2 := matches[1]
	if match2.TotalKills != 2 {
		t.Fatalf("Expected 2 total kills, got %d", match2.TotalKills)
	}

	if len(match2.Players) != 2 {
		t.Fatalf("Expected 2 players, got %d", len(match2.Players))
	}

	if match2.Kills["HarryPotter"] != -1 {
		t.Fatalf("Expected HarryPotter to have 1 kill, got %d", match2.Kills["Player_3"])
	}

	if match2.Kills["Voldemort"] != -1 {
		t.Fatalf("Expected Voldemort to have -1 kill, got %d", match2.Kills["Player_2"])
	}

	matchDBM := matchesDBM[0]
	if matchDBM.DeathsByMeans["MOD_TRIGGER_HURT"] != 1 {
		t.Fatalf("Expected 1 death by MOD_TRIGGER_HURT, got %d", matchDBM.DeathsByMeans["MOD_TRIGGER_HURT"])
	}

	if matchDBM.DeathsByMeans["MOD_RAILGUN"] != 1 {
		t.Fatalf("Expected 1 death by MOD_RAILGUN, got %d", matchDBM.DeathsByMeans["MOD_RAILGUN"])
	}

	matchDBM2 := matchesDBM[1]
	if matchDBM2.DeathsByMeans["MOD_TRIGGER_HURT"] != 2 {
		t.Fatalf("Expected 1 death by MOD_TRIGGER_HURT, got %d", matchDBM2.DeathsByMeans["MOD_TRIGGER_HURT"])
	}
}
