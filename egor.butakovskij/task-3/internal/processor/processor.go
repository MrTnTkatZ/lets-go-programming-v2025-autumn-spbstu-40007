package processor

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/tntkatz/task-3/internal/config"
	"github.com/tntkatz/task-3/internal/pathcreator"
	"github.com/tntkatz/task-3/internal/xmldecoder"
	"gopkg.in/yaml.v3"
)

func Run(configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	cfg := config.Config{
		InputFile:  "",
		OutputFile: "",
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	inputFile, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	valCurs := config.ValCurs{
		Date:   "",
		Name:   "",
		Valute: nil,
	}

	err = xmldecoder.DecodeXML(inputFile, &valCurs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	processedValutes := make([]config.ProcessedValute, 0, len(valCurs.Valute))

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

	currencyResults := make([]config.CurrencyResult, 0, len(processedValutes))

	for _, pVal := range processedValutes {
		currencyResult := config.CurrencyResult{
			NumCode:  pVal.NumCode,
			CharCode: pVal.CharCode,
			Value:    pVal.SortValue,
		}

		currencyResults = append(currencyResults, currencyResult)
	}

	jsonData, err := json.Marshal(currencyResults)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = pathcreator.PathCreator(cfg.OutputFile)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = os.WriteFile(cfg.OutputFile, jsonData, 0600)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
