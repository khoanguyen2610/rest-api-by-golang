package filter

import "user-service/repositories"

type basicFilter struct {
	where repositories.Where
	keys  map[string]bool
}

func BasicFilter() *basicFilter {
	return &basicFilter{
		where: repositories.Where{},
		keys:  map[string]bool{},
	}
}

// implement repository.Filter interface
func (f *basicFilter) GetLimit() int {
	return 0
}

// implement repository.Filter interface
func (f *basicFilter) GetOffset() int {
	return 0
}

// implement repository.Filter interface
func (f *basicFilter) GetWhere() repositories.Where {
	return f.where
}

func (f *basicFilter) AddWhere(key string, query string, values ...interface{}) *basicFilter {
	f.where[query] = values
	f.keys[key] = true
	return f
}
