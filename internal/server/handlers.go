package server

import (
	"rest-company/internal/company/delivery/http"
	"rest-company/internal/company/repository"
	"rest-company/internal/company/usecase"
	"rest-company/internal/middleware"
)

func (s *Server) MapHandlers() {
	companyRepo := repository.NewCompanyRepository(s.redis)
	companyUseCase := usecase.NewCompanyUseCase(s.cfg, companyRepo, s.logger)
	companyHandlers := http.NewCompanyHandlers(s.cfg, companyUseCase, s.logger)

	v1 := s.app.Group("/api/v1")
	companies := v1.Group("/companies")
	companies.Use(middleware.CheckIsCyprus)
	http.MapCompaniesRoutes(companies, companyHandlers)
}
