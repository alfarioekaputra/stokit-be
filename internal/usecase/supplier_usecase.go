package usecase

import (
	"context"
	"net/http"
	"stokit/external/helper"
	"stokit/internal/entity"
	"stokit/internal/model"
	"stokit/internal/model/converter"
	"stokit/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SupplierUsecase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	SupplierRepository *repository.SupplierRepository
}

func NewSupplierUsecase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, supplierRepository *repository.SupplierRepository) *SupplierUsecase {
	return &SupplierUsecase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		SupplierRepository: supplierRepository,
	}
}

func (c *SupplierUsecase) List(req *http.Request, filter *model.SupplierFilter) (*model.PaginatedResponse[*model.SupplierResponse], error) {
	raw, err := repository.FetchAllWithFilter[entity.Supplier](
		c.DB,
		req,
		filter,
		repository.ApplySupplierFilter,
		helper.Preloads("Products"),
	)
	if err != nil {
		c.Log.Warnf("failed fetch user: %+v", err)
		return nil, fiber.ErrNotFound
	}

	var suppliers []*model.SupplierResponse
	for _, supplier := range raw.Items {
		supplierResponse := converter.SupplierToResponse(&supplier)
		suppliers = append(suppliers, supplierResponse)
	}

	return &model.PaginatedResponse[*model.SupplierResponse]{
		Items: suppliers,
		Page:  raw.Page,
		Size:  raw.Size,
		Total: raw.Total,
		First: raw.First,
		Last:  raw.Last,
	}, nil
}

func (c *SupplierUsecase) Create(ctx context.Context, request *model.CreateSupplierRequest) (*model.SupplierResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	supplier := &entity.Supplier{
		ID:      uuid.New().String(),
		Name:    request.Name,
		Address: request.Address,
		Phone:   request.Phone,
		Email:   request.Email,
	}

	if err := c.SupplierRepository.Create(tx, supplier); err != nil {
		c.Log.Warnf("Failed create supplier to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.SupplierToResponse(supplier), nil
}

func (c *SupplierUsecase) View(ctx context.Context, request *model.ViewSupplierRequest) (*model.SupplierResponse, error) {
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request bidy : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	supplier := new(entity.Supplier)
	if err := c.SupplierRepository.FindById(c.DB, supplier, request.ID); err != nil {
		c.Log.Warnf("Supplier Not Found")
		return nil, fiber.ErrNotFound
	}

	return converter.SupplierToResponse(supplier), nil
}

func (c *SupplierUsecase) Update(ctx context.Context, request *model.UpdateSupplierRequest) (*model.SupplierResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	supplier := new(entity.Supplier)
	if err := c.SupplierRepository.FindById(tx, supplier, request.ID); err != nil {
		c.Log.Warnf("Failed find supplier by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	supplier.Name = request.Name
	supplier.Address = request.Address
	supplier.Phone = request.Phone
	supplier.Email = request.Email

	if err := c.SupplierRepository.Update(tx, supplier); err != nil {
		c.Log.Warnf("Failed save supplier : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.SupplierToResponse(supplier), nil
}

func (c *SupplierUsecase) Delete(ctx context.Context, request *model.DeleteSupplierRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return fiber.ErrBadRequest
	}

	supplier := new(entity.Supplier)
	if err := c.SupplierRepository.FindById(tx, supplier, request.ID); err != nil {
		c.Log.Warnf("Failed find supplier by id : %+v", err)
		return fiber.ErrNotFound
	}

	supplier.ID = request.ID

	if err := c.SupplierRepository.Delete(tx, supplier); err != nil {
		c.Log.Warnf("Failed delete supplier : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}
