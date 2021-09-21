package payment

import (

	"bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {

	max:= types.Payment{}
 
	for i:=0; i<len(payments); i++ {

		
		if payments[0].Amount <= payments[i].Amount {
			
			max = payments[i]
			
			
		}
	}
       return  max
}