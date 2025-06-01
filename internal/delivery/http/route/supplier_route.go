package route

func SetupSupplierRoutes(c *RouteConfig) {
	supplier := c.App.Group("/api", c.AuthMiddleware)
	supplier.Get("/supplier", c.SupplierController.List)
	supplier.Get("/supplier/:supplierId/view", c.SupplierController.View)
	supplier.Post("/supplier", c.SupplierController.Create)
	supplier.Put("/supplier/:supplierId/update", c.SupplierController.Update)
	supplier.Delete("/supplier/:supplierId/delete", c.SupplierController.Delete)
}
