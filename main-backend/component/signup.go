package component

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserSignUpInfo struct {
	Id  string `json:"user_id"`
	Pwd string `json:"user_password"`
}

func (u UserSignUpInfo) checkValid() error {
	if !(0 < len(u.Id) && len(u.Id) <= 20) {
		return errors.New("invalid id")
	}
	if !(0 < len(u.Pwd) && len(u.Pwd) <= 60) {
		return errors.New("invalid password")
	}
	return nil
}

func (u UserSignUpInfo) checkUserExist() error {
	rows, err := DB.Query("SELECT id FROM users WHERE id=$1", u.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return errors.New("user id exist. please change your id")
	}
	return nil
}

func (u UserSignUpInfo) postToDB() error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(u.Pwd), 10)
	if err != nil {
		return err
	}
	rows, err := DB.Query("INSERT INTO users (id, password) VALUES ($1, $2)", u.Id, hashPwd)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func SignUpFun(c *fiber.Ctx) error {
	usersignupinfo := UserSignUpInfo{}
	var err error

	if err = c.BodyParser(&usersignupinfo); err != nil {
		return err
	}
	if err = usersignupinfo.checkValid(); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if err = usersignupinfo.checkUserExist(); err != nil {
		return c.Status(409).SendString(err.Error())
	}
	if err = usersignupinfo.postToDB(); err != nil {
		return err
	}
	return c.SendStatus(201)
}
