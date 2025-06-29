package resources

import "github.com/linkeunid/api.linkeun.com/app/models"

type UserResource struct {
	User *models.User
}

func MakeUserResource(user *models.User) *UserResource {
	return &UserResource{User: user}
}

func MakeUserCollection(users []*models.User) []map[string]interface{} {
	var result []map[string]interface{}
	for _, u := range users {
		result = append(result, MakeUserResource(u).ToArray())
	}
	return result
}

func (r *UserResource) ToArray() map[string]interface{} {
	return map[string]interface{}{
		"id":         r.User.ID,
		"name":       r.User.Name,
		"email":      r.User.Email,
		"created_at": r.User.CreatedAt,
		"updated_at": r.User.UpdatedAt,
	}
}
