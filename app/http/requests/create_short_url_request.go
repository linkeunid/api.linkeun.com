package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateShortUrlRequest struct {
	OriginalUrl  string  `form:"original_url" json:"original_url"`
	CustomAlias  *string `form:"custom_alias" json:"custom_alias"`
	Password     *string `form:"password" json:"password"`
	Description  *string `form:"description" json:"description"`
}

func (r *CreateShortUrlRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateShortUrlRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateShortUrlRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"original_url": "required|full_url",
		"custom_alias": "string|max_len:50|alpha_dash",
		"password":     "string|min_len:6",
		"description":  "string|max_len:255",
	}
}

func (r *CreateShortUrlRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateShortUrlRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateShortUrlRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
