package commands

 
type DeleteStockCommand struct {
	StockId     string
}

func NewDeleteStockCommand(stockId string) DeleteStockCommand{
       return DeleteStockCommand{
		 StockId: stockId,
	   }
}
