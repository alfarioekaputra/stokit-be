package route

import (
	"stokit/internal/delivery/http"

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
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	// c.App.Static("/", "./views/dist", fiber.Static{
	// 	Compress:  true,
	// 	ByteRange: true,
	// 	Browse:    true,
	// 	Index:     "index.html",
	// })
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)

	// c.App.Use(func(ctx *fiber.Ctx) error {
	// 	path := ctx.Path()
	// 	method := ctx.Method()

	// 	// Hanya intercept GET request, bukan ke /api atau file statis (.js, .css, .png, dll)
	// 	if method == fiber.MethodGet &&
	// 		!strings.HasPrefix(path, "/api") &&
	// 		!strings.Contains(path, ".") {
	// 		return ctx.SendFile("./views/dist/index.html")
	// 	}

	// 	return ctx.Next()
	// })

}

func (c *RouteConfig) SetupAuthRoute() {
	api := c.App.Group("/api", c.AuthMiddleware)
	api.Delete("/users", c.UserController.Logout)
	api.Patch("/users/_current", c.UserController.Update)
	api.Get("/users/_current", c.UserController.Current)
	api.Get("/users", c.UserController.FetchAll)

	//category
	api.Get("/category", c.CategoryController.List)
	api.Post("/category", c.CategoryController.Create)

}
