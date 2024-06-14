package dbutils

import "gorm.io/gorm"

// ValidatePageParam validates the page and pageSize parameters.
// If pageSize is less than 0, it sets pageSize to -1 and page to 1.
// If page is less than 1, it sets page to 1.
// It returns the validated page and pageSize.
func ValidatePageParam(page int, pageSize int) (int, int) {
	if pageSize < 0 {
		pageSize = -1
		page = 1
	}
	if page < 1 {
		page = 1
	}
	return page, pageSize
}

// Paginate is a function that returns a function which applies pagination to a gorm.DB instance.
// If pageSize is -1, it returns the db instance without any modification.
// Otherwise, it applies the Offset and Limit methods to the db instance based on the page and pageSize parameters.
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageSize == -1 {
			return db
		}
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

// CountTotalPage calculates the total number of pages based on the pageSize and totalData parameters.
// If totalData is 0, it returns 0.
// If pageSize is -1, it returns 1.
// Otherwise, it calculates the total number of pages and increments it by 1 if there are remaining data after dividing totalData by pageSize.
func CountTotalPage(pageSize int, totalData int64) int64 {
	if totalData == 0 {
		return 0
	} else if pageSize == -1 {
		return 1
	} else {
		totalPage := totalData / int64(pageSize)
		if totalData%int64(pageSize) > 0 {
			totalPage++
		}
		return totalPage
	}
}

// ResultListData is a generic struct that represents the result of a paginated list data.
// It contains the current page, page size, total number of pages, total number of data per page, total number of data, and the actual data.
type ResultListData[T any] struct {
	Page             int   `json:"page"`                // The current page
	PageSize         int   `json:"limit"`               // The size of the page
	TotalPage        int64 `json:"total_page"`          // The total number of pages
	TotalDataPerPage int64 `json:"total_data_per_page"` // The total number of data per page
	TotalData        int64 `json:"total_data"`          // The total number of data
	Data             T     `json:"data"`                // The actual data
}

type ResultListData2[T any] struct {
	Page      int   `json:"current_page"` // The current page
	Count     int   `json:"count"`        // The count of listed data in the page
	TotalPage int64 `json:"total_page"`   // The total number of pages
	PerPage   int64 `json:"per_page"`     // The total number of data per page
	Total     int64 `json:"total"`        // The total number of data
	Data      T     `json:"data"`         // The actual data
}
