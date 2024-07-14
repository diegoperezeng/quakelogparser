package parser

import (
	"fmt"
	"quakelogparser/internal/parser"
	"quakelogparser/internal/service"
	"quakelogparser/internal/utils"
	"quakelogparser/pkg/report"
)

func Run() {
	logFilePath := "files/qgames.log"
	data, err := utils.ReadFile(logFilePath)
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	matches, matchesDBM, err := parser.ParseLogFile(data)
	if err != nil {
		fmt.Println("Error parsing log file:", err)
		return
	}

	matchService := service.NewMatchService()
	groupedMatches := matchService.GroupMatches(matches)
	groupedMatchesDBM := matchService.GroupMatchesDBM(matchesDBM)

	report.GenerateReport(groupedMatches)
	report.GenerateReportDeathByMeans(groupedMatchesDBM)
}
