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
