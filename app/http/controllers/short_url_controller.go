package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type ShortUrlController struct {
	// Dependent services
}

func NewShortUrlController() *ShortUrlController {
	return &ShortUrlController{
		// Inject services
	}
}

func (r *ShortUrlController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *ShortUrlController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *ShortUrlController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *ShortUrlController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *ShortUrlController) Destroy(ctx http.Context) http.Response {
	return nil
}
