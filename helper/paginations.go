package helper

import (
	"beet_pos/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GeneratePagination(ctx *gin.Context) *dto.Pagination{
	limit := 10
	page := 0
	sort := ""

	query := ctx.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	// check if query param key contains dot

	return &dto.Pagination{
		Limit:        limit,
		Page:         page,
		Sort:         sort,
	}
}