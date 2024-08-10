package main

import (
	"encoding/json"
	"log"
	"net/http"

	useCaseCreateStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Application/UseCase/Stock/Commands/Create"
	useCaseDeleteStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Application/UseCase/Stock/Commands/Delete"
	commandCreateStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Create"
	commandDeleteStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Commands/Stock/Delete"
	 
	dbconnection "github.com/mohammadrahimi/Inventory_Service_Go/src/Infrastructure/Persistence.Sql/DbConnection"
	stock "github.com/mohammadrahimi/Inventory_Service_Go/src/Infrastructure/Persistence.Sql/Repository/Stock"

	cqrs "github.com/mohammadrahimi/Inventory_Service_Go/src/Framework.Core/Bus"

	useCaseQuery "github.com/mohammadrahimi/Inventory_Service_Go/src/Application/UseCase/Stock/Queries"
	queryByIdProduct "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ByIdProduct"
	queryByIdStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ByIdStock"
	querListStock "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ListStock"
)

var (
	commandBus cqrs.CommandBus
	queryBus   cqrs.QueryBus
)

func init() {

	//dsn := "sqlserver://userID:password@host?database=Database_name"
	DbConnection := dbconnection.NewSQLConnection("sqlserver://@192.168.56.1?database=Inventory")
	DB, err := DbConnection.DBSQL()
	if err != nil {
		panic(" ConnectionSql is Error !  " + err.Error())
	}

	repo := stock.NewStockRepository(DB)

	CreateStockCommandHandler := useCaseCreateStock.NewCreateStockCommandHandler(repo)
	DeleteStockCommandHandler := useCaseDeleteStock.NewDeleteStockCommandHandler(repo)

	commandBus = *cqrs.NewCommandBus()
	commandBus.RegisterHandler(CreateStockCommandHandler, &commandCreateStock.CreateStockCommand{})
	commandBus.RegisterHandler(DeleteStockCommandHandler, &commandDeleteStock.DeleteStockCommand{})

	ByIdStockQueryHandler := useCaseQuery.NewByIdStockQueryHandler(repo)
	ByIdProductQueryHandler := useCaseQuery.NewByIdProductQueryHandler(repo)
	ListStockQueryHandler := useCaseQuery.NewListStockQueryHandler(repo)

	queryBus = *cqrs.NewQueryBus()
	queryBus.RegisterHandler(ByIdStockQueryHandler, &queryByIdStock.ByIdStockQuery{})
	queryBus.RegisterHandler(ByIdProductQueryHandler, &queryByIdProduct.ByIdProductQuery{})
	queryBus.RegisterHandler(ListStockQueryHandler, &querListStock.ListStockQuery{})

}

// @title  Inventory API
// @version 1.0
// @description This is a s Inventory API in My Microservice
// @termsOfService http://swagger.io/terms/
// @host localhost:8088
// @BasePath /

func main() {

	//app := fiber.New()
	//app.Get("/swagger/*", swagger.HandlerDefault)

	mux := ControllerApi()
	if err := http.ListenAndServe(":8088", mux); err != nil {
		log.Fatal(err)
	}
}

func  ControllerApi() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
 
			command := &commandCreateStock.CreateStockCommand{}

			err := json.NewDecoder(r.Body).Decode(&command)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			errbus := commandBus.Send(command)
			if errbus != nil {
				http.Error(w, errbus.Error(), http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("{State:OkCreate,Status:200}"))
		}

	})

	mux.HandleFunc("/del/{id}", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodDelete {

			id := r.PathValue("id")
			command := &commandDeleteStock.DeleteStockCommand{StockId: id}
			err := json.NewDecoder(r.Body).Decode(&command)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			errbus := commandBus.Send(command)
			if errbus != nil {
				http.Error(w, errbus.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{State:OkDelete,Status:200}"))
		}

	})

	mux.HandleFunc("/ByProductId/{id}", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPut {
 
			id := r.PathValue("id")
			query := queryByIdProduct.NewByIdProductQuery(id)

			result, err := queryBus.Send(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			data, err := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}

	})

	mux.HandleFunc("/ByStockId/{id}", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPut{

			id := r.PathValue("id")
			query := queryByIdStock.NewByIdStockQuery(id)

			result, err := queryBus.Send(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			data, err := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}

	})

	mux.HandleFunc("/ByAll", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			query := querListStock.NewListStockQuery()

			result, err := queryBus.Send(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				//http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data, err := json.Marshal(result)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			 
		}

	})

	return mux
}
