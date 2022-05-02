package repository

import (
	"context"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/om"

	"rest-company/internal/company"
	"rest-company/internal/models"
)

const indexName = "hashidx:companies"
const prefix = "companies"

type companyRepo struct {
	repo om.Repository
}

func (c companyRepo) Create(ctx context.Context, company *models.Company) (*models.CompanyWithID, error) {
	record := c.repo.NewEntity().(*models.CompanyWithID)
	record.Name = company.Name
	record.Code = company.Code
	record.Country = company.Country
	record.Website = company.Website
	record.Phone = company.Phone
	err := c.repo.Save(ctx, record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (c companyRepo) Update(ctx context.Context, company *models.CompanyWithID) error {
	return c.repo.Save(ctx, company)
}

func (c companyRepo) GetByID(ctx context.Context, id string) (*models.CompanyWithID, error) {
	record, err := c.repo.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return record.(*models.CompanyWithID), nil
}

func (c companyRepo) Delete(ctx context.Context, id string) error {
	return c.repo.Remove(ctx, id)
}

func (c companyRepo) GetList(ctx context.Context, query models.CompanySearch) ([]*models.CompanyWithID, error) {
	_, records, err := c.repo.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query(query.String()).Limit().OffsetNum(query.Offset, query.Limit).Build()
	})
	if err != nil {
		return nil, err
	}
	return records.([]*models.CompanyWithID), nil
}
func initDB(db rueidis.Client) {
	db.Do(context.Background(),
		db.B().FtCreate().Index(indexName).OnHash().Prefix(1).Prefix(prefix).Schema().
			FieldName("name").Tag().Sortable().
			FieldName("code").Tag().Sortable().
			FieldName("country").Tag().Sortable().
			FieldName("website").Tag().Sortable().
			FieldName("phone").Tag().Sortable().
			Build(),
	)
}

func NewCompanyRepository(db rueidis.Client) company.Repository {
	initDB(db)
	repo := om.NewHashRepository(prefix, models.CompanyWithID{}, db)
	return &companyRepo{repo: repo}
}
