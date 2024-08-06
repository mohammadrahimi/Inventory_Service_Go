package commands

import (
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Dto/Stock/MoneyDto"
	 
)



type CreateStockCommand struct {
	Price         MoneyDto.MoneyDto
	Quantity      int
	ProductId     string
}

func NewCreateStockCommand(price MoneyDto.MoneyDto,quantity int, productId string) CreateStockCommand{
       return CreateStockCommand{
		  Price: price,
		  Quantity: quantity,
		  ProductId: productId,
	   }
}
