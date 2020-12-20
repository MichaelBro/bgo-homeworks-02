package credit

import "math"

func Calculate(amountCredit, periodCredit, percentPerYearCredit int) (monthlyPayment, overPayment, totalPayment int) {
	ratePerMonthPercent := math.Round((float64(percentPerYearCredit)/12/100)*1000) / 1000

	exp := math.Pow(1+ratePerMonthPercent, float64(periodCredit))
	coefficient := (ratePerMonthPercent * exp) / (exp - 1)

	monthlyPayment = int(math.Round(coefficient*1_000_000)) * amountCredit / 1_000_000
	totalPayment = monthlyPayment * periodCredit
	overPayment = totalPayment - amountCredit

	return
}
