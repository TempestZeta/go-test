package main

import (
	"fmt"
	"learngo/accounts"
)

func main() {

	account := accounts.NewAccount("Testbed")
	fmt.Println(account)
}
