package useCase

import (
	"fmt"
 

	commands "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Delete"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
)
 
type DeleteStockCommandHandler struct {
	// repository add
}

func NewDeleteStockCommandHandler() *DeleteStockCommandHandler {
	return &DeleteStockCommandHandler{
	}
}
 
func (h *DeleteStockCommandHandler) Handle(command  cqrs.Command) error {
  
	switch c := command.(type) {

			case *commands.DeleteStockCommand:{
				 
				fmt.Println(" quantity = " + c.StockId   )  

			}
    }
	
	return nil
}