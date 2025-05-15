package route

func SetupCategoryRoutes(c *RouteConfig) {
	category := c.App.Group("/api", c.AuthMiddleware)
	category.Get("/category", c.CategoryController.List)
	category.Get("/category/:categoryId/view", c.CategoryController.View)
	category.Post("/category", c.CategoryController.Create)
	category.Put("/category/:categoryId/update", c.CategoryController.Update)
	category.Delete("/category/:categoryId/delete", c.CategoryController.Delete)
}
