# Rt-Leaderboard
A **real-time leaderboard system** built with **Go** and **Redis**.  
It allows users to compete in games or activities, submit scores, and view rankings globally or by specific periods.

---

## Description
This project involves creating a backend system for a real-time leaderboard service. The service will allow users to compete in various games or activities, track their scores, and view their rankings on a leaderboard. The system will feature user authentication, score submission, real-time leaderboard updates, and score history tracking. Redis sorted sets will be used to manage and query the leaderboards efficiently.
---

## Project Requirements
It is a build of a real-time leaderboard system that ranks users based on their scores in various games or activities. The system should meet the following 
---

## Requirements:
-User Authentication: Users should be able to register and log in to the system.

-Score Submission: Users should be able to submit their scores for different games or activities.

-Leaderboard Updates: Display a global leaderboard showing the top users across all games.

-User Rankings: Users should be able to view their rankings on the leaderboard.

-Top Players Report: Generate reports on the top players for a specific period.
---

## 🚀 Features

- **User Authentication**: Register and log in with JWT.
- **Score Submission**: Submit scores in real time.
- **Global Leaderboard**: View the top players across all games.
- **User Rankings**: Check individual user positions.
- **Reports**: Generate top player reports for specific periods (`YYYY-MM`).
- **Redis Sorted Sets**: Efficient ranking management and queries.

---

## 📂 Project Structure
rt-leaderboard/ 
│── cmd/ 
│ └── main.go # API entry point 
│── internal/ 
│ ├── auth/ # Registration, login, JWT 
│ ├── leaderboard/ # Ranking logic with Redis 
│ ├── scores/ # Score submission 
│ ├── reports/ # Period-based reports 
│── pkg/ 
│ └── redis/ # Centralized Redis client 
│── Dockerfile 
│── docker-compose.yml 
│── Makefile 
│── README.md
---

## ⚙️ Installation

### Requirements
- Go 1.22+
- Docker & Docker Compose

### Steps
```bash
# Clone repository
git clone https://github.com/youruser/rt-leaderboard.git
cd rt-leaderboard

# Start services with Docker
make up

# Stop services
make down