package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"quakelogparser/internal/domain"
	"sort"
	"strconv"
)

func GenerateReport(groupedMatches map[string]*domain.Match) {
	ensureOutputDirectory()

	allMatches := make(map[string]interface{})
	for _, match := range groupedMatches {
		reportData := map[string]interface{}{
			"TotalKills": match.TotalKills,
			"Players":    extractPlayerNames(match.Players),
			"Kills":      match.Kills,
		}
		allMatches[match.ID] = reportData
	}

	filePath := filepath.Join("output", "matches.json")
	saveToFile(filePath, allMatches)
}

func GenerateReportDeathByMeans(groupedMatchesDBM map[string]*domain.MatchDeathsByMeans) {
	ensureOutputDirectory()

	allMatchesDBM := make(map[string]interface{})
	for _, match := range groupedMatchesDBM {
		reportData := map[string]interface{}{
			"kills_by_means": match.DeathsByMeans,
		}
		allMatchesDBM[match.ID] = reportData
	}

	filePath := filepath.Join("output", "matches_dbm.json")
	saveToFile(filePath, allMatchesDBM)
}

func extractPlayerNames(players map[string]*domain.Player) []string {
	names := make([]string, 0, len(players))
	for name := range players {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func ensureOutputDirectory() {
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		err = os.Mkdir("output", 0755)
		if err != nil {
			fmt.Println("Error creating output directory:", err)
		}
	}
}

func saveToFile(filePath string, data map[string]interface{}) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	orderedData := make(map[string]interface{})
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		numI, _ := strconv.Atoi(keys[i][5:])
		numJ, _ := strconv.Atoi(keys[j][5:])
		return numI < numJ
	})

	for _, k := range keys {
		orderedData[k] = data[k]
	}

	report, err := json.MarshalIndent(orderedData, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}
	file.Write(report)
}
