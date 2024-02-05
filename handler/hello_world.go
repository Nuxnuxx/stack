package handler

import (
	"stack/view/hello_world"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct {
}

func (h *HelloHandler) Show(c echo.Context) error {
	return render(c, helloworld.Hello())
}
