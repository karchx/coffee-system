package products

import (
	"context"
	"github/karchx/coffee-system/internal/product/domain"

	"github.com/pkg/errors"
)


var _ UseCase = (*service)(nil)


type service struct {
  repo domain.ProductRepo
}

func (s *service) GetItemTypes(ctx context.Context) ([]*domain.ItemTypeDto, error) {
  results, err := s.repo.GetAll(ctx)
  if err != nil {
    return nil, errors.Wrap(err, "service.GetItemTypes")
  }

  return results,nil
}

func (s *service) GetItemsBYTypes() {}
