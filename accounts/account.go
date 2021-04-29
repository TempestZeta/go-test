package accounts

type account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 10}
	return &account
}
