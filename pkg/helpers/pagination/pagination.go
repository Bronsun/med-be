package pagination

import (
	"math"
)

type Param struct {
	Page   int64
	Limit  int64
	Offset int64
}

type Result struct {
	TotalRecord int64       `json:"total_record"`
	TotalPage   int64       `json:"total_page"`
	Offset      int64       `json:"offset"`
	Limit       int64       `json:"limit"`
	Page        int64       `json:"page"`
	PrevPage    int64       `json:"prev_page"`
	NextPage    int64       `json:"next_page"`
	Data        interface{} `json:"data"`
}

// BuildPaginationQuery build offset and checks limit
func BuildPaginationQuery(page, limit int64) (int64, int64) {
	var offset int64

	if page < 1 {
		page = 1
	}

	if limit == 0 {
		limit = 25
	}
	if limit > 50 {
		limit = 50
	}

	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	return offset, limit
}

// PaginationResponseBuilder build pagination response for clinics
func PaginationResponseBuilder(param Param, resultData interface{}, count int64) *Result {
	var result Result

	result.TotalRecord = count
	result.Data = resultData
	result.Page = param.Page

	result.Offset = param.Offset
	result.Limit = param.Limit
	result.TotalPage = int64(math.Ceil(float64(count) / float64(param.Limit)))

	if param.Page > 1 {
		result.PrevPage = param.Page - 1
	} else {
		result.PrevPage = param.Page
	}

	if param.Page == result.TotalPage {
		result.NextPage = param.Page
	} else {
		result.NextPage = param.Page + 1
	}
	return &result
}
