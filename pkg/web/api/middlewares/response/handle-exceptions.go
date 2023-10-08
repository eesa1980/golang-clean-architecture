package web

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	. "wire-demo-2/pkg/application/common/exceptions"
	"wire-demo-2/pkg/web/crosscutting"
)

type HandleException func(ctx *fiber.Ctx, err error) error

func MakeHandleException(deps *crosscutting.Dependencies) HandleException {
	return func(ctx *fiber.Ctx, err error) error {
		var exception Exception

		var internalServerError *InternalServerException
		var badRequestError *BadRequestException
		var httpError *HttpException
		var notFoundError *NotFoundException

		switch {
		case errors.As(err, &httpError):
			exception = httpException(httpError)
			break
		case errors.As(err, &notFoundError):
			exception = notFoundException(notFoundError)
			break
		case errors.As(err, &badRequestError):
			exception = badRequestException(badRequestError)
			break
		case errors.As(err, &internalServerError):
			exception = internalServerException(internalServerError)
			break
		default:
			exception = internalServerException(&InternalServerException{
				Message: err.Error(),
			})
		}

		stringified, err := json.Marshal(exception)

		if err != nil {
			return ctx.Status(exception.StatusCode).JSON(stringified)
		}

		return ctx.Status(exception.StatusCode).JSON(exception)
	}

}

func badRequestException(e *BadRequestException) Exception {
	return Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.1",
		Message:     e.Message,
		StatusCode:  400,
		Title:       "Bad Request",
	}
}

func httpException(e *HttpException) Exception {
	return Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		Message:     e.Message,
		StatusCode:  e.StatusCode,
		Title:       "Http Exception",
	}
}

func internalServerException(e *InternalServerException) Exception {
	return Exception{
		Explanation: "https://datatracker.ietf.org/doc/html/rfc7231#section-6.6.1",
		Message:     e.Error(),
		StatusCode:  500,
		Title:       "Internal Server Error",
	}
}

func notFoundException(e *NotFoundException) Exception {
	return Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		Message:     e.Message,
		StatusCode:  404,
		Title:       "Not Found",
	}
}
