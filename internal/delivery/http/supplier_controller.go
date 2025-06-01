package http

import (
	"stokit/external/helper"
	"stokit/internal/model"
	"stokit/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SupplierController struct {
	Log             *logrus.Logger
	SupplierUsecase *usecase.SupplierUsecase
}

func NewSupplierController(useCase *usecase.SupplierUsecase, logger *logrus.Logger) *SupplierController {
	return &SupplierController{
		Log:             logger,
		SupplierUsecase: useCase,
	}
}

func (c *SupplierController) List(ctx *fiber.Ctx) error {
	filter := new(model.SupplierFilter)
	if err := ctx.QueryParser(filter); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid filter"})
	}

	stdReq, err := helper.ConvertFiberToHTTPRequest(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to convert request",
		})
	}

	suppliers, err := c.SupplierUsecase.List(stdReq, filter)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to fetch all suppliers")
		return err
	}

	return ctx.JSON(suppliers)
}

func (c *SupplierController) View(ctx *fiber.Ctx) error {
	supplierId := ctx.Params("supplierId")

	request := &model.ViewSupplierRequest{
		ID: supplierId,
	}

	response, err := c.SupplierUsecase.View(ctx.Context(), request)
	if err != nil {
		c.Log.Warnf("Supplier not found : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.SupplierResponse]{Data: response})
}

func (c *SupplierController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateSupplierRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.SupplierUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.SupplierResponse]{Data: response})
}

func (c *SupplierController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateSupplierRequest)
	if err := ctx.BodyParser(&request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("supplierId")

	response, err := c.SupplierUsecase.Update(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("Error Update Category")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.SupplierResponse]{Data: response})
}

func (c *SupplierController) Delete(ctx *fiber.Ctx) error {
	supplierId := ctx.Params("supplierId")

	request := &model.DeleteSupplierRequest{
		ID: supplierId,
	}

	if err := c.SupplierUsecase.Delete(ctx.Context(), request); err != nil {
		c.Log.WithError(err).Error("Error Deleting Supplier")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
