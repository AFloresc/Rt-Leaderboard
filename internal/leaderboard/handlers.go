package leaderboard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LeaderboardService struct {
	rdb *redis.Client
}

func NewLeaderboardService(rdb *redis.Client) *LeaderboardService {
	return &LeaderboardService{rdb: rdb}
}

// Handler: Top N jugadores
func (s *LeaderboardService) GetTopPlayersHandler(c *gin.Context) {
	// Query param: ?limit=10
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		limit = 10
	}

	results, err := s.rdb.ZRevRangeWithScores(c, "global:leaderboard", 0, limit-1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch leaderboard"})
		return
	}

	// Formatear respuesta
	players := []gin.H{}
	for i, r := range results {
		players = append(players, gin.H{
			"rank":   i + 1,
			"userID": r.Member,
			"score":  r.Score,
		})
	}

	c.JSON(http.StatusOK, gin.H{"leaderboard": players})
}

// Handler: Ranking de un usuario
func (s *LeaderboardService) GetUserRankHandler(c *gin.Context) {
	userID := c.Param("userID")

	rank, err := s.rdb.ZRevRank(c, "global:leaderboard", userID).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found in leaderboard"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch rank"})
		return
	}

	score, _ := s.rdb.ZScore(c, "global:leaderboard", userID).Result()

	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
		"rank":   rank + 1, // Redis rank es 0-based
		"score":  score,
	})
}
