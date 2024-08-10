package queries
 

type ByIdStockQuery struct {
	StockId     string
}

func NewByIdStockQuery(stockId string) ByIdStockQuery{
       return ByIdStockQuery{
		   StockId: stockId,
	   }
}

