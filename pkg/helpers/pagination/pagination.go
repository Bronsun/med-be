package pagination

import (
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
)

type Param struct {
	Page      int64
	Limit     int64
	Offset    int64
	OrderBy   string
	Where     []string
	Values    []interface{}
	BenefitID string
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
func PaginationResponseBuilder(db *gorm.DB, param Param, counter string, resultData interface{}) *Result {

	done := make(chan bool, 1)
	var result Result
	var count int64

	switch counter {
	case "clinicsAll":
		go countAllClinics(db, done, &count)
	case "clinicsWithWhere":
		go countClinicsWithDynamicWhereClause(db, done, param.Values, param.Where, &count)
	case "clinicsBenefits":
		go countClinicsBenefitsWithDynamicWhereClause(db, done, param.BenefitID, param.Values, param.Where, &count)
	}

	<-done
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

func countAllClinics(db *gorm.DB, done chan bool, count *int64) {
	db.Table("clinics").Count(count)
	done <- true
}

func countClinicsWithDynamicWhereClause(db *gorm.DB, done chan bool, values []interface{}, where []string, count *int64) {
	db.Raw("SELECT COUNT (*) FROM clinics WHERE "+strings.Join(where, " AND "), values...).Count(count)

	done <- true
}

func countClinicsBenefitsWithDynamicWhereClause(db *gorm.DB, done chan bool, bID string, values []interface{}, where []string, count *int64) {
	if where != nil {
		db.Raw("SELECT clinics.*, clinic_benefits.visit_date FROM clinics FULL OUTER JOIN clinic_benefits on clinic_benefits.clinic_id = clinics.id WHERE clinic_benefits.benefit_id = "+bID+" AND "+strings.Join(where, " AND "), values...).Count(count)
		done <- true
	} else {
		db.Raw("SELECT clinics.*, clinic_benefits.visit_date FROM clinics FULL OUTER JOIN clinic_benefits on clinic_benefits.clinic_id = clinics.id WHERE clinic_benefits.benefit_id = " + bID).Count(count)

		fmt.Println(*count)
		done <- true
	}

}
