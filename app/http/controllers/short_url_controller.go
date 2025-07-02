package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/http/helpers"
	"github.com/linkeunid/api.linkeun.com/app/http/requests"
	"github.com/linkeunid/api.linkeun.com/app/http/resources"
	"github.com/linkeunid/api.linkeun.com/app/models"
	"github.com/linkeunid/api.linkeun.com/app/services"
)

type ShortUrlController struct {
	shortUrlService *services.ShortUrlService
}

func NewShortUrlController() *ShortUrlController {
	return &ShortUrlController{
		shortUrlService: services.NewShortUrlService(),
	}
}

func (r *ShortUrlController) Index(ctx http.Context) http.Response {
	var user models.User
	if err := facades.Auth(ctx).User(&user); err != nil {
		return helpers.JsonResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, err)
	}

	urls, err := r.shortUrlService.GetUrlsByUserId(uint64(user.ID))
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	}

	data := resources.MakeUrlCollection(urls)
	return helpers.JsonResponse(ctx, http.StatusOK, "success", data, nil)
}

func (r *ShortUrlController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *ShortUrlController) Store(ctx http.Context) http.Response {
	var request requests.CreateShortUrlRequest
	errors, _ := ctx.Request().ValidateRequest(&request)
	if errors != nil {
		return helpers.JsonResponse(ctx, http.StatusUnprocessableEntity, "error", nil, errors.All())
	}

	var userId *uint64
	var user models.User
	if err := facades.Auth(ctx).User(&user); err == nil {
		id := uint64(user.ID)
		userId = &id
	}

	url, err := r.shortUrlService.CreateShortUrl(
		ctx.Context(),
		userId,
		request.OriginalUrl,
		request.CustomAlias,
		request.Password,
		request.Description,
	)
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusInternalServerError, "error", nil, err)
	}

	data := resources.MakeUrlResource(url).ToArray()
	return helpers.JsonResponse(ctx, http.StatusCreated, "success", data, nil)
}

func (r *ShortUrlController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *ShortUrlController) Destroy(ctx http.Context) http.Response {
	return nil
}

func (r *ShortUrlController) Redirect(ctx http.Context) http.Response {
	shortCode := ctx.Request().Route("shortCode")

	url, err := r.shortUrlService.GetUrlByShortCode(shortCode)
	if err != nil {
		return helpers.JsonResponse(ctx, http.StatusNotFound, "URL not found", nil, nil)
	}

	if url.PasswordHash != nil {
		password := ctx.Request().Input("password")
		if password == "" {
			return helpers.JsonResponse(ctx, http.StatusUnauthorized, "Password required", nil, nil)
		}

		if !facades.Hash().Check(password, *url.PasswordHash) {
			return helpers.JsonResponse(ctx, http.StatusUnauthorized, "Invalid password", nil, nil)
		}
	}

	go func() {
		ipAddress := ctx.Request().Ip()
		userAgent := ctx.Request().Header("User-Agent")
		referrer := ctx.Request().Header("Referer")
		r.shortUrlService.ProcessClick(url.ID, ipAddress, userAgent, referrer)
	}()

	return ctx.Response().Redirect(http.StatusFound, url.OriginalUrl)
}
