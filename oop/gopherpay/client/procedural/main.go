package main

import (
	"awesomeProject/oop/payment/procedural"
	"fmt"
	"strings"
)

func main() {
	const amount = 500
	fmt.Println("Paying with cash")
	procedural.PayWithCash(amount)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	credit := &procedural.CreditCard{
		OwnerName:       "Debra Ingram",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 5,
		ExpirationYear:  2021,
		SecurityCode:    123,
		AvailableCredit: 5000,
	}
	fmt.Println("Paying with credit")
	fmt.Printf("Initial balance: $%.2f\n", credit.AvailableCredit)
	procedural.PayWithCredit(credit, amount)
	fmt.Printf("balance now: $%.2f\n", credit.AvailableCredit)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := &procedural.CheckingAccount{
		AccountOwner:  "Debra Ingram",
		RoutingNumber: "01010011",
		AccountNumber: "0551005115100",
		Balance:       250,
	}
	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("balance now: $%.2f\n", checking.Balance)

	fmt.Println("Hmm, not enough in the account. We can fix that!")
	fmt.Println("Changing account balance")
	checking.Balance = 5000

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: $%.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("balance now: $%.2f\n", checking.Balance)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
}
