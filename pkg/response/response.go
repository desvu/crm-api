package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	e "github.com/pkg/errors"
	"github.com/qilin/crm-api/pkg/errors"
)

type HTTPError struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
}

func New(c echo.Context, value interface{}) error {
	return c.JSON(http.StatusOK, value)
}

func Err(err error, c echo.Context) {
	switch v := e.Cause(err).(type) {
	case errors.Error:
		_ = c.JSON(getStatus(v), HTTPError{
			Error: fmt.Sprintf("errors.com.qilin.crm.%s", v.Key),
			Msg:   getMsg(v),
		})
	default:
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

func getMsg(err errors.Error) string {
	switch err.Type {
	case errors.ErrInternal:
		return "internal server error"
	default:
		return err.Err.Error()
	}
}
