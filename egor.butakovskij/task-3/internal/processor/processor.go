package processor

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/tntkatz/task-3/internal/config"
	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

func Run(configPath string) error {

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	cfg := config.Config{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	inputFile, err := os.ReadFile(cfg.InputFile)
	if err != nil {
		return err
	}

	valCurs := config.ValCurs{}

	decoder := xml.NewDecoder(bytes.NewReader(inputFile))

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.ToLower(charset) == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf(charset)
	}

	err = decoder.Decode(&valCurs)
	if err != nil {
		return err
	}

	var processedValutes []config.ProcessedValute

	for _, valute := range valCurs.Valute {
		newValue := strings.Replace(valute.Value, ",", ".", 1)

		sortValue, err := strconv.ParseFloat(newValue, 64)
		if err != nil {
			return err
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

	var currencyResults []config.CurrencyResult

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
		return err
	}

	outputDir := filepath.Dir(cfg.OutputFile)
	if outputDir != "" {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return err
		}
	}

	err = os.WriteFile(cfg.OutputFile, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
