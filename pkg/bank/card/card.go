package payment

import (

	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {

	

	Numb :=  []types.PaymentSource{ }
    
	
	for i := 0; i < len(cards); i++ {

       
		if !cards[i].Active {
			continue
		}
		if cards[i].Balance<=0 {
			continue
		}
		
		exemple1 := types.PaymentSource{ Number: cards[i].PAN , Balance: cards[i].Balance }
       Numb = append(Numb, exemple1)
	}
return Numb
	
}