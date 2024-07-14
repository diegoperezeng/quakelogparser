package report

import (
	"encoding/json"
	"fmt"
	"quakelogparser/internal/domain"
)

func GenerateReport(groupedMatches map[string]*domain.Match) {
	for _, match := range groupedMatches {
		report, err := json.MarshalIndent(match, "", "  ")
		if err != nil {
			fmt.Println("Error generating report:", err)
			continue
		}
		fmt.Println(string(report))
	}
}

func GenerateReportDeathByMeans(groupedMatchesDBM map[string]*domain.MatchDeathsByMeans) {
	for _, match := range groupedMatchesDBM {
		reportDBM, err := json.MarshalIndent(match, "", "  ")
		if err != nil {
			fmt.Println("Error generating DeathByMeans report:", err)
			continue
		}
		fmt.Println(string(reportDBM))
	}
}
