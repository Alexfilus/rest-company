package models

import (
	"fmt"
	"strings"
)

type Company struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type CompanyWithID struct {
	ID      string `json:"id" redis:",key"`
	Ver     int64  `redis:",ver"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

const defaultSearchLimit = 20

type CompanySearch struct {
	Limit   int64  `json:"limit" query:"limit" default:"20"`
	Offset  int64  `json:"offset" query:"offset"`
	Name    string `json:"name,omitempty" query:"name,omitempty"`
	Code    string `json:"code,omitempty" query:"code,omitempty"`
	Country string `json:"country,omitempty" query:"country,omitempty"`
	Website string `json:"website,omitempty" query:"website,omitempty"`
	Phone   string `json:"phone,omitempty" query:"phone,omitempty"`
}

func (c *CompanySearch) String() string {
	if c.Limit == 0 {
		c.Limit = defaultSearchLimit
	}
	parts := make([]string, 0, 5)
	if c.Name != "" {
		parts = append(parts, fmt.Sprintf("@%s:{%s}", "name", c.Name))
	}
	if c.Code != "" {
		parts = append(parts, fmt.Sprintf("@%s:{%s}", "code", c.Code))
	}
	if c.Country != "" {
		parts = append(parts, fmt.Sprintf("@%s:{%s}", "country", c.Country))
	}
	if c.Website != "" {
		parts = append(parts, fmt.Sprintf("@%s:{%s}", "website", c.Website))
	}
	if c.Phone != "" {
		parts = append(parts, fmt.Sprintf("@%s:{%s}", "phone", c.Phone))
	}
	if len(parts) == 0 {
		return "*"
	}

	return strings.Join(parts, " ")
}
