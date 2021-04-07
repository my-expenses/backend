package pagination

import (
	paginationData "backend/models/pagination"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ExtractPaginationData(c *echo.Context) *paginationData.Data {
	queryParams := (*c).Request().URL.Query()
	sortDesc := convertToBoolArray(queryParams["sortDesc[]"])
	sortBy := queryParams["sortBy[]"]
	page, _ := strconv.Atoi(queryParams["page"][0])
	itemsPerPage, _ := strconv.Atoi(queryParams["itemsPerPage"][0])
	return &paginationData.Data{
		SortDesc:     sortDesc,
		SortBy:       sortBy,
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}
}

func convertToBoolArray(strArr []string) []bool {
	length := len(strArr)
	boolArr := make([]bool, length)
	for i := 0; i < length; i++ {
		val, err := strconv.ParseBool(strArr[i])
		if err != nil {
			boolArr[i] = false
		} else {
			boolArr[i] = val
		}
	}
	return boolArr
}