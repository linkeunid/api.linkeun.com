package helpers

import (
	"strings"

	"github.com/goravel/framework/contracts/http"
)

type PaginationInput struct {
	Page    int
	PerPage int
	Sort    string // "name desc" or "created_at asc"
	Search  string
	Filters map[string]string
}

func ParsePaginationRequest(ctx http.Context) PaginationInput {
	request := ctx.Request()

	page := request.QueryInt("page", 1)
	perPage := request.QueryInt("per_page", 10)

	// Ensure sensible pagination limits
	if page < 1 {
		page = 1
	}
	if perPage < 1 {
		perPage = 10
	}
	if perPage > 100 {
		perPage = 100
	}

	// Support sortBy + sort or fallback to sort only
	sortBy := request.Query("sortBy", "")
	sortDir := request.Query("sort", "")
	sort := ""
	if sortBy != "" {
		if sortDir != "" && (strings.ToLower(sortDir) == "asc" || strings.ToLower(sortDir) == "desc") {
			sort = sortBy + " " + strings.ToLower(sortDir)
		} else {
			sort = sortBy + " asc" // default to ascending
		}
	} else if sortDir != "" {
		sort = sortDir // fallback: "name" or "name desc"
	}

	search := strings.TrimSpace(request.Query("search", ""))

	// Extract filters
	filters := map[string]string{}
	// Get all query parameters
	for key := range request.All() {
		if key != "page" && key != "per_page" && key != "sort" && key != "sortBy" && key != "search" {
			value := request.Query(key, "")
			if value != "" {
				filters[key] = value
			}
		}
	}

	return PaginationInput{
		Page:    page,
		PerPage: perPage,
		Sort:    sort,
		Search:  search,
		Filters: filters,
	}
}
