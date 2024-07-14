package service

import "quakelogparser/internal/domain"

type MatchService struct{}

func NewMatchService() *MatchService {
	return &MatchService{}
}

func (ms *MatchService) GroupMatches(matches []*domain.Match) map[string]*domain.Match {
	groupedMatches := make(map[string]*domain.Match)
	for _, match := range matches {
		groupedMatches[match.ID] = match
	}
	return groupedMatches
}

func (ms *MatchService) GroupMatchesDBM(matches []*domain.MatchDeathsByMeans) map[string]*domain.MatchDeathsByMeans {
	groupedMatchesDBM := make(map[string]*domain.MatchDeathsByMeans)
	for _, match := range matches {
		groupedMatchesDBM[match.ID] = match
	}
	return groupedMatchesDBM
}
