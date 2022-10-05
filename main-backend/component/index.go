package component

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func IndexFun(c *fiber.Ctx) error {
	jwt_claim, ok := c.Locals("jwt_claim").(jwt.MapClaims)
	j := map[string]interface{}{}
	if !ok {
		j["user_id"] = false
		return c.Status(200).JSON(j)
	}
	j["user_id"] = jwt_claim["user"]
	return c.Status(200).JSON(j)
}
