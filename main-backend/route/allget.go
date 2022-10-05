package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/streaming-server/component"
)

func AllGet(app *fiber.App) {
	app.Get("/", component.IndexFun)
	app.Get("/logout", component.LogoutFun)
}
