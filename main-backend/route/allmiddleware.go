package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
	"github.com/yusong-offx/streaming-server/component"
)

func AllMiddleware(app *fiber.App) {
	app.Use(recover.New())

	app.Use(csrf.New(csrf.Config{
		KeyLookup: "cookie:csrf_",
	}))

	// file server
	app.Static("/streamer", "/filestorage")

	// login checker
	app.Use(func(c *fiber.Ctx) error {
		jwt_token := c.Cookies("jwt", "")
		if jwt_token != "" {
			claim, valid := component.ValidateJWT(jwt_token)
			if !valid {
				return c.Redirect("/logout", fiber.StatusMovedPermanently)
			}
			c.Locals("jwt_claim", claim)
		}
		return c.Next()
	})

	// chat websocket
	app.Use("/ws", component.RequestWebsocket)
	app.Get("/ws/:user_id", websocket.New(component.ChatApp))

}
