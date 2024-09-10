package getfilter

import "regexp"

type QueryField struct {
	Operator string
	Value    string
}

func ArrQuery(queryString string) map[string][]QueryField {
	regex := regexp.MustCompile(`(\w+):([^|]+):(\w+)`)
	matches := regex.FindAllStringSubmatch(queryString, -1)

	arrQuery := make(map[string][]QueryField)

	for _, match := range matches {
		field := match[1]
		value := match[2]
		operator := match[3]

		symbol, exists := QueryParserOperators[operator]
		if !exists {
			symbol = operator
		}

		queryField := QueryField{
			Operator: symbol,
			Value:    value,
		}

		// Check if the field already exists in the map
		if _, found := arrQuery[field]; found {
			// If found, append the new QueryField to the existing slice
			arrQuery[field] = append(arrQuery[field], queryField)
		} else {
			// If not found, create a new slice with the current QueryField
			arrQuery[field] = []QueryField{queryField}
		}
	}

	return arrQuery
}

func ArrSort(queryString string) map[string]QueryField {
	regex := regexp.MustCompile(`(\w+):([^|]+)`)
	matches := regex.FindAllStringSubmatch(queryString, -1)

	arrSort := make(map[string]QueryField)

	for _, match := range matches {
		field := match[1]
		value := match[2]

		arrSort[field] = QueryField{
			Operator: field,
			Value:    value,
		}
	}

	return arrSort
}
