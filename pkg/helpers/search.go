package helpers

// LikeStatement adds % to search clause
func LikeStatement(search string) string {
	return search + "_%"
}
