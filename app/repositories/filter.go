package repositories

type Where map[string][]interface{}
type Filter interface {
	GetWhere() Where
	GetLimit() int
	GetOffset() int
	//GetOrder(allowedFields map[string]string, defaultSort string) string
}
