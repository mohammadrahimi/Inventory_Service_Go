package stock

import (
	"log"

	queries "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ResultStock"
	aggregate "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/Money"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/ProductId"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/StockId"
	models "github.com/mohammadrahimi/Inventory_Service_Go/src/Infrastructure/Persistence.Sql/Models"
)

 

func toDBStock(stock *models.StockEntity) *aggregate.Stock {

	stockId,err := StockId.Get(stock.Id.String())
	if(err != nil){
		log.Fatal("stockId is false")
    }
	price,err := Money.New(stock.Amount,stock.Currency)
	if(err != nil){
	   log.Fatal("money is false")
	}
	productId,err := ProductId.Get(stock.ProductId.String())
	if(err != nil){
		log.Fatal("productId is false")
    }

	var s = &aggregate.Stock{
		 Id: stockId,
		 CreatedAt: stock.CreatedAt,
		 Price: price,
		 Quantity: stock.Quantity,
		 ProductId: productId,
	}

	return s
}

func fromDBStock(stock *aggregate.Stock) *models.StockEntity {

	var s = &models.StockEntity{
		Id: stock.Id.Id,
	    Amount: stock.Price.Amount(),
		Currency: stock.Price.Currency(),
		Quantity: stock.Quantity,
		CreatedAt: stock.CreatedAt,
		ProductId: stock.ProductId.Id,
	}
	 
	return s
}

func fromDbStockQuery(stock *models.StockEntity)  queries.ResultStockQuery{
	var q =  queries.ResultStockQuery{
		 Amount: stock.Amount,
		 Currency: stock.Currency,
		 Quantity: stock.Quantity,
		 ProductId: stock.ProductId.String(),
		 ProductName: "",
	}
	 
	return q
}