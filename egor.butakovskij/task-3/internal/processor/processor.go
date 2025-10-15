package processor

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tntkatz/task-3/internal/config"
	"github.com/tntkatz/task-3/internal/crp"
	"github.com/tntkatz/task-3/internal/pathcreator"
	"github.com/tntkatz/task-3/internal/vp"
	"github.com/tntkatz/task-3/internal/xmldecoder"
	"gopkg.in/yaml.v3"
)

const DefaultFilePermissions = 0o600

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

	processedValutes, err := vp.ValuteProcess(valCurs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	currencyResults := crp.CurrencyProcess(processedValutes)

	jsonData, err := json.Marshal(currencyResults)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = pathcreator.PathCreator(cfg.OutputFile)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = os.WriteFile(cfg.OutputFile, jsonData, DefaultFilePermissions)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
