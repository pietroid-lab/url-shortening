package internal

import "context"

type RepositoryInterface interface {
}

type Service struct {
	repository RepositoryInterface
}

// NewService constructs a Service from any type that satisfies
// RepositoryInterface. Accepting the interface decouples Service from a
// concrete repository implementation and makes wiring simpler.
func NewService(repository RepositoryInterface) *Service {
	return &Service{repository: repository}
}

func (s *Service) GenerateShorterURL(ctx context.Context, originaURL, slug  string ) (string, error) {


	return "", nil
}
