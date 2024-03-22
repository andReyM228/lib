package auth

import (
	"errors"
	"fmt"
	"github.com/andReyM228/lib/errs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

const (
	secret        = "nwciudamsdeqdincacm"
	Authorization = "Authorization"
	ChatIdHeader  = "h_chat_id"
)

func CreateToken(chatID, userID int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем поля токена
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = chatID
	claims["user_id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	// Подписываем токен
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error while signing token:", err)
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string, chatID ...int64) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что используется метод подписи HS256 и возвращаем секретный ключ
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("Error while verifying token:", err)
		return err
	}

	// Получаем поля токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"].(float64)
		userID := claims["user_id"].(float64)
		iat := claims["iat"].(float64)
		exp := claims["exp"].(float64)

		if len(chatID) != 0 {
			if int64(sub) != chatID[0] {
				return errors.New("wrong chatID")
			}
		}

		fmt.Printf("Token verified. sub=%f, user_id=%f, iat=%s, exp=%s\n", sub, userID, time.Unix(int64(iat), 0), time.Unix(int64(exp), 0))

	} else {
		return errors.New("invalid token")
	}

	return nil
}

func GetChatID(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("Error while verifying token:", err)
		return 0, err
	}

	// Получаем поля токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"].(float64)

		return int64(sub), nil
	} else {
		return 0, errors.New("invalid token")
	}
}

func AuthMiddleware(enabled bool) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if !enabled {
			return ctx.Next()
		}

		token := ctx.Get(Authorization)
		if token == "" {
			return errs.UnauthorizedError{Cause: "token"}
		}

		if err := VerifyToken(token); err != nil {
			return errs.UnauthorizedError{Cause: err.Error()}
		}

		chatID, err := GetChatID(token)
		if err != nil {
			return errs.UnauthorizedError{Cause: err.Error()}
		}

		ctx.Request().Header.Add(ChatIdHeader, strconv.FormatInt(chatID, 10))

		return ctx.Next()
	}
}

func GetChatIDFromHeader(ctx *fiber.Ctx) (int64, error) {
	chatID := ctx.Get(ChatIdHeader)
	if chatID == "" {
		return 0, errs.UnauthorizedError{Cause: "chatID"}
	}

	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return 0, errs.UnauthorizedError{Cause: err.Error()}
	}

	return id, nil
}
