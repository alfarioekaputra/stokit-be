package middleware

import (
	"fmt"
	"stokit/internal/model"
	"stokit/internal/usecase"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func NewAuth(userUserCase *usecase.UserUsecase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &model.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		userUserCase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUserCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userUserCase.Log.Debugf("User : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("user").(*model.Auth)
}

func VerifyToken(ctx *fiber.Ctx) error {
	var token string

	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		token = strings.TrimPrefix(authorization, "Bearer ")
	} else if ctx.Cookies("token") != "" {
		token = ctx.Cookies("token")
	}

	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenByte, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {

		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}
		return []byte(viper.GetString("auth.jwtKey")), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	userId, ok := claims["sub"].(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token subject"})
	}

	ctx.Locals("user_id", claims["sub"])
	ctx.Locals("name", claims["name"])

	ctx.Locals("user", &model.Auth{
		ID: userId,
	})

	return ctx.Next()

}
