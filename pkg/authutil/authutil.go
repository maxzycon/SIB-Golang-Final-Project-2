package authutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-2/pkg/errors"
)

type UserClaims struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
	Role     uint   `json:"role"`
}

func GetCredential(f *fiber.Ctx) (resp *UserClaims, err error) {
	resp, ok := f.Locals("user_profile").(*UserClaims)
	if !ok {
		log.Errorf("err parse data profile to userClaims")
		err = errors.ErrUnauthorized
		return
	}
	return
}
