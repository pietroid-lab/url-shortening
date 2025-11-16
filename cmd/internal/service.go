package internal

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
)

const (
	DefaultDomain = "http://localhost:8080/"
)

type RepositoryInterface interface {
	SaveURL(ctx context.Context, originalURL, hash string)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) *Service {
	return &Service{repository: repository}
}

func (s *Service) GenerateShorterURL(ctx context.Context, originaURL, slug string) (string, error) {

	if slug == "" {
		slug = s.generateShortURL(originaURL)
		s.repository.SaveURL(ctx, originaURL, slug)
		return DefaultDomain + slug, nil
	}

	s.repository.SaveURL(ctx, originaURL, slug)

	return "", nil
}

func (s *Service) generateShortURL(originalURL string) string {
	hasher := sha256.New()
	hasher.Write([]byte(originalURL))
	hash := hasher.Sum(nil)
	encoded := base64.URLEncoding.EncodeToString(hash)
	return encoded[:8]
}
