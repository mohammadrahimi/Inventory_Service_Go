package StockId

import (

	"github.com/google/uuid"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Errors"
)

type StockId struct{
    Id        uuid.UUID
}

 

func  New() StockId{
      return StockId{
		    Id: uuid.New(),
		} 
}
 
func Get(value string) (StockId,error){
	if( value ==  ""){
            return StockId{}, Errors.ErrorStockId
	}
	return StockId{
		  Id:  uuid.MustParse(value),
	  } ,nil
}

