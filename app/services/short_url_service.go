package services

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/linkeunid/api.linkeun.com/app/models"
	"gorm.io/gorm"
)

type ShortUrlService struct {
	userService *UserService
}

func NewShortUrlService() *ShortUrlService {
	return &ShortUrlService{
		userService: NewUserService(),
	}
}

func (s *ShortUrlService) CreateShortUrl(ctx context.Context, userId *uint64, originalUrl string, customAlias *string, password *string, description *string) (*models.Url, error) {
	var shortCode string
	var passwordHash *string

	if customAlias != nil && *customAlias != "" {
		shortCode = *customAlias
	} else {
		var err error
		shortCode, err = s.generateShortCode()
		if err != nil {
			return nil, err
		}

		for s.shortCodeExists(shortCode) {
			shortCode, err = s.generateShortCode()
			if err != nil {
				return nil, err
			}
		}
	}

	if password != nil && *password != "" {
		hash, err := facades.Hash().Make(*password)
		if err != nil {
			return nil, err
		}
		passwordHash = &hash
	}

	url := &models.Url{
		UserId:       userId,
		ShortCode:    shortCode,
		OriginalUrl:  originalUrl,
		IsActive:     true,
		CustomAlias:  customAlias,
		PasswordHash: passwordHash,
		Description:  description,
		ClicksCount:  0,
	}

	if err := facades.Orm().Query().Create(url); err != nil {
		return nil, err
	}

	return url, nil
}

func (s *ShortUrlService) generateShortCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func (s *ShortUrlService) shortCodeExists(shortCode string) bool {
	var count int64
	facades.Orm().Query().Model(&models.Url{}).Where("short_code = ? OR custom_alias = ?", shortCode, shortCode).Count(&count)
	return count > 0
}

func (s *ShortUrlService) GetUrlByShortCode(shortCode string) (*models.Url, error) {
	var url models.Url
	err := facades.Orm().Query().Where("(short_code = ? OR custom_alias = ?) AND is_active = ?", shortCode, shortCode, true).First(&url)
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (s *ShortUrlService) GetUrlsByUserId(userId uint64) ([]*models.Url, error) {
	var urls []*models.Url
	err := facades.Orm().Query().Where("user_id = ?", userId).Find(&urls)
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (s *ShortUrlService) ProcessClick(urlId uint, ipAddress, userAgent, referrer string) {
	err := facades.Orm().Transaction(func(tx orm.Query) error {
		result, err := tx.Model(&models.Url{}).Where("id = ?", urlId).Update("clicks_count", gorm.Expr("clicks_count + 1"))
		if err != nil {
			return err
		}

		if result.RowsAffected == 0 {
			return errors.New("url not found or not updated")
		}

		click := models.Click{
			UrlID:     uint64(urlId),
			ClickedAt: time.Now(),
		}

		if ipAddress != "" {
			click.IpAddress = &ipAddress
		}
		if userAgent != "" {
			click.UserAgent = &userAgent
		}
		if referrer != "" {
			click.Referrer = &referrer
		}

		if err := tx.Create(&click); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		facades.Log().Errorf("failed to process click for url id %d: %v", urlId, err)
	}
}
