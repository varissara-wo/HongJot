package api

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/varissara-wo/hongjot/api/handler"
	"github.com/varissara-wo/hongjot/api/repository"
	"github.com/varissara-wo/hongjot/api/service"
)

type Server struct {
	*echo.Echo
}

func NewServer(db *sql.DB) *Server {
	e := echo.New()

	v1 := e.Group("api/v1")

	{
		r := repository.NewSpenderRepository(db)
		s := service.NewSpenderService(r)
		h := handler.NewSpenderHandler(s)
		v1.POST("/spenders", h.NewSpender)
		v1.GET("/spenders", h.GetAllSpenders)
	}

	return &Server{e}
}
