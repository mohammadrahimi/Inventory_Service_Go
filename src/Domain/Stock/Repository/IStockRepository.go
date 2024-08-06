package repository

import (
	"context"

	aggregate "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock"
)

type IStockRepository interface {
	Create(ctx context.Context, stock *aggregate.Stock) ( status string, err error)
	FindById(ctx context.Context, id string) (*aggregate.Stock, error)
	FindAll(ctx context.Context) ([]*aggregate.Stock, error)
	Update(ctx context.Context, stock *aggregate.Stock) ( status string, err error)
	Delete(ctx context.Context, id string) ( status string, err error)
}