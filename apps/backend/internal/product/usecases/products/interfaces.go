package products

import (
	"context"
	"github/karchx/coffee-system/internal/product/domain"
)

type UseCase interface {
  GetItemTypes(context.Context)([]*domain.ItemTypeDto, error)
  GetItemsByType(context.Context, string)([]*domain.ItemDto, error)
}
