package component

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutFun(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: "randomvalue",
		// Set expiry date to the past
		Expires:  time.Now().Add(-(time.Hour * 2)),
		HTTPOnly: true,
	})
	return c.SendStatus(200)
}
