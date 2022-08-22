package helper

import (
	"beet_pos/dto"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


func GeneratePagination(ctx *gin.Context) *dto.Pagination{
	limit := 10
	page := 0
	sort := ""

	var searchs []dto.Search

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

		// check if query param key contains dot
		if strings.Contains(key, ".") {
			// split query parameter key by dot
			searchKeys := strings.Split(key, ".")

			// create search object
			search := dto.Search{
				Column: searchKeys[0],
				Action: searchKeys[1],
				Query: queryValue,
			}

			// add search object to searchs array
			searchs = append(searchs, search)

		}
	}
	

	return &dto.Pagination{
		Limit:        limit,
		Page:         page,
		Sort:         sort,
		Searchs: searchs,
	}
}