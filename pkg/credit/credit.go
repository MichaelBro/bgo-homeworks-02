package credit

import "math"

func Calculate(amountCredit, periodCredit, percentPerYearCredit int) (monthlyPayment, overPayment, totalPayment int) {
	ratePerMonthPercent := float64(percentPerYearCredit) / 12 / 100

	exp := math.Pow(1+ratePerMonthPercent, float64(periodCredit))
	coefficient := (ratePerMonthPercent * exp) / (exp - 1)

	monthlyPayment = int(coefficient * float64(amountCredit))
	totalPayment = monthlyPayment * periodCredit
	overPayment = totalPayment - amountCredit

	return
}
