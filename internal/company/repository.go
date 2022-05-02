package company

import (
	"context"

	"rest-company/internal/models"
)

type Repository interface {
	Create(ctx context.Context, company *models.Company) (*models.CompanyWithID, error)
	Update(ctx context.Context, company *models.CompanyWithID) error
	GetByID(ctx context.Context, id string) (*models.CompanyWithID, error)
	Delete(ctx context.Context, id string) error
	GetList(ctx context.Context, query models.CompanySearch) ([]*models.CompanyWithID, error)
}
