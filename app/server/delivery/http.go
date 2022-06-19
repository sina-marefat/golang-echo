package delivery

import "github.com/labstack/echo/v4"

type handler struct {
}

func (h *handler) SayHi(ctx echo.Context) error {
	return ctx.JSON(200, "hi there")
}

func RegisterRoutes() {
	h := handler{}
	router := echo.New()
	group := router.Group("/v1")
	{
		group.GET("/test", h.SayHi)
	}

	router.Start(":8080")
}
