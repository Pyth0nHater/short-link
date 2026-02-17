package service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	"github.com/Pyth0nHater/link-shorter/internal/models"
	"github.com/Pyth0nHater/link-shorter/internal/repository"
)

type LinkService struct {
	linkRepo repository.LinkRepositoryInterface
}

func NewLinkService(linkRepo repository.LinkRepositoryInterface) *LinkService {
	return &LinkService{linkRepo: linkRepo}
}

func (s *LinkService) CreateLink(ctx context.Context, mainLink *models.MainLink) (string, error) {

	h := sha1.New()
	h.Write([]byte(mainLink.MainLink))
	ShortLink := hex.EncodeToString(h.Sum(nil))[:10]

	createLink := &models.CreateLink{
		MainLink: mainLink.MainLink,
		ShortLink: ShortLink,
	}

	err := s.linkRepo.CreateLink(ctx, createLink)
	if err != nil {
		return "", err
	}

	return ShortLink, nil
}

func (s *LinkService) GetLink(ctx context.Context, getLinkFilter models.GetLinkFilter) (*models.GetLink, error) {
	return s.linkRepo.GetLink(ctx, getLinkFilter)
}

