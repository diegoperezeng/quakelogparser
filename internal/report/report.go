package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"quakelogparser/internal/domain"
	"sort"
)

func GenerateReport(groupedMatches map[string]*domain.Match) {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	keys := make([]string, 0, len(groupedMatches))
	for k := range groupedMatches {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		match := groupedMatches[k]
		reportData := map[string]interface{}{
			"TotalKills": match.TotalKills,
			"Players":    extractPlayerNames(match.Players),
			"Kills":      match.Kills,
		}
		report, err := json.MarshalIndent(reportData, "", "  ")
		if err != nil {
			fmt.Println("Error generating report:", err)
			continue
		}
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s", match.ID, report))
	}
	buffer.WriteString("}")
	fmt.Println(buffer.String())
}

func GenerateReportDeathByMeans(groupedMatchesDBM map[string]*domain.MatchDeathsByMeans) {
	var buffer bytes.Buffer
	buffer.WriteString("{")
	keys := make([]string, 0, len(groupedMatchesDBM))
	for k := range groupedMatchesDBM {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		match := groupedMatchesDBM[k]
		reportData := map[string]interface{}{
			"kills_by_means": match.DeathsByMeans,
		}
		reportDBM, err := json.MarshalIndent(reportData, "", "  ")
		if err != nil {
			fmt.Println("Error generating DeathByMeans report:", err)
			continue
		}
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf("\"%s\":%s", match.ID, reportDBM))
	}
	buffer.WriteString("}")
	fmt.Println(buffer.String())
}

func extractPlayerNames(players map[string]*domain.Player) []string {
	names := make([]string, 0, len(players))
	for name := range players {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
