package scores

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ScoreService struct {
	// Redis client aquí
	rdb *redis.Client
}

func NewScoreService(rdb *redis.Client) *ScoreService {
	return &ScoreService{rdb: rdb}
}

// Handler: Enviar puntuación
func (s *ScoreService) SubmitScoreHandler(c *gin.Context) {
	var req struct {
		UserID string  `json:"user_id"`
		Score  float64 `json:"score"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Validación básica
	if req.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id required"})
		return
	}

	// Actualizar leaderboard global
	err := s.rdb.ZAdd(c, "global:leaderboard", redis.Z{
		Score:  req.Score,
		Member: req.UserID,
	}).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update global leaderboard"})
		return
	}

	// Actualizar leaderboard del periodo actual (YYYY-MM)
	period := time.Now().Format("2006-01") // ej: "2025-11"
	periodKey := "leaderboard:" + period

	err = s.rdb.ZAdd(c, periodKey, redis.Z{
		Score:  req.Score,
		Member: req.UserID,
	}).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update period leaderboard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "score submitted",
		"userID":  req.UserID,
		"score":   req.Score,
		"period":  period,
	})
}
