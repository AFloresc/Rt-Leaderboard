package reports

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type ReportService struct {
	rdb *redis.Client
}

func NewReportService(rdb *redis.Client) *ReportService {
	return &ReportService{rdb: rdb}
}

// Handler: Reporte de top players por periodo
func (s *ReportService) GetReportHandler(c *gin.Context) {
	period := c.Param("period") // ej: "2025-11"
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		limit = 10
	}

	key := "leaderboard:" + period

	results, err := s.rdb.ZRevRangeWithScores(c, key, 0, limit-1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch report"})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no data for this period"})
		return
	}

	players := []gin.H{}
	for i, r := range results {
		players = append(players, gin.H{
			"rank":   i + 1,
			"userID": r.Member,
			"score":  r.Score,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"period": period,
		"report": players,
	})
}
