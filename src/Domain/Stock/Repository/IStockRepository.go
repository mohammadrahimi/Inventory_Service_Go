package repository

import (
	queries "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ResultStock"
	aggregate "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock"
)

type IStockRepository interface {
	Create(stock *aggregate.Stock) ( status string, err error)
	FindByProductId(id string) (queries.ResultStockQuery, error)
	FindById(id string) (queries.ResultStockQuery, error)
	FindAll() ([]queries.ResultStockQuery, error)
	Update(stock *aggregate.Stock) ( status string, err error)
	Delete(id string) ( status string, err error)
}