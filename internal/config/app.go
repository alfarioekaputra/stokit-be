package config

import (
	"stokit/internal/delivery/http"
	"stokit/internal/delivery/http/middleware"
	"stokit/internal/delivery/http/route"
	"stokit/internal/repository"
	"stokit/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	productRepository := repository.NewProductRepository(config.Log)
	categoryRepository := repository.NewCategoryRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUsecase(config.DB, config.Log, config.Validate, userRepository)
	productUseCase := usecase.NewProductUsecase(config.DB, config.Log, config.Validate, productRepository)
	categoryUseCase := usecase.NewCategoryUsecase(config.DB, config.Log, config.Validate, categoryRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	productController := http.NewProductController(productUseCase, config.Log)
	categoryController := http.NewCategoryController(categoryUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                config.App,
		UserController:     userController,
		ProductController:  productController,
		CategoryController: categoryController,
		AuthMiddleware:     authMiddleware,
	}
	routeConfig.Setup()
}
