package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/varissara-wo/hongjot/api/errs"
	"github.com/varissara-wo/hongjot/api/service"
)

type spenderHandler struct {
	spenderService service.SpenderService
}

func NewSpenderHandler(ss service.SpenderService) spenderHandler {
	return spenderHandler{spenderService: ss}
}

func (h spenderHandler) NewSpender(c echo.Context) error {
	// validate
	req := service.NewSpenderRequest{}
	err := c.Bind(&req)
	if err != nil {
		handleError(c, errs.NewBadRequestError())
	}
	res, err := h.spenderService.NewSpender(req)
	if err != nil {
		handleError(c, err)
	}
	return c.JSON(http.StatusCreated, res)
}

func (h spenderHandler) GetAllSpenders(c echo.Context) error {
	res, err := h.spenderService.GetAllSpenders()
	if err != nil {
		handleError(c, err)
		// return
	}
	return c.JSON(http.StatusOK, res)
}
