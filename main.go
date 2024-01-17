package main

import (
	"fmt"
	"time"
)

type Message struct {
	From     string
	Payloads string
}

type Server struct {
	msgCh chan Message
}

func (s *Server) StartAndListen() {
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payloads)
		default:

		}
	}
}

func sendMessageToServer(msgCh chan Message, from string, paylaod string) {
	msg := Message{
		From:     from,
		Payloads: paylaod,
	}

	msgCh <- msg
	fmt.Println("sending message")
}

func main() {
	s := &Server{
		msgCh: make(chan Message),
	}
	go s.StartAndListen()

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(2 * time.Second)
			sendMessageToServer(s.msgCh, "John Doe", "Hello gophers")
		}()
	}

	select {}

}
