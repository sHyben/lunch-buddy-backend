package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

// Offset returns the starting number of result for pagination
// It returns 0 if the offset is not a number
func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

// Limit returns the number of result for pagination
// It returns 25 if the limit is not a number
// It returns 100 if the limit is greater than 100
// It returns 1 if the limit is less than 1
// It returns the limit if it is between 1 and 100
func Limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}
	return limitInt
}

// SortOrder returns the string for sorting and ordering data
// It returns the table name, the snake cased sort and the snake cased order
func SortOrder(table, sort, order string) string {
	return table + "." + ToSnakeCase(sort) + " " + ToSnakeCase(order)
}

// ToSnakeCase changes string to database table
// It returns the snake cased string
// It is used to convert camel case to snake case
// It is used to convert the sort and order parameters to database table
func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}
