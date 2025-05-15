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
	c.SetupGuestRoute()
	c.SetupStaticRoute()
	SetupCategoryRoutes(c)
	SetupUserRoutes(c)
}

func (c *RouteConfig) SetupGuestRoute() {
	//Guest
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)

	//GetCagtegoryTree
	c.App.Get("/api/category/tree", c.CategoryController.GetTree)
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
