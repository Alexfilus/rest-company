package models

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type TestCase struct {
	Name   string
	Input  CompanySearch
	Output string
}

func TestCompanySearch(t *testing.T) {
	cases := []TestCase{
		{
			Name:   "empty query",
			Input:  CompanySearch{},
			Output: "* limit 0 20",
		},
		{
			Name: "one field",
			Input: CompanySearch{
				Name: "Google",
			},
			Output: "@name:{Google} limit 0 20",
		},
		{
			Name: "complex query",
			Input: CompanySearch{
				Name:    "Google",
				Country: "USA",
				Phone:   "123456789",
				Limit:   5,
				Offset:  10,
			},
			Output: "@name:{Google} @country:{USA} @phone:{123456789} limit 10 5",
		},
	}
	for _, testCase := range cases {
		assert.Equal(t, testCase.Input.String(), testCase.Output, testCase.Name)
	}
}
