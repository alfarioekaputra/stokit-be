package route

func SetupUserRoutes(c *RouteConfig) {
	// Auth
	user := c.App.Group("/api", c.AuthMiddleware)
	user.Delete("/users", c.UserController.Logout)
	user.Patch("/users/_current", c.UserController.Update)
	user.Get("/users/_current", c.UserController.Current)
	user.Get("/users", c.UserController.FetchAll)
}
