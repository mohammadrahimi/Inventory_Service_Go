package stock

import (
	"context"
	aggregate "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock"
	repository "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Repository"
)

type StockRepository struct {
	connectionstring string
}

func NewStockRepository(connectionstring string) repository.IStockRepository {
	return &StockRepository{connectionstring: connectionstring}
}


// Create implements repository.IStockRepository.
func (s *StockRepository) Create(ctx context.Context, stock *aggregate.Stock) (status string, err error) {
	panic("unimplemented")
}

// Delete implements repository.IStockRepository.
func (s *StockRepository) Delete(ctx context.Context, id string) (status string, err error) {
	panic("unimplemented")
}

// FindAll implements repository.IStockRepository.
func (s *StockRepository) FindAll(ctx context.Context) ([]*aggregate.Stock, error) {
	panic("unimplemented")
}

// FindById implements repository.IStockRepository.
func (s *StockRepository) FindById(ctx context.Context, id string) (*aggregate.Stock, error) {
	panic("unimplemented")
}

// Update implements repository.IStockRepository.
func (s *StockRepository) Update(ctx context.Context, stock *aggregate.Stock) (status string, err error) {
	panic("unimplemented")
}


