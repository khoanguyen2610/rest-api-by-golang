package filter

import (
	"strings"
)

type basicOrder struct {
	OrderBy  string `json:"order_by" schema:"order_by"`
	OrderDir string `json:"order_dir" schema:"order_dir"`
}

func BasicOrder() *basicOrder {
	return &basicOrder{}
}

func (s *basicOrder) GetOrderBy(allowedFields map[string]string, defaultSort string) string {
	for orderParam, orderField := range allowedFields {
		if orderParam == s.OrderBy {
			return orderField
		}
	}

	return allowedFields[defaultSort]
}

func (s *basicOrder) GetOrderDir() string {
	for _, orderDir := range []string{"ASC", "DESC"} {
		if strings.ToUpper(s.OrderDir) == orderDir {
			return orderDir
		}
	}

	return "DESC"
}

func (s *basicOrder) GetOrder(allowedFields map[string]string, defaultSort string) string {
	orderDir := s.GetOrderDir()
	orderBy := s.GetOrderBy(allowedFields, defaultSort)

	return orderBy + " " + orderDir
}
