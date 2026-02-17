package repository

import (
	"context"

	"github.com/Pyth0nHater/link-shorter/internal/models"
)

type MemoryLinkRepository struct {
	mapa *models.MemoryMap
}

func NewMemeoryLinkRepository(mapa *models.MemoryMap) LinkRepositoryInterface {
	return &MemoryLinkRepository{mapa: mapa}
}

func (r *MemoryLinkRepository) CreateLink(ctx context.Context, createLink *models.CreateLink) error {
	r.mapa.Mutex.Lock()
	defer r.mapa.Mutex.Unlock()

	r.mapa.Storage[createLink.MainLink] = createLink.ShortLink
	r.mapa.ReverseStorage[createLink.ShortLink] = createLink.MainLink
	return nil
}


func (r *MemoryLinkRepository) GetLink(ctx context.Context, getLink models.GetLinkFilter) (*models.GetLink, error) {
	r.mapa.Mutex.RLock()
	defer r.mapa.Mutex.RUnlock()

	if getLink.MainLink != nil && getLink.ShortLink != nil{
		return nil, ErrInvalidGetFilter
	}

	if getLink.MainLink != nil{
		res, ok:=r.mapa.Storage[*getLink.MainLink]

		if ok{
			return &models.GetLink{ShortLink: res, MainLink: *getLink.MainLink}, nil
		}
	
		return nil, ErrNotFound
	}

	if getLink.ShortLink != nil{
		res, ok:=r.mapa.ReverseStorage[*getLink.ShortLink]

		if ok{
			return &models.GetLink{MainLink: res, ShortLink: *getLink.ShortLink}, nil
		}

		return nil, ErrNotFound

	}

	return nil, ErrInvalidGetFilter
}

