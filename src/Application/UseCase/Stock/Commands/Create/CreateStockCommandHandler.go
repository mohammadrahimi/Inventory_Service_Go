package useCase

import (
	"fmt"

	commands "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Create"
	aggregate "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock"
	repository "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Repository"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
	 
)
 
type CreateStockCommandHandler struct {
	  repository  repository.IStockRepository
}

func NewCreateStockCommandHandler(repository  repository.IStockRepository) *CreateStockCommandHandler{
	return &CreateStockCommandHandler{
		repository: repository,
	}
}
 
func (h *CreateStockCommandHandler) Handle(command  cqrs.Command) error {
  
	switch c := command.(type) {

			case *commands.CreateStockCommand:{
				 
				  stock,err:= aggregate.NewStock(c)
				  if(err != nil){
					  return err
				  }
				 
				status,err:= h.repository.Create(stock)
				if(err != nil){
					return err
				}

				fmt.Println("  status = " +   status  )    //s := strconv.FormatFloat(c.Price.Amount, 'f', -1, 64)

			}
    }
	
	return nil
}