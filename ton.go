package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tonkeeper/tonapi-go"
)

func tonGetAccoutBalance() {
	client, err := tonapi.New()
	if err != nil {
		log.Fatal(err)
	}
	account, err := client.GetAccount(context.Background(), tonapi.GetAccountParams{AccountID: "EQC12muMPDLSfjzlNBgRz5gs3TEswm6k59A_rkdFNkK7PInS"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account ton amount: %v\n", account.Balance)
}


