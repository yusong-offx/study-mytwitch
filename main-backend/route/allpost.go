package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/streaming-server/component"
)

func AllPost(app *fiber.App) {

	app.Post("/login", component.LoginFun)
	app.Post("/signup", component.SignUpFun)
}
