package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExamplePaymentSources() {
	cards := []types.Card {
		{Balance: 10_000, Active: true, PAN: "5058 xxxx xxxx 9999",}, 
		{Balance: 10_000, Active: true, PAN: "5058 xxxx xxxx 9991",},
		{Balance: 0, Active: true, PAN: "5058 xxxx xxxx 9992",},
		{Balance: 10_000, Active: false, PAN: "5058 xxxx xxxx 9993",},
		
				}

				
				srcPay := PaymentSources (cards)
                  for _, i := range srcPay {
                      fmt.Println(i.Number)
}

				//Output:
				//5058 xxxx xxxx 9999
				//5058 xxxx xxxx 9991
}