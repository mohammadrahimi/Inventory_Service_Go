package useCase

import (
	 
	cqrs "github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
	repository "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Repository"
)

type ListStockQueryHandler struct {
	repository repository.IStockRepository
}

func NewListStockQueryHandler(repository repository.IStockRepository) *ListStockQueryHandler {
	return &ListStockQueryHandler{
		repository: repository,
	}
}

func (h *ListStockQueryHandler) Handle(query cqrs.Query)  (cqrs.QueryResult,error)  {  
  
			 
			  queryResult,err:= h.repository.FindAll() 
			  if(err != nil){
				  return   nil,err
			  }
			  return queryResult,nil
			 
}