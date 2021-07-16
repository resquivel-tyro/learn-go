package pointers

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

//When calling func (w Wallet) Deposit(amount int)
//the w is a copy of whatever we called the method from.

//So rather than taking a copy of the whole Wallet,
//we instead take a pointer to that wallet
//so that we can change the original values within it.
//The difference is the receiver type is *Wallet rather than Wallet which you can read as
//"a pointer to a wallet".


func (w *Wallet) Deposit(i Bitcoin) {
	w.balance += i
	// the language permits us to write w.balance, without an explicit dereference like (*w).balance
	// these are struct pointers and are automatically dereferenced
}

func (w *Wallet) Balance() (i Bitcoin) {
	return w.balance
	//	Technically you do not need to change Balance to use a pointer receiver as taking a copy of the balance is fine.
	//	However, by convention you should keep your method receiver types the same for consistency.
}

