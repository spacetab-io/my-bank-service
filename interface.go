package main

type Currency string

const (
	CurrencySBP Currency = "SBP"
	CurrencyRUB Currency = "RUB"
)

type AccountInterface interface {
	// AddFunds Позволяет внести на счёт сумму sum
	AddFunds(sum float64)
	// SumProfit Рассчитывает процент по вкладу и полученные деньги вносит на счёт
	SumProfit()
	// Withdraw Производит списание со счёта по указанным правилам. Если списание выходит за рамки правил, выдаёт ошибку
	Withdraw(f float64) error
	// GetCurrency Выдаёт валюту счёта
	GetCurrency() Currency
	// GetAccountCurrencyRate Выдаёт курс валюты счёта к передаваемой валюте cur
	GetAccountCurrencyRate(cur Currency) float64
	// GetBalance Выдаёт баланс счёта в указанной валюте
	GetBalance(cur Currency) float64
}
