package queries

 

type ResultStockQuery struct {
	Amount       float64
	Currency     string  
	Quantity     int
	ProductId    string
	ProductName  string
}

 
func NewResultStockQuery(amount float64,currency string,quantity int, productId string,productName string) ResultStockQuery{
	return ResultStockQuery{
	   Amount: amount,
	   Currency: currency,
	   Quantity: quantity,
	   ProductId: productId,
	   ProductName: productName,
	}
}
