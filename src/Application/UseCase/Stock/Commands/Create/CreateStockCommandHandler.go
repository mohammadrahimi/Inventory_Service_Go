package useCase

import (
	"fmt"
	"strconv"

	commands "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Create"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
)
 
type CreateStockCommandHandler struct {
	// repository add
}

func NewCreateStockCommandHandler() *CreateStockCommandHandler {
	return &CreateStockCommandHandler{
	}
}
 
func (h *CreateStockCommandHandler) Handle(command  cqrs.Command) error {
  
	switch c := command.(type) {

			case *commands.CreateStockCommand:{
				 
				s := strconv.FormatFloat(c.Price.Amount, 'f', -1, 64)
				fmt.Println(" quantity = " + string(c.Quantity)  + "  price = " + string(s) + " : "+ c.Price.Currency  )  

			}
    }
	
	return nil
}