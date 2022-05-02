//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package company

import (
	"context"

	"rest-company/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, company *models.Company) (*models.CompanyWithID, error)
	Update(ctx context.Context, id string, company *models.Company) error
	GetByID(ctx context.Context, id string) (*models.CompanyWithID, error)
	Delete(ctx context.Context, id string) error
	GetList(ctx context.Context, query models.CompanySearch) ([]*models.CompanyWithID, error)
}
