package usecase

import (
	"context"
	"net/http"
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

type CategoryUsecase struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CategoryRepository *repository.CategoryRepository
}

func NewCategoryUsecase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, categoryRepository *repository.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		CategoryRepository: categoryRepository,
	}
}

func (c *CategoryUsecase) List(req *http.Request, filter *model.CategoryFilter) (*model.PaginatedResponse[*model.CategoryResponse], error) {
	raw, err := repository.FetchAllWithFilter[entity.Category](
		c.DB,
		req,
		filter,
		repository.ApplyCategoryFilter,
	)
	if err != nil {
		c.Log.Warnf("failed fetch user: %+v", err)
		return nil, fiber.ErrNotFound
	}

	var categories []*model.CategoryResponse
	for _, category := range raw.Items {
		categoryResponse := converter.CategoryToResponse(&category)
		categories = append(categories, categoryResponse)
	}

	return &model.PaginatedResponse[*model.CategoryResponse]{
		Items: categories,
		Page:  raw.Page,
		Size:  raw.Size,
		Total: raw.Total,
		First: raw.First,
		Last:  raw.Last,
	}, nil
}

func (c *CategoryUsecase) Create(ctx context.Context, request *model.CreateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	category := &entity.Category{
		ID:       uuid.New().String(),
		Name:     request.Name,
		ParentID: request.ParentID,
	}

	if err := c.CategoryRepository.Create(tx, category); err != nil {
		c.Log.Warnf("Failed create category to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUsecase) View(ctx context.Context, request *model.ViewCategoryRequest) (*model.CategoryResponse, error) {
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request bidy : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(c.DB, category, request.ID); err != nil {
		c.Log.Warnf("Category Not Found")
		return nil, fiber.ErrNotFound
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUsecase) Update(ctx context.Context, request *model.UpdateCategoryRequest) (*model.CategoryResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		c.Log.Warnf("Failed find category by id : %+v", err)
		return nil, fiber.ErrNotFound
	}

	category.Name = request.Name
	category.ParentID = request.ParentID

	if err := c.CategoryRepository.Update(tx, category); err != nil {
		c.Log.Warnf("Failed save category : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (c *CategoryUsecase) Delete(ctx context.Context, request *model.DeleteCategoryRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return fiber.ErrBadRequest
	}

	category := new(entity.Category)
	if err := c.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		c.Log.Warnf("Failed find category by id : %+v", err)
		return fiber.ErrNotFound
	}

	category.ID = request.ID

	if err := c.CategoryRepository.Delete(tx, category); err != nil {
		c.Log.Warnf("Failed delete category : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}
