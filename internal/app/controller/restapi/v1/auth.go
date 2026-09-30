package v1

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// TODO
// The REAL authorization process should be here
// This is a fake, for demonstration only

func (r *V1) getToken(ctx *fiber.Ctx) error {
	claims := jwt.MapClaims{
		"sub": "test-user-id",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "Something went wrong")
	}

	return ctx.Status(http.StatusCreated).JSON(tokenString)
}
