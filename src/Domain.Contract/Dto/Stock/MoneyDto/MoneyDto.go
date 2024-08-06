
package MoneyDto
 

type MoneyDto struct {
	 
	Amount     float64
	Currency   string      

}

func NewMoneyDto(amount float64,currency string) MoneyDto{
      return MoneyDto{
		Amount: amount,
		Currency: currency,
	  }
}
