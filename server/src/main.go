package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matthewyuh246/websocket/src/domain"
	"github.com/matthewyuh246/websocket/src/handlers"
	"github.com/matthewyuh246/websocket/src/services"
)

func main() {
	pubsub := services.NewPubSubService()
	hub := domain.NewHub(pubsub)
	go hub.SubscribeMessages()
	go hub.RunLoop()

	http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

	port := "8080"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}
