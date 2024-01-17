package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	userID := 10
	respCh := make(chan string)

	go fetchUserData(userID, respCh)
	go fetchUserRecommendations(userID, respCh)
	go fetchUserLikes(userID, respCh)

	close(respCh)
	
	for resp := range respCh {
		fmt.Println(resp)
	}

	fmt.Println("ELAPSED TIME", time.Since(now))
}

func fetchUserData(userID int, respCh chan string) {
	time.Sleep(80 * time.Millisecond)

	respCh <- "user data"
}

func fetchUserRecommendations(userID int, respCh chan string) {
	time.Sleep(120 * time.Millisecond)

	respCh <- "user recommendations"
}

func fetchUserLikes(userID int, respCh chan string) {
	time.Sleep(50 * time.Millisecond)

	respCh <- "user likes"
}
