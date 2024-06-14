package getfilter

var QueryParserSymbols = []string{
	"=",
	"<",
	">",
	"<=",
	">=",
	"in",
	"like",
	"is",
	"not",
}

var QueryParserOperators = map[string]string{
	"eq":       "=",
	"lt":       "<",
	"gt":       ">",
	"lte":      "<=",
	"gte":      ">=",
	"in":       "in",
	"like":     "like",
	"contains": "like",
	"is":       "is",
	"not":      "not in",
}
