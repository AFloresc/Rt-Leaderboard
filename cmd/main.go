package main

import (
	"real-time-leaderboard/internal/auth"
	"real-time-leaderboard/internal/leaderboard"
	"real-time-leaderboard/internal/reports"
	"real-time-leaderboard/internal/scores"
	"real-time-leaderboard/pkg/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Inicializar Redis
	rdb := redis.NewClient()

	// Servicios
	authService := auth.NewAuthService()
	leaderboardService := leaderboard.NewLeaderboardService(rdb)
	scoreService := scores.NewScoreService(rdb)
	reportService := reports.NewReportService(rdb)

	// Grupo público (sin JWT)
	r.POST("/register", authService.RegisterHandler)
	r.POST("/login", authService.LoginHandler)

	// Grupo protegido con JWT
	api := r.Group("/")
	api.Use(auth.JWTMiddleware())

	{
		api.POST("/scores", scoreService.SubmitScoreHandler)
		api.GET("/leaderboard", leaderboardService.GetTopPlayersHandler)
		api.GET("/leaderboard/:userID", leaderboardService.GetUserRankHandler)
		api.GET("/reports/:period", reportService.GetReportHandler)
	}

	r.Run(":8080")
}
