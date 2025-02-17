package routes

import (
	"Bookify/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterCategoryRoutes(e *echo.Echo, h *handlers.CategoryHandler)  {
	e.GET("/categories",h.GetAllCategory)
}


func RegisterCategoryAllRoutes(e *echo.Echo, categoryHandler *handlers.CategoryHandler) {
	RegisterCategoryRoutes(e,categoryHandler)
}