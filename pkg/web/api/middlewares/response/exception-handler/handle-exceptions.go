package exceptionhandler

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"wire-demo-2/pkg/application/common/exception"
	"wire-demo-2/pkg/web/crosscutting"
)

type HandleException func(ctx *fiber.Ctx, err error) error

func New(deps *container.Dependencies) HandleException {
	return func(ctx *fiber.Ctx, err error) error {
		var e exception.Exception

		var internalServerError *exception.InternalServer
		var badRequestError *exception.BadRequest
		var httpError *exception.Http
		var notFoundError *exception.NotFound

		switch {
		case errors.As(err, &httpError):
			e = httpException(httpError)
			break
		case errors.As(err, &notFoundError):
			e = notFoundException(notFoundError)
			break
		case errors.As(err, &badRequestError):
			e = badRequestException(badRequestError)
			break
		case errors.As(err, &internalServerError):
			e = internalServerException(internalServerError)
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

func badRequestException(e *exception.BadRequest) exception.Exception {
	return exception.Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.1",
		Message:     e.Message,
		StatusCode:  400,
		Title:       "Bad Request",
	}
}

func httpException(e *exception.Http) exception.Exception {
	return exception.Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		Message:     e.Message,
		StatusCode:  e.StatusCode,
		Title:       "Http Exception",
	}
}

func internalServerException(e *exception.InternalServer) exception.Exception {
	return exception.Exception{
		Explanation: "https://datatracker.ietf.org/doc/html/rfc7231#section-6.6.1",
		Message:     e.Error(),
		StatusCode:  500,
		Title:       "Internal Server Error",
	}
}

func notFoundException(e *exception.NotFound) exception.Exception {
	return exception.Exception{
		Explanation: "https://tools.ietf.org/html/rfc7231#section-6.5.4",
		Message:     e.Message,
		StatusCode:  404,
		Title:       "Not Found",
	}
}
