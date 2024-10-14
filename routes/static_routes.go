package routes

import "github.com/labstack/echo/v4"

func RoutePhotoStatic(apiv1 *echo.Group) {
	apiv1.Static("/storage", "storage")
	apiv1.GET("/storage/:filename", func(c echo.Context) error {
		format := c.Param("filename")[len(c.Param("filename"))-3:]
		if format == "jpg" || format == "png" {
			return c.File("storage/" + c.Param("filename"))
		} else {
			return c.File("storage/default.png")
		}
	})
}
