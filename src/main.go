package main

import (
	"log"
	"net/http"

      useCaseCreateStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Application/UseCase/Stock/Commands/Create"
	  useCaseDeleteStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Application/UseCase/Stock/Commands/Delete"
	  commandCreateStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Create"
	  commandDeleteStock  "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Delete"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Dto/Stock/MoneyDto"

	"github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"
)


var (
	 
	commandBus cqrs.CommandBus
)

func init() {

	 
	CreateStockCommandHandler := useCaseCreateStock.NewCreateStockCommandHandler()
	DeleteStockCommandHandler := useCaseDeleteStock.NewDeleteStockCommandHandler()

	commandBus = *cqrs.NewCommandBus()
	commandBus.RegisterHandler(CreateStockCommandHandler, &commandCreateStock.CreateStockCommand{})
	commandBus.RegisterHandler(DeleteStockCommandHandler, &commandDeleteStock.DeleteStockCommand{})
	 
}


func main() {
	mux := setupHandlers()
	if err := http.ListenAndServe(":8088", mux); err != nil {
		log.Fatal(err)
	}
}

func setupHandlers() *http.ServeMux {

    mux := http.NewServeMux()

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		 
		if r.Method == http.MethodGet {
            
			command := &commandCreateStock.CreateStockCommand{
				Price:  MoneyDto.MoneyDto{Amount: 1000,Currency: "rial"},
				Quantity: 2,
				ProductId: "ff-dd-cc-mm",
		    }

			err := commandBus.Send(command)
			if err != nil {
				log.Println(err)
			}

			w.Write([]byte(" Create Stock Command " ))
		}
		 
	})

	mux.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
		 
		if r.Method == http.MethodGet {
            
			command := &commandDeleteStock.DeleteStockCommand{
				 StockId : "ff-dd-cc-mm",
		    }

			err := commandBus.Send(command)
			if err != nil {
				log.Println(err)
			}

			w.Write([]byte(" Delete Stock Command " ))
		}
		 
	})

	return mux
}