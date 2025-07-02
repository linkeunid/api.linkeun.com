package resources

import "github.com/linkeunid/api.linkeun.com/app/models"

type UrlResource struct {
	Url *models.Url
}

func MakeUrlResource(url *models.Url) *UrlResource {
	return &UrlResource{Url: url}
}

func MakeUrlCollection(urls []*models.Url) []map[string]interface{} {
	var result []map[string]interface{}
	for _, u := range urls {
		result = append(result, MakeUrlResource(u).ToArray())
	}
	return result
}

func (r *UrlResource) ToArray() map[string]interface{} {
	return map[string]interface{}{
		"id":           r.Url.ID,
		"user_id":      r.Url.UserId,
		"short_code":   r.Url.ShortCode,
		"original_url": r.Url.OriginalUrl,
		"is_active":    r.Url.IsActive,
		"custom_alias": r.Url.CustomAlias,
		"description":  r.Url.Description,
		"clicks_count": r.Url.ClicksCount,
		"created_at":   r.Url.CreatedAt,
		"updated_at":   r.Url.UpdatedAt,
		"deleted_at":   r.Url.DeletedAt,
	}
}
