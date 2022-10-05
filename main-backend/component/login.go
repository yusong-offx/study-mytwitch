package component

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	Id  string `json:"user_id"`
	Pwd string `json:"user_password"`
}

func (u UserInfo) checkValid() error {
	if !(0 < len(u.Id) && len(u.Id) <= 20) {
		return errors.New("invalid id")
	}
	if !(0 < len(u.Pwd) && len(u.Pwd) <= 60) {
		return errors.New("invalid password")
	}
	return nil
}

func (u UserInfo) checkUserExist() error {
	rows, err := DB.Query("SELECT password FROM users WHERE id=$1", u.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var pwd []byte
	if rows.Next() {
		if err = rows.Scan(&pwd); err != nil {
			return err
		}
		if err = bcrypt.CompareHashAndPassword(pwd, []byte(u.Pwd)); err != nil {
			return errors.New("wrong password")
		}
	} else {
		return errors.New("no user")
	}
	return nil
}

func LoginFun(c *fiber.Ctx) error {
	inputinfo := UserInfo{}
	var err error
	if err = c.BodyParser(&inputinfo); err != nil {
		return err
	}
	if err = inputinfo.checkValid(); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if err = inputinfo.checkUserExist(); err != nil {
		return c.Status(403).SendString(err.Error())
	}

	// send jwt token
	jwt_token, err := GernerateJWT(inputinfo.Id)
	if err != nil {
		return err
	}
	newCookie := new(fiber.Cookie)
	newCookie.Name = "jwt"
	newCookie.Value = jwt_token
	newCookie.Expires = time.Now().Add(time.Minute * 15)
	newCookie.HTTPOnly = true
	c.Cookie(newCookie)
	c.Set("loginStatus", "true")

	return c.SendStatus(200)
}
