package component

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients = map[*websocket.Conn]struct{}{}

func RequestWebsocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func ChatApp(c *websocket.Conn) {
	defer c.Close()

	clients[c] = struct{}{}
	log.Printf("'%s' connect", c.Params("user_id", "unknown"))
	log.Println("user count :", len(clients))
	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			delete(clients, c)
			log.Println("user count :", len(clients), "/", err)
			break
		}
		msg = append([]byte(c.Params("user_id")+" : "), msg...)
		for k := range clients {
			if err = k.WriteMessage(mt, msg); err != nil {
				log.Println(err)
			}
		}
	}
}
