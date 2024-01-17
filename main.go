package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()

	userID := 10
	respCh := make(chan string)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go fetchUserData(userID, respCh, wg)

	wg.Add(1)
	go fetchUserRecommendations(userID, respCh, wg)

	wg.Add(1)
	go fetchUserLikes(userID, respCh, wg)

	wg.Wait()

	close(respCh)

	for resp := range respCh {
		fmt.Println(resp)
	}

	fmt.Println("ELAPSED TIME", time.Since(now))
}

func fetchUserData(userID int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(80 * time.Millisecond)

	respCh <- "user data"

	wg.Done()
}

func fetchUserRecommendations(userID int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(120 * time.Millisecond)

	respCh <- "user recommendations"

	wg.Done()
}

func fetchUserLikes(userID int, respCh chan string, wg *sync.WaitGroup) {
	time.Sleep(50 * time.Millisecond)

	respCh <- "user likes"

	wg.Done()
}
