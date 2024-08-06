package ProductId

import (
	"github.com/google/uuid"
	"github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Errors"
)

type ProductId struct{
    Id        uuid.UUID
}

 
 
func Get(value string) (ProductId,error){
	if( value ==  ""){
            return ProductId{}, Errors.ErrorProductId
	}
	return ProductId{
		  Id:  uuid.MustParse(value),
	  } ,nil
}


