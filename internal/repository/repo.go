package repository

import (
	"context"
	"errors"

	"github.com/Pyth0nHater/link-shorter/internal/models"
)


var (
	ErrLinkAlreadyExists = errors.New("Ссылка уже существует")
	ErrInvalidGetFilter = errors.New("Невалидная структура фильтра")
	ErrNotFound = errors.New("Ссылка не найдена")
)

type LinkRepositoryInterface interface {
	CreateLink(ctx context.Context, link *models.CreateLink) (error)
	GetLink(ctx context.Context, getLink models.GetLinkFilter) (*models.GetLink, error)
}