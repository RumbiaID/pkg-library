package sortfilter

import (
	"github.com/gin-gonic/gin"
)

type SortItem struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type SortMiddleware struct {
	ArrQuery []SortItem
	Message  string
}

func Initiate(c *gin.Context) []SortItem {
	sortmiddleware := &SortMiddleware{}

	sort, exists := c.GetQuery("sort")
	if exists {
		sortItems := make([]SortItem, 0)
		for field, q := range ArrQuery(sort) {
			sortItems = append(sortItems, SortItem{
				Field: field,
				Value: q.Value,
			})
		}
		sortmiddleware.ArrQuery = sortItems
	}

	return sortmiddleware.ArrQuery
}

func Validation(query []SortItem) bool {
	for _, item := range query {
		if !contains(item.Value, QueryParserSymbols) {
			return false
		}
	}
	return true
}

func Handle(c *gin.Context) bool {
	if Validation(Initiate(c)) {
		c.Set("sortQuery", Initiate(c))
		return false
	}
	return true
}

func contains(item string, arr []string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
