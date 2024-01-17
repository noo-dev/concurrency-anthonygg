package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	userID := 10
	userData := go fetchUserData(userID)
	userRecs := go fetchUserRecommendations(userID)
	userLikes := go fetchUserLikes(userID)

	fmt.Println(userData)
	fmt.Println(userRecs)
	fmt.Println(userLikes)

	fmt.Println("ELAPSED TIME", time.Since(now))
}

func fetchUserData(userID int) string {
	time.Sleep(80 * time.Millisecond)

	return "user data"
}

func fetchUserRecommendations(userID int) string {
	time.Sleep(120 * time.Millisecond)

	return "user recommendations"
}

func fetchUserLikes(userID int) string {
	time.Sleep(50 * time.Millisecond)

	return "user likes"
}
