package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/yusong-offx/streaming-server/component"
	"github.com/yusong-offx/streaming-server/route"
)

func main() {
	// DB Connect
	if err := component.DBConnect(); err != nil {
		log.Fatal("Database disconnect", err)
	}
	defer component.DB.Close()

	// Fiber Start
	app := fiber.New()

	route.AllMiddleware(app)
	route.AllGet(app)
	route.AllPost(app)

	log.Fatal(app.Listen(":8000"))
}
