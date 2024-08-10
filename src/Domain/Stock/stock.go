package aggregate

import (
	"time"

	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Create"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Errors"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/Money"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/ProductId"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/StockId"
)


type Stock struct{
    Id          StockId.StockId
    CreatedAt   time.Time
    Price       Money.Money
    Quantity    int
    ProductId   ProductId.ProductId
}



func NewStock(createStockCommand *commands.CreateStockCommand) (*Stock,error){

     if(createStockCommand.Quantity <= 0){
        return &Stock{}, Errors.ErrorQuantity
     }
      
     price,err := Money.New(createStockCommand.Price.Amount,createStockCommand.Price.Currency)
     if(err != nil){
         return &Stock{},   err
     }

     productId,err := ProductId.Get(createStockCommand.ProductId)
     if(err != nil){
        return &Stock{},   err
    }

     return &Stock{
         Id:  StockId.New(),
         CreatedAt: time.Now(),
         Price: price,
         Quantity: createStockCommand.Quantity,
         ProductId: productId,
     }, nil

}

func (s *Stock) GetPrice() Money.Money{
     return s.Price
}

func (s *Stock) GetQuantity() int{
    return s.Quantity
}

func (s *Stock) GetProductId() ProductId.ProductId{
    return s.ProductId
}
