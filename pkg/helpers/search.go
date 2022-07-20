package helpers

import "fmt"

// LikeStatement adds % to search clause
func LikeStatement(search string) string {
	return search + "%"
}

// DynamicWhereBuilder builds dynamic where clause for search using queries
func DynamicWhereLikeBuilder(query map[string]string, tableNameJoin string) ([]interface{}, []string, []string) {
	var values []interface{}
	var where []string
	var whereJoin []string

	for index, value := range query {
		if value == "" {
			continue
		}
		values = append(values, LikeStatement(value))
		where = append(where, fmt.Sprintf("%s LIKE ?", index))
		whereJoin = append(whereJoin, fmt.Sprintf("%s.%s LIKE ?", tableNameJoin, index))
	}
	return values, where, whereJoin
}
