package useCase

import (
	"reflect"

    queries "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ByIdProduct"
	repository "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Repository"
	cqrs "github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
)

type ByIdProductQueryHandler struct {
	repository repository.IStockRepository
}

func NewByIdProductQueryHandler(repository repository.IStockRepository) *ByIdProductQueryHandler {
	return &ByIdProductQueryHandler{
		repository: repository,
	}
}

func (h *ByIdProductQueryHandler) Handle(query cqrs.Query) (cqrs.QueryResult,error)  {  

	 
			obj := reflect.ValueOf(query)
			byIdProductQuery := obj.Interface().(queries.ByIdProductQuery)
	 
			queryResult,err:= h.repository.FindByProductId(byIdProductQuery.ProductId) 
			if(err != nil){
				return   nil,err
			}
			return queryResult,nil
	 
}