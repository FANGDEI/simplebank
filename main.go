package main

import (
	"log"

	"github.com/FANGDEI/simplebank/ecrypto"
	"github.com/FANGDEI/simplebank/gateway"
	"github.com/FANGDEI/simplebank/store/cache"
	"github.com/FANGDEI/simplebank/store/local"
	"github.com/go-playground/validator/v10"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func main() {
	var err error

	gateway.C.Cryptoer = ecrypto.New()
	gateway.C.Validator = validator.New()

	gateway.C.Cacher, err = cache.New()
	if err != nil {
		log.Fatalln(err)
	}

	gateway.C.Localer, gateway.C.Storer, err = local.New()
	if err != nil {
		log.Fatalln(err)
	}

	m := gateway.New()
	err = m.Run()
	if err != nil {
		log.Fatalln(err)
	}
	// s := gateway.C.Storer
	// results, errs := make(chan local.TransferTxResult), make(chan error)
	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		result, err := s.TransferTx(local.TransferTxParams{
	// 			FromAccountID: 1,
	// 			ToAccountID:   2,
	// 			Amount:        10,
	// 		})
	// 		results <- result
	// 		errs <- err
	// 	}()
	// }
	// for i := 0; i < 5; i++ {
	// 	result, err := <-results, <-errs
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	log.Printf("%v %v\n", result.FromAccount.Balance, result.ToAccount.Balance)
	// }
}
