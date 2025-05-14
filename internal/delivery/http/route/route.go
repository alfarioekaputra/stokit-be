package route

import (
	"stokit/internal/delivery/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	UserController     *http.UserController
	ProductController  *http.ProductController
	CategoryController *http.CategoryController
	AuthMiddleware     fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupStaticRoute()
	SetupCategoryRoutes(c)
	SetupUserRoutes(c)
}

func (c *RouteConfig) SetupStaticRoute() {
	c.App.Static("/", "./views/dist", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		Index:     "index.html",
	})

	c.App.Use(func(ctx *fiber.Ctx) error {
		path := ctx.Path()
		method := ctx.Method()

		// Hanya intercept GET request, bukan ke /api atau file statis (.js, .css, .png, dll)
		if method == fiber.MethodGet &&
			!strings.HasPrefix(path, "/api") &&
			!strings.Contains(path, ".") {
			return ctx.SendFile("./views/dist/index.html")
		}

		return ctx.Next()
	})

}
