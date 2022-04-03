package status

import (
	"github.com/labstack/echo"
	"net/http"
)

type responseError struct {
	Message string
	Code    string
}
type response struct {
	Data interface{}    `json:"data,omitempty"`
	err  *responseError `json:"err,omitempty"`
}

func Success(c echo.Context, data interface{}) error {
	if data == nil {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, response{
		Data: data,
	})
}
