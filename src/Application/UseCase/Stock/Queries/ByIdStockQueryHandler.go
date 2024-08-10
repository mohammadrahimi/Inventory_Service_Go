package useCase

import (
	"reflect"

	queries "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ByIdStock"
	repository "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Repository"
	cqrs "github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
)

type ByIdStockQueryHandler struct {
	repository repository.IStockRepository
}

func NewByIdStockQueryHandler(repository repository.IStockRepository) *ByIdStockQueryHandler {
	return &ByIdStockQueryHandler{
		repository: repository,
	}
}

func (h *ByIdStockQueryHandler) Handle(query cqrs.Query) (cqrs.QueryResult,error) {

			obj := reflect.ValueOf(query)
			byIdStockQuery := obj.Interface().(queries.ByIdStockQuery)
			
			queryResult,err:= h.repository.FindById(byIdStockQuery.StockId) 
			if(err != nil){
				return   nil,err
			}
			return queryResult,nil
			
}