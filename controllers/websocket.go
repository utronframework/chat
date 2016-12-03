package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron/controller"
	"github.com/gorilla/websocket"
	"github.com/utronframework/chat/chatroom"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Websocket struct {
	controller.BaseController
	Routes []string
}

func (w *Websocket) Room() {
	r := w.Ctx.Request()
	user := r.URL.Query().Get("user")
	w.Ctx.Data["title"] = "Chat Room"
	w.Ctx.Data["user"] = user
	w.Ctx.Template = "WebSocket/Room"
	w.HTML(http.StatusOK)
}

func (w *Websocket) RoomSocket() {
	fmt.Println("HOME WS")
	r := w.Ctx.Request()
	res := w.Ctx.Response()
	user := r.URL.Query().Get("user")
	conn, err := upgrader.Upgrade(res, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	subscription := chatroom.Subscribe()
	defer subscription.Cancel()

	chatroom.Join(user)
	defer chatroom.Leave(user)

	// Send down the archive.
	for _, event := range subscription.Archive {
		if conn.WriteJSON(event) != nil {
			return
		}
	}

	// In order to select between websocket messages and subscription events, we
	// need to stuff websocket events into a channel.
	newMessages := make(chan string)
	go func() {
		for {
			_, b, err := conn.ReadMessage()
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- string(b)
		}
	}()

	// Now listen for new events from either the websocket or the chatroom.
	for {
		select {
		case event := <-subscription.New:
			if conn.WriteJSON(&event) != nil {
				return
			}
		case msg, ok := <-newMessages:
			// If the channel is closed, they disconnected.
			if !ok {
				return
			}

			// Otherwise, say something.
			chatroom.Say(user, msg)
		}
	}
	return
}

func NewWebsocket() controller.Controller {
	return &Websocket{
		Routes: []string{
			"get;/websocket/room;Room",
			"get;/websocket/room/socket;RoomSocket",
		},
	}
}
