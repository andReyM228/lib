package responder

import (
	"github.com/andReyM228/lib/errs"
	"github.com/gofiber/fiber/v2"
)

const (
	tokenName = "Authorization"
)

// Respond - helper for create response
func Respond(ctx *fiber.Ctx, statusCode int, payload interface{}) error {
	ctx.Response().SetStatusCode(statusCode)

	if err := ctx.JSON(payload); err != nil {
		return err
	}

	return nil
}

func HandleError(ctx *fiber.Ctx, err error) error {
	switch err.(type) {
	case errs.NotFoundError:
		if err := Respond(ctx, fiber.StatusNotFound, err); err != nil {
			return err
		}
	case errs.BadRequestError:
		if err := Respond(ctx, fiber.StatusBadRequest, err); err != nil {
			return err
		}
	case errs.ForbiddenError:
		if err := Respond(ctx, fiber.StatusForbidden, err); err != nil {
			return err
		}
	case errs.UnauthorizedError:
		if err := Respond(ctx, fiber.StatusUnauthorized, err); err != nil {
			return err
		}
	default:
		if err := Respond(ctx, fiber.StatusInternalServerError, err); err != nil {
			return err
		}

	}

	return nil
}

func GetToken(ctx *fiber.Ctx) (string, error) {
	headers := ctx.GetReqHeaders()

	value, ok := headers[tokenName]
	if ok {
		return value, nil
	}

	return "", errs.UnauthorizedError{Cause: "invalid authorization header"}
}
