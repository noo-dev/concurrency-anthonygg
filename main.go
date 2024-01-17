package main

import "fmt"

type Message struct {
	From     string
	Payloads string
}

type Server struct {
	msgCh chan Message
}

func (s *Server) StartAndListen() {
	for {
		msg := <-s.msgCh
		fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payloads)
	}
}

func sendMessageToServer(msgCh chan Message, paylaod string) {
	msg := Message{
		From:     "john doe",
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

	sendMessageToServer(s.msgCh, "Hello gophers")
}
