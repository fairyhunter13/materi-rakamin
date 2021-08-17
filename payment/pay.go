package payment

import (
	"fmt"
	"regexp"
)

type IPayment interface {
	Validate() (err error)
	Pay() string
}

const (
	TypeWallet       = "wallet"
	TypeCC           = "cc"
	TypeBankTransfer = "bank_transfer"
)

func New(paymentType, arg string) (pay IPayment) {
	switch paymentType {
	case TypeWallet:
		pay = &Wallet{arg}
	case TypeCC:
		pay = &CC{arg}
	case TypeBankTransfer:
		pay = &BankTransfer{arg}
	}
	return
}

type Wallet struct{ PhoneNumber string }

func (w *Wallet) Pay() string {
	return fmt.Sprintf("Pembayaran dengan metode wallet #%s sukses!", w.PhoneNumber)
}

var regexPhone = regexp.MustCompile("[0-9]{12,}")

func (w *Wallet) Validate() (err error) {
	if !regexPhone.MatchString(w.PhoneNumber) {
		err = fmt.Errorf("Invalid phone number: %s", w.PhoneNumber)
	}
	return
}

type CC struct{ CreditCardNumber string }

var regexCreditCard = regexp.MustCompile("[0-9]{12}")

func (cc *CC) Pay() string {
	return fmt.Sprintf("Pembayaran dengan metode kartu kredit #%s sukses!", cc.CreditCardNumber)
}
func (cc *CC) Validate() (err error) {
	if !regexCreditCard.MatchString(cc.CreditCardNumber) {
		err = fmt.Errorf("Invalid cerdit card number: %s", cc.CreditCardNumber)
	}
	return
}

type BankTransfer struct{ AccountNumber string }

func (bt *BankTransfer) Pay() string {
	return fmt.Sprintf("Pembayaran dengan metode bank transfer #%s sukses!", bt.AccountNumber)
}

var regexAccountNumber = regexp.MustCompile("[0-9]{6,}")

func (bt *BankTransfer) Validate() (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = rec.(error)
		}
	}()
	if !regexAccountNumber.MatchString(bt.AccountNumber) {
		panic(fmt.Errorf("Invalid account number: %s", bt.AccountNumber))
	}
	return
}
