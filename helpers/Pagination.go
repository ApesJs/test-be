package helpers

import (
	"github.com/ApesJs/test-be/initializers"
	"math"
)

type PaginationData struct {
	CurrentPage, NextPage, PreviousPage, TotalPages, Offset int
}

func GetPaginationData(page, perPage int, model interface{}) PaginationData {
	var totalRows int64
	initializers.DB.Model(model).Count(&totalRows)
	totalPages := math.Ceil(float64(totalRows / int64(perPage)))

	offset := (page - 1) * perPage

	return PaginationData{
		CurrentPage:  page,
		NextPage:     page + 1,
		PreviousPage: page - 1,
		TotalPages:   int(totalPages),
		Offset:       offset,
	}
}
