package usecase

import (
	"context"

	"rest-company/config"
	"rest-company/internal/company"
	"rest-company/internal/models"
	"rest-company/pkg/logger"
)

type companyUseCase struct {
	cfg    *config.Config
	repo   company.Repository
	logger logger.Logger
}

func NewCompanyUseCase(cfg *config.Config, repo company.Repository, logger logger.Logger) company.UseCase {
	return &companyUseCase{
		cfg:    cfg,
		repo:   repo,
		logger: logger,
	}
}

func (uc *companyUseCase) Create(ctx context.Context, company *models.Company) (*models.CompanyWithID, error) {
	return uc.repo.Create(ctx, company)
}

func (uc *companyUseCase) Update(ctx context.Context, id string, company *models.Company) error {
	return uc.repo.Update(ctx, &models.CompanyWithID{
		ID:      id,
		Name:    company.Name,
		Code:    company.Code,
		Country: company.Country,
		Website: company.Website,
		Phone:   company.Phone,
	})
}

func (uc *companyUseCase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *companyUseCase) GetByID(ctx context.Context, id string) (*models.CompanyWithID, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *companyUseCase) GetList(ctx context.Context, query models.CompanySearch) ([]*models.CompanyWithID, error) {
	return uc.repo.GetList(ctx, query)
}
