package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	e "github.com/pkg/errors"
	"github.com/qilin/crm-api/pkg/errors"
	"go.uber.org/zap"
)

//swagger:model
type HTTPError struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

func New(c echo.Context, value interface{}) error {
	return c.JSON(http.StatusOK, value)
}

func Err(err error, c echo.Context) {
	var (
		domainError errors.Error
		echoError   *echo.HTTPError
	)
	switch {
	// domain errors
	case e.As(err, &domainError):
		_ = c.JSON(getStatus(domainError), HTTPError{
			Error: fmt.Sprintf("errors.com.qilin.crm.%s", domainError.Key),
			Msg:   err.Error(), // full message with wrappers
		})

	// catch echo std errors
	case e.Is(err, echo.ErrMethodNotAllowed):
		_ = c.JSON(http.StatusMethodNotAllowed, HTTPError{
			Error: "errors.com.qilin.crm.http.method_not_allowed",
			Msg:   "Method Not Allowed",
		})
	case e.Is(err, echo.ErrUnsupportedMediaType):
		_ = c.JSON(http.StatusMethodNotAllowed, HTTPError{
			Error: "errors.com.qilin.crm.http.unsupported_media_type",
			Msg:   "Unsupported Media Type",
		})
	case e.Is(err, echo.ErrNotFound):
		_ = c.JSON(http.StatusNotFound, HTTPError{
			Error: "errors.com.qilin.crm.http.not_found",
			Msg:   "Not Found",
		})
	case e.Is(err, echo.ErrUnauthorized):
		_ = c.JSON(http.StatusUnauthorized, HTTPError{
			Error: "errors.com.qilin.crm.http.unauthorized",
			Msg:   "Unauthorized",
		})
	case e.As(err, &echoError):
		// bind errors
		if echoError.Code == http.StatusBadRequest {
			_ = c.JSON(http.StatusBadRequest, HTTPError{
				Error: "errors.com.qilin.crm.http.bad_request",
				Msg:   fmt.Sprint(echoError.Message),
			})
		} else {
			zap.L().Error("Unknown error", zap.Error(err))
			_ = c.JSON(http.StatusInternalServerError, HTTPError{
				Error: "errors.com.qilin.crm.internal_error",
				Msg:   "internal server error",
			})
		}

	// unknown errors
	default:
		zap.L().Error("Unknown error", zap.Error(err))
		_ = c.JSON(http.StatusInternalServerError, HTTPError{
			Error: "errors.com.qilin.crm.internal_error",
			Msg:   "internal server error",
		})

	}
}

func getStatus(err errors.Error) int {
	switch err.Type {
	case errors.ErrNotFound:
		return http.StatusNotFound
	case errors.ErrAlreadyExist:
		return http.StatusBadRequest
	case errors.ErrValidation:
		return http.StatusBadRequest
	case errors.ErrInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
