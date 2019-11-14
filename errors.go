package main

import "fmt"

// Represents a error, which occurred trying to query the game list from IGDB.
type QueryError struct {
	Title string `json:"title"`
	Status int `json:"status"`
	Err error
}

// Returns a string representation of the current QueryError.
func (q QueryError) Error() string {
	return fmt.Sprintf("query error: (%d) %s", q.Status, q.Title)
}

// Creates a new QueryError using the specified information.
func NewQueryError(statusCode int, title string, err error) QueryError {
	return QueryError{
		Title:  title,
		Status: statusCode,
		Err:    err,
	}
}