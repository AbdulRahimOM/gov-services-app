package dberror

import "fmt"

var (
	ErrRecordNotFound error = fmt.Errorf("@db: record not found")
)