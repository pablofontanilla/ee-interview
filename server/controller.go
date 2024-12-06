package server

import (
	"fmt"
	"log"
	"mywsapp/fibonacci"
	"os"
	"os/signal"
	"syscall"
)

var prompt = "<b>What fibbonaci index do you want to know?</b>"

type IncommingMsg struct {
	Source  string
	Payload []byte
}

type WebsocketServerController struct {
	// Registered clients
	clients map[*ServerClient]bool

	// Inbound messages from the client
	broadcast chan IncommingMsg

	// register requests from clients
	register chan *ServerClient

	// Unregister request from clients
	unregister chan *ServerClient

	// Shutdown the controller
	shutdown chan os.Signal
}

func newWebsocketHub() *WebsocketServerController {
	return &WebsocketServerController{
		broadcast:  make(chan IncommingMsg),
		register:   make(chan *ServerClient),
		unregister: make(chan *ServerClient),
		clients:    make(map[*ServerClient]bool),
		shutdown:   nil,
	}
}

func (h *WebsocketServerController) run() {
	h.shutdown = make(chan os.Signal)

	signal.Notify(h.shutdown, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.sendClientMsg(client, fmt.Sprintf("<b>Registered as %s</b>", client.id))
			h.broadcastMsg(fmt.Sprintf("<i>User %s has joined</i>", client.id))
			h.broadcastMsg(prompt)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			// Echo the input to other clients
			h.echoMsgToOtherClients(message)

			// Parse the user input
			resp, err := fibonacci.ParseIndex(string(message.Payload))

			// Send the response
			if err != nil {
				h.broadcastMsg(fmt.Sprintf("server error: unable to parse the given index: %s", err))
			} else {
				h.broadcastMsg(fmt.Sprintf("<b>Answer:</b> %s", resp))
			}

			// resend the prompt
			h.broadcastMsg(prompt)

		case <-h.shutdown:
			log.Println("Killing the controller")
			return
		}
	}
}

func (h *WebsocketServerController) echoMsgToOtherClients(msg IncommingMsg) {
	for client := range h.clients {
		if client.id != msg.Source {
			h.sendClientMsg(client, fmt.Sprintf("%s: %s", msg.Source, string(msg.Payload)))
		}
	}
}

func (h *WebsocketServerController) broadcastMsg(msg string) {
	for client := range h.clients {
		h.sendClientMsg(client, msg)
	}
}

func (h *WebsocketServerController) sendClientMsg(c *ServerClient, msg string) {
	c.send <- []byte(msg)
}
