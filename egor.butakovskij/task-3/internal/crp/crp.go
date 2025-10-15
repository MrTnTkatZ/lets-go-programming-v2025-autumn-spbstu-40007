package crp

import "github.com/tntkatz/task-3/internal/config"

func CurrencyProcess(processedValutes []config.ProcessedValute) []config.CurrencyResult {
	currencyResults := make([]config.CurrencyResult, 0, len(processedValutes))

	for _, pVal := range processedValutes {
		currencyResult := config.CurrencyResult{
			NumCode:  pVal.NumCode,
			CharCode: pVal.CharCode,
			Value:    pVal.SortValue,
		}

		currencyResults = append(currencyResults, currencyResult)
	}

	return currencyResults
}
