package payment 

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampleMax() {
	payments := []types.Payment {
		types.Payment{ID:3, Amount: 10,}, 
		types.Payment{ID:2, Amount: 40,},
		types.Payment{ID:1, Amount: 50,},
		types.Payment{ID:4, Amount: 50,},
		types.Payment{ID:5, Amount: 50,},
				}

				max :=Max(payments)
				fmt.Println(max.ID)

				//Output:
				//5
}