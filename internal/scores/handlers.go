package scores

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ScoreService struct {
	// Redis client aquí
}

func NewScoreService(rdb *redis.Client) *ScoreService {
	return &ScoreService{}
}

func (s *ScoreService) SubmitScoreHandler(c *gin.Context) {
	var req struct {
		UserID string  `json:"user_id"`
		Score  float64 `json:"score"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Lógica de SubmitScore con Redis
	// s.rdb.ZAdd(...)

	c.JSON(http.StatusOK, gin.H{"message": "score submitted"})
}
