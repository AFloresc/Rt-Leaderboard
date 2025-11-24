package leaderboard

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type LeaderboardService struct {
	rdb *redis.Client
}

func NewLeaderboardService(rdb *redis.Client) *LeaderboardService {
	return &LeaderboardService{rdb: rdb}
}

// Añadir puntuación
func (s *LeaderboardService) SubmitScore(userID string, score float64) error {
	return s.rdb.ZAdd(context.Background(), "global:leaderboard", redis.Z{
		Score:  score,
		Member: userID,
	}).Err()
}

// Obtener top N jugadores
func (s *LeaderboardService) GetTopPlayers(limit int64) ([]redis.Z, error) {
	return s.rdb.ZRevRangeWithScores(context.Background(), "global:leaderboard", 0, limit-1).Result()
}

// Obtener ranking de un usuario
func (s *LeaderboardService) GetUserRank(userID string) (int64, error) {
	return s.rdb.ZRevRank(context.Background(), "global:leaderboard", userID).Result()
}
