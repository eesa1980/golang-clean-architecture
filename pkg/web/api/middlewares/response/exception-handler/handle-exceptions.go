package exceptionhandler

import (
	"clean-architecture/pkg/application/common/exception"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type HandleException func(ctx *fiber.Ctx, err error) error

func New() HandleException {
	return func(ctx *fiber.Ctx, err error) error {
		var e exception.Exception

		var badRequestError *exception.BadRequest
		var customError *exception.Custom
		var internalServerError *exception.InternalServer
		var notFoundError *exception.NotFound

		switch {
		case errors.As(err, &badRequestError):
			e = badRequestException(badRequestError)
			break
		case errors.As(err, &customError):
			e = customException(customError)
			break
		case errors.As(err, &internalServerError):
			e = internalServerException(internalServerError)
			break
		case errors.As(err, &notFoundError):
			e = notFoundException(notFoundError)
			break
		default:
			e = internalServerException(&exception.InternalServer{
				Message: err.Error(),
			})
		}

		stringified, err := json.Marshal(e)

		if err != nil {
			return ctx.Status(e.StatusCode).JSON(stringified)
		}

		return ctx.Status(e.StatusCode).JSON(e)
	}

}

func statusCodeHandler(statusCode int) (title string, code int) {
	err := *fiber.NewError(statusCode)
	return err.Message, err.Code
}

func badRequestException(e *exception.BadRequest) exception.Exception {
	title, code := statusCodeHandler(fiber.StatusBadRequest)

	return exception.Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.1",
		Message:     e.Message,
		StatusCode:  code,
		Title:       title,
	}
}

func customException(e *exception.Custom) exception.Exception {
	title, code := statusCodeHandler(e.StatusCode)

	return exception.Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		Message:     e.Message,
		StatusCode:  code,
		Title:       title,
	}
}

func internalServerException(e *exception.InternalServer) exception.Exception {
	title, code := statusCodeHandler(fiber.StatusInternalServerError)

	return exception.Exception{
		Explanation: "https://datatracker.ietf.org/doc/html/rfc7231#section-6.6.1",
		Message:     e.Error(),
		StatusCode:  code,
		Title:       title,
	}
}

func notFoundException(e *exception.NotFound) exception.Exception {
	title, code := statusCodeHandler(fiber.StatusNotFound)

	return exception.Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		Message:     e.Message,
		StatusCode:  code,
		Title:       title,
	}
}
