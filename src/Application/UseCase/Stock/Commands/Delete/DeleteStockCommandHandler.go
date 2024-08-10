package useCase

import (
	"fmt"

	commands "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Delete"
	repository "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Repository"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
)
 
type DeleteStockCommandHandler struct {
	repository  repository.IStockRepository
}

func NewDeleteStockCommandHandler(repository  repository.IStockRepository) *DeleteStockCommandHandler {
	return &DeleteStockCommandHandler{
		repository: repository,
	}
}
 
func (h *DeleteStockCommandHandler) Handle(command  cqrs.Command) error {
  
	switch c := command.(type) {

			case *commands.DeleteStockCommand:{
				 
				status,err:= h.repository.Delete(c.StockId)
				if(err != nil){
					return err
				}

				fmt.Println("  status = " +   status  ) 

			}
    }
	
	return nil
}