package domain

import (
	"github.com/google/uuid"
)

type Account struct {
	userUUID    uuid.UUID
	accountUUID uuid.UUID
	status      AccountStatus
	balance     Money
}

type CheckingAccount struct {
	Account
}

type SavingsAccount struct {
	Account
	BonusProgram BonusProgram
}

type CreditAccount struct {
	Account
	overdraftLimit Money
}

type AccountStatus int

const (
	Active AccountStatus = iota + 1
	Frozen
	Closed
)

func NewAccount(userUUID uuid.UUID, accountUUID uuid.UUID, curr Currency) Account {
	return Account{
		userUUID:    userUUID,
		accountUUID: accountUUID,
		status:      Active,
		balance:     Money{amount: 0, curr: curr},
	}
}

func NewCheckingAccount(userUUID uuid.UUID, accountUUID uuid.UUID, curr Currency) *CheckingAccount {
	return &CheckingAccount{Account: NewAccount(userUUID, accountUUID, curr)}
}

func NewSavingsAccount(userUUID uuid.UUID, accountUUID uuid.UUID, curr Currency, tier BonusTier) *SavingsAccount {
	return &SavingsAccount{
		Account:      NewAccount(userUUID, accountUUID, curr),
		BonusProgram: NewBonusProgram(tier),
	}
}

func NewCreditAccount(userUUID uuid.UUID, accountUUID uuid.UUID, curr Currency, overdraftLimit Money) *CreditAccount {
	return &CreditAccount{
		Account:        NewAccount(userUUID, accountUUID, curr),
		overdraftLimit: overdraftLimit,
	}
}

func (a *Account) UserUUID() uuid.UUID {
	return a.userUUID
}

func (a *Account) AccountUUID() uuid.UUID {
	return a.accountUUID
}

func (a *Account) Status() AccountStatus {
	return a.status
}

func (a *Account) SetStatus(status AccountStatus) {
	a.status = status
}

func (a *Account) Balance() Money {
	return a.balance
}

func (a *Account) Deposit(value Money) {
	if a.status != Active {
		return
	}
	a.balance = a.balance.Add(value)
}

func (s *SavingsAccount) Deposit(value Money) {
	s.Account.Deposit(value)
	s.BonusProgram.Accrue(value)
}

func (a *Account) Withdraw(value Money) bool {
	if !a.CanWithdraw(value) {
		return false
	}
	a.balance = a.balance.Sub(value)
	return true
}

func (a *Account) CanWithdraw(value Money) bool {
	if a.status != Active {
		return false
	}
	return a.balance.amount >= value.amount
}

func (c *CreditAccount) CanWithdraw(value Money) bool {
	if c.status != Active {
		return false
	}
	return c.balance.amount+c.overdraftLimit.amount >= value.amount
}
