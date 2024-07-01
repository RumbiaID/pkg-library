package structType

import (
	"fmt"
	"github.com/RumbiaID/pkg-library/app/pkg/constants"
	"github.com/RumbiaID/pkg-library/app/pkg/loggingdata"
	"reflect"
	"time"
)

// RETURN 1 for json where, RETURN 2 for column
func GetType(dbType string, x interface{}, dst []string) ([]string, []string) {
	v := reflect.ValueOf(x)
	s := reflect.TypeOf(x)
	selected := []string{}
	selected2 := []string{}

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
						selected2 = append(selected2, field+"::text")
					case reflect.Int64:
						selected = append(selected, "(new_value->>'"+field+"')::int AS "+field)
						selected2 = append(selected2, field+"::int")
					default:
						selected = append(selected, "(new_value->>'"+field+"')::"+fieldType.Name()+" AS "+field)
						selected2 = append(selected2, field+"::"+fieldType.Name())
					}
				} else if dbType == "mysql" {
					selected = append(selected, "JSON_VALUE(new_value, '$."+field+"') AS "+field)
					selected2 = append(selected2, field)
				} else {

					// MYSQL
					selected = append(selected, "JSON_VALUE(new_value, '$."+field+"') AS "+field)
					selected2 = append(selected2, field)
				}
				fmt.Println("Field", field, "TYPE ", s.Field(i).Type.Name())

			}
		}
	}

	return selected, selected2
}

func DeclarePendingInsert(x interface{}, requestHeader *loggingdata.InsertReturn) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()
	now := time.Now()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status":
			field.SetInt(constants.SYSROW_STATUS_PENDING_INSERT) // Replace with appropriate constant
		case "sys_created_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_created_host":
			field.SetString(requestHeader.CreatedHost)
		case "sys_last_pending_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_last_pending_host":
			field.SetString(requestHeader.CreatedHost)
		case "sys_last_pending_time":
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.ValueOf(&now))
			}
		}
	}
}

func DeclarePendingUpdate(x interface{}, requestHeader *loggingdata.InsertReturn, pendingId *int) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()
	now := time.Now()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status":
			field.SetInt(constants.SYSROW_STATUS_PENDING_UPDATE) // Replace with appropriate constant
		case "sys_last_pending_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_last_approve_by":
			field.SetZero()
		case "sys_last_approve_host":
			field.SetZero()
		case "sys_last_approval_notes":
			field.SetZero()
		case "sys_last_pending_host":
			field.SetString(requestHeader.CreatedHost)
		case "sys_last_pending_time":
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.ValueOf(&now))
			}
		case "sys_last_approve_time":

			field.SetZero()

		case "pending_id":
			{
				if field.Kind() == reflect.Ptr {
					switch field.Type().Elem().Kind() {
					case reflect.Int:
						value := *pendingId
						field.Set(reflect.ValueOf(&value))
					case reflect.Int64:
						value := int64(*pendingId)
						field.Set(reflect.ValueOf(&value))
					}
				}
			}
		}
	}
}

func DeclarePendingDelete(x interface{}, requestHeader *loggingdata.InsertReturn) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()
	now := time.Now()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status":
			field.SetInt(constants.SYSROW_STATUS_PENDING_DELETE) // Replace with appropriate constant
		case "sys_last_pending_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_last_approve_by":
			field.SetZero()
		case "sys_last_approve_host":
			field.SetZero()
		case "sys_last_approval_notes":
			field.SetZero()
		case "sys_last_pending_host":
			field.SetString(requestHeader.CreatedHost)
		case "sys_last_pending_time":
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.ValueOf(&now))
			}
		case "sys_last_approve_time":
			field.SetZero()

		case "pending_id":
			field.SetZero()
		}
	}
}

func DeclareApproveUpsert(x interface{}, requestHeader *loggingdata.InsertReturn, remarks string) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()
	now := time.Now()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status":
			field.SetInt(constants.SYSROW_STATUS_ACTIVE) // Replace with appropriate constant
		case "sys_last_approve_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_last_approve_host":
			field.SetString(requestHeader.CreatedHost)
		case "sys_last_approval_notes":
			field.SetString(remarks)
		case "sys_last_pending_time":
			if field.Kind() == reflect.Ptr {
				field.SetZero()
			}
		case "sys_last_approve_time":
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.ValueOf(&now))
			}
		case "pending_id":
			field.SetZero()
		}

	}
}

func DeclareReturnInsert(x interface{}, requestHeader *loggingdata.InsertReturn, remarks string) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status":
			field.SetInt(constants.SYSROW_STATUS_RETURN_INSERT) // Replace with appropriate constant
		case "sys_last_approval_notes":
			field.SetString(remarks)
		}
	}
}

func DeclareReturnUpdate(x interface{}, requestHeader *loggingdata.InsertReturn, remarks string) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status", "row_status":
			field.SetInt(constants.SYSROW_STATUS_RETURN_UPDATE) // Replace with appropriate constant
		case "sys_last_approval_notes", "return_notes":
			field.SetString(remarks)
		}
	}
}

func DeclareRejectDelUp(x interface{}, requestHeader *loggingdata.InsertReturn, remarks string) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()
	now := time.Now()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status":
			field.SetInt(constants.SYSROW_STATUS_ACTIVE) // Replace with appropriate constant
		case "sys_last_approve_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_last_approve_time":
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.ValueOf(&now))
			}
		case "sys_last_approve_host":
			field.SetString(requestHeader.CreatedHost)
		case "sys_last_approval_notes":
			field.SetString(remarks)
		case "pending_id":
			field.SetZero()
		}

	}
}

func DeclareRetryInsert(x interface{}, requestHeader *loggingdata.InsertReturn) {
	body := reflect.ValueOf(x).Elem()
	bodyType := reflect.TypeOf(x).Elem()
	now := time.Now()

	for i := 0; i < body.NumField(); i++ {
		field := body.Field(i)
		fieldType := bodyType.Field(i)

		switch fieldType.Tag.Get("json") {
		case "sys_row_status":
			field.SetInt(constants.SYSROW_STATUS_PENDING_INSERT) // Replace with appropriate constant
		case "sys_created_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_created_host":
			field.SetString(requestHeader.CreatedHost)
		case "sys_last_pending_by":
			field.SetString(requestHeader.CreatedBy)
		case "sys_last_pending_time":
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.ValueOf(&now))
			}
		case "sys_last_pending_host":
			field.SetString(requestHeader.CreatedHost)
		}
	}
}
