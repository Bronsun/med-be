package helpers

import (
	"fmt"
)

// LikeStatement adds % to search clause
func LikeStatement(search string) string {
	return search + "%"
}

// BuildQueryWithPagination
func BuildQueryWithPagination(sql string, limit, offset int64) string {
	return fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, limit, offset)
}

// BuildQuery builcx query
func BuildQuery(query map[string]string, benefit string) string {
	var res string = "SELECT clinics.*"
	var flag = false

	if benefit != "" {
		res += ", clinic_benefits.awaiting,clinic_benefits.visit_date,clinic_benefits.average_period FROM clinics FULL OUTER JOIN clinic_benefits on clinic_benefits.clinic_id = clinics.id WHERE clinic_benefits.benefit_id = (SELECT id FROM benefits WHERE name = " + "'" + benefit + "'" + ")"
		flag = true
	} else {
		res += " FROM clinics"
	}

	for key, value := range query {
		if value == "" {
			continue
		}
		if key == "benefits_for_children" {
			res += decideAndOrWhere(&flag)
			res += key + " = " + value
		} else {
			res += decideAndOrWhere(&flag)
			res += key + " LIKE " + "'" + LikeStatement(value) + "'"
		}
	}
	return res
}

func decideAndOrWhere(flag *bool) string {
	if *flag {
		return " AND "
	}
	*flag = true
	return " WHERE "
}
