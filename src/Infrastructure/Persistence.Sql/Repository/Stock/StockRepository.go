package stock

import (
	"errors"

	queries "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain.Contract/Queries/Stock/ResultStock"
	aggregate "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock"
	repository "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Repository"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/ProductId"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/ValueObject/StockId"
	models "github.com/mohammadrahimi/Inventory_Service_Go/src/Infrastructure/Persistence.Sql/Models"

	"gorm.io/gorm"
)

var(
	ErrorFailCreate = errors.New("Stock Not Create!")
	ErrorFailDelete = errors.New("Stock Not Delete!")
)
 
type StockRepository struct {
	db    *gorm.DB
}

func NewStockRepository(db *gorm.DB)  repository.IStockRepository{
	return &StockRepository{db: db}
}
 
func (s *StockRepository) Create(stock *aggregate.Stock) (status string, err error) {
	 
	dbStock := fromDBStock(stock)

	res:= s.db.Table("Stock").Create(dbStock) 
    if(res.RowsAffected > 0) {
		return  "Status.200",nil
	}
	return "Status.400", ErrorFailCreate
}

 
func (s *StockRepository) Delete(id string) (status string, err error) {
	stockId,_ := StockId.Get(id)
	var  stock = models.StockEntity{Id: stockId.Id}
	res:= s.db.Table("Stock").Delete(stock)
	if(res.RowsAffected > 0) {
		return  "Status.200",nil
	}
	return "Status.400", ErrorFailDelete
}
 
func (s *StockRepository) Update(stock *aggregate.Stock) (status string, err error) {
	panic("unimplemented")
}


func (s *StockRepository) FindAll() ([]queries.ResultStockQuery, error) {

		var dbStocks  []models.StockEntity
		s.db.Table("Stock").Find(&dbStocks)  
		stocks := make([]queries.ResultStockQuery, len(dbStocks))
		for i,stock := range dbStocks {
			stocks[i] = fromDbStockQuery(&stock)
		}
		return stocks,nil
}

 
func (s *StockRepository) FindById(id string) (queries.ResultStockQuery, error) {
	 
	stockId,_ := StockId.Get(id)
	var  stock   models.StockEntity 
	s.db.Table("Stock").Select("*").Where("id=?", stockId.Id).Find(&stock)
    return fromDbStockQuery(&stock),nil  
	 
}

func (s *StockRepository) FindByProductId(id string) (queries.ResultStockQuery, error) {
	 
	productId,_ := ProductId.Get(id)
	var  stock   models.StockEntity 
	s.db.Table("Stock").Select("*").Where("product_id=?", productId.Id).Find(&stock)
    return fromDbStockQuery(&stock),nil  
	 
}
