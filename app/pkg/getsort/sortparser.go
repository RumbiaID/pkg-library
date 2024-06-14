package sortfilter

import "regexp"

type QueryField struct {
	Value string
}

func ArrQuery(queryString string) map[string]QueryField {
	regex := regexp.MustCompile(`(\w+):(\w+)`)
	matches := regex.FindAllStringSubmatch(queryString, -1)

	arrQuery := make(map[string]QueryField)

	for _, match := range matches {
		field := match[1]
		value := match[2]

		symbol, exists := QueryParserOperators[value]
		if !exists {
			symbol = value
		}

		arrQuery[field] = QueryField{
			Value: symbol,
		}
	}

	return arrQuery
}
