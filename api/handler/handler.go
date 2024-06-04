package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/varissara-wo/hongjot/api/errs"
)

func handleError(c echo.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		c.JSON(e.Code, e.Message)
	case error:
		c.JSON(http.StatusInternalServerError, e.Error())
	}
}
