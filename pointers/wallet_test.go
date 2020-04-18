package pointers

import "testing"


func assertBalance(t *testing.T, wallet Wallet,
	expected Bitcoin) {
	t.Helper()
	res := wallet.Balance()
	if res != expected {
		t.Errorf("res: %d, expected: %d", res, expected)
	}
}

func assertError(t *testing.T, err error, expected string){
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didnt get one")
	}
	if err.Error() != expected {
		t.Errorf("res: %q, expected: %q", err,expected)
	}
}

func assertNoError(t *testing.T, err error){
	t.Helper()
	if err != nil{
		t.Fatal("got an error but didnt want one")
	}
}

func TestWallet(t *testing.T){

	t.Run("Deposit", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(10))
		assertNoError(t,err)
	})

	t.Run("Withdraw insufficient funds",
		func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t,wallet,startingBalance)
		assertError(t,err,"insufficient funds")
	})

}
