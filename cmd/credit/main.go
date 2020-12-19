package main

import (
	"bgo-homeworks-02/pkg/credit"
	"fmt"
)

func main() {

	amountCredit, periodCredit, percentPerYearCredit := 1_000_000_00, 36, 20

	monthlyPayment, overPayment, totalPayment := credit.Calculate(amountCredit, periodCredit, percentPerYearCredit)

	fmt.Println("Ежемесячный платёж кредита:", monthlyPayment)
	fmt.Println("Переплата по кредиту:", overPayment)
	fmt.Println("Общая сумма платежа:", totalPayment)
}
