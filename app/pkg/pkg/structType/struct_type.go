package structType

import (
	"fmt"
	"reflect"
)

func isPointer(i interface{}) bool {
	return reflect.TypeOf(i).Kind() == reflect.Ptr
}

// RETURN 1 for json where, RETURN 2 for column
func GetType(dbType string, x interface{}, dst []string) ([]string, []string) {
	v := reflect.ValueOf(x)
	s := reflect.TypeOf(x)
	selected := []string{}
	selected2 := []string{}
	if isPointer(x) {
		return nil, nil
	}
	for i := 0; i < v.NumField(); i++ {
		for _, field := range dst {
			if field == s.Field(i).Tag.Get("json") {
				fieldType := s.Field(i).Type

				// Handle pointer types
				if fieldType.Kind() == reflect.Ptr {
					fieldType = fieldType.Elem()
				}

				if dbType == "postgres" {

					// PG
					switch fieldType.Kind() {
					case reflect.String:
						selected = append(selected, "(new_value->>'"+field+"')::text AS "+field)
					case reflect.Int64:
						selected = append(selected, "(new_value->>'"+field+"')::int AS "+field)
					default:
						selected = append(selected, "(new_value->>'"+field+"')::"+fieldType.Name()+" AS "+field)
					}
				} else if dbType == "mysql" {
					selected = append(selected, "JSON_VALUE(new_value, '$."+field+"')"+" AS "+field)
				} else {

					// MYSQL
					selected = append(selected, "JSON_VALUE(new_value, '$."+field+"') AS "+field)
				}
				selected2 = append(selected2, field)
				fmt.Println("Field", field, "TYPE ", s.Field(i).Type.Name())

			}
		}
	}

	return selected, selected2
}
