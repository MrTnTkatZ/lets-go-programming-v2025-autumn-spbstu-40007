package vp

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tntkatz/task-3/internal/config"
)

func ValuteProcess(valCurs config.ValCurs) ([]config.ProcessedValute, error) {
	processedValutes := make([]config.ProcessedValute, 0, len(valCurs.Valute))

	for _, valute := range valCurs.Valute {
		newValue := strings.Replace(valute.Value, ",", ".", 1)

		sortValue, err := strconv.ParseFloat(newValue, 64)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		processedValute := config.ProcessedValute{
			ID:        valute.ID,
			NumCode:   valute.NumCode,
			CharCode:  valute.CharCode,
			Nominal:   valute.Nominal,
			Name:      valute.Name,
			Value:     valute.Value,
			VunitRate: valute.VunitRate,
			SortValue: sortValue,
		}

		processedValutes = append(processedValutes, processedValute)
	}

	sort.Sort(config.ByValue(processedValutes))

	return processedValutes, nil
}
