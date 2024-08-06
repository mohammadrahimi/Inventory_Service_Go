package Errors

import "errors"
 
 
var (
	ErrorQuantity = errors.New("stock quantity is not correct")
	ErrorMoney = errors.New(" Money is not correct")
	ErrorProductId = errors.New(" ProductId is Empty")
	ErrorStockId = errors.New(" StockId is Empty")
)
