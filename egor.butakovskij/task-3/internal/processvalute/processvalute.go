package processvalute

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tntkatz/task-3/internal/config"
)

func ProcessValute(valCurs config.ValCurs, processedValutes []config.ProcessedValute) error {
	for _, valute := range valCurs.Valute {
		newValue := strings.Replace(valute.Value, ",", ".", 1)

		sortValue, err := strconv.ParseFloat(newValue, 64)
		if err != nil {
			return fmt.Errorf("%w", err)
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

	return nil
}
