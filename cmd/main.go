package main

import (
	"rt-leaderboard/internal/auth"
	"rt-leaderboard/internal/leaderboard"
	"rt-leaderboard/internal/reports"
	"rt-leaderboard/internal/scores"
	"rt-leaderboard/pkg/redis"

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
