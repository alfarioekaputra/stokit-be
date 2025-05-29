package helper

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractUnverifiedClaims(tokenString string, c *fiber.Ctx) (err error) {

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err.Error())})
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (_ interface{}, err error) {
		return nil, err
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Fail to parse data token"})
	}

	defer func() {
		exp := claims["exp"].(float64)
		expInt := int64(exp)
		expTime := time.Unix(expInt, 0)
		currentTime := time.Now()
		if currentTime.After(expTime) {
			err = errors.New("token expired")
			return
		}
		username := claims["username"].(string)

		c.Locals("username", username)

		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				// Fallback err (per specs, error strings should be lowercase w/o punctuation
				err = errors.New("parse fail please asccess with correct token")
			}
		}
	}()
	return err
}

func GenerateTokenPair(user_id string, username string) (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user_id
	claims["name"] = username
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Add(time.Minute * -5).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("rahasia"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
