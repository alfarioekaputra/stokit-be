package http

import (
	"stokit/external"
	"stokit/internal/model"
	"stokit/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoryController struct {
	Log             *logrus.Logger
	CategoryUsecase *usecase.CategoryUsecase
}

func NewCategoryController(categoryUsecase *usecase.CategoryUsecase, logger *logrus.Logger) *CategoryController {
	return &CategoryController{
		Log:             logger,
		CategoryUsecase: categoryUsecase,
	}
}

func (c *CategoryController) List(ctx *fiber.Ctx) error {
	filter := new(model.CategoryFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid filter"})
	}

	stdReq, err := external.ConvertFiberToHTTPRequest(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to convert request",
		})
	}

	categories, err := c.CategoryUsecase.List(stdReq, filter)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to fetch all categories")
		return err
	}

	return ctx.JSON(categories)
}

func (c *CategoryController) View(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")

	request := &model.ViewCategoryRequest{
		ID: categoryId,
	}

	response, err := c.CategoryUsecase.View(ctx.Context(), request)
	if err != nil {
		c.Log.Warnf("Category not found : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateCategoryRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.CategoryUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateCategoryRequest)
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("categoryId")

	response, err := c.CategoryUsecase.Update(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Error Update Category")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CategoryResponse]{Data: response})
}

func (c *CategoryController) Delete(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("categoryId")

	request := &model.DeleteCategoryRequest{
		ID: categoryId,
	}

	if err := c.CategoryUsecase.Delete(ctx.Context(), request); err != nil {
		c.Log.WithError(err).Error("Error Deleting Category")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
