package main

import (
	"log"

	"github.com/FANGDEI/simplebank/store/local"
)

func main() {
	m, _ := local.NewStore()

	errs := make(chan error)
	results := make(chan local.TransferTxResult)

	for i := 0; i < 5; i++ {
		go func() {
			result, err := m.TransferTx(local.TransferTxParams{
				FromAccountID: 1,
				ToAccountID:   2,
				Amount:        10,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < 5; i++ {
		err := <-errs
		result := <-results

		if err != nil {
			log.Println(err)
			log.Print(i)
		}

		log.Printf("%v %v\n", result.FromAccount.Balance, result.ToAccount.Balance)
	}
}
