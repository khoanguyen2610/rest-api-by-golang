package filter

const (
	defaultPerPage = 50
	maxPerPage     = 200
)

type paginationFilter struct {
	basicFilter
	basicOrder
	Page    int `json:"page" schema:"page"`
	PerPage int `json:"per_page" schema:"per_page"`
}

func PaginationFilter() *paginationFilter {
	return &paginationFilter{
		basicFilter: *BasicFilter(),
		basicOrder:  *BasicOrder(),
	}
}

// implement repository.Filter interface
func (f *paginationFilter) GetLimit() int {
	return f.GetPerPage() + 1
}

// implement repository.Filter interface
func (f *paginationFilter) GetOffset() int {
	return (f.GetPage() - 1) * f.GetPerPage()
}

func (f *paginationFilter) GetPage() int {
	if f.Page < 1 {
		return 1
	}
	return f.Page
}

func (f *paginationFilter) GetPerPage() int {
	if f.PerPage < 1 || f.PerPage > maxPerPage {
		return defaultPerPage
	}

	return f.PerPage
}
