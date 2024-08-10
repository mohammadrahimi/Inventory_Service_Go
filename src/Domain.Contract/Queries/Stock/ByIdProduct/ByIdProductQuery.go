package queries
 

type ByIdProductQuery struct {
	ProductId     string
}

func NewByIdProductQuery(productId string) ByIdProductQuery{
       return ByIdProductQuery{
		  ProductId: productId,
	   }
}
