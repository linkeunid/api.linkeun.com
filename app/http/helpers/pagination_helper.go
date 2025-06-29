package helpers

import "math"

func MakePaginationMeta(total int64, page, perPage int) map[string]any {
	lastPage := int(math.Ceil(float64(total) / float64(perPage)))

	return map[string]any{
		"total":     total,
		"page":      page,
		"per_page":  perPage,
		"last_page": lastPage,
		"has_next":  page < lastPage,
		"has_prev":  page > 1,
	}
}
