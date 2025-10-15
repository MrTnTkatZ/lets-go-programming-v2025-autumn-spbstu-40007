package xmldecoder

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/tntkatz/task-3/internal/config"
	"golang.org/x/text/encoding/charmap"
)

func DecodeXML(inputFile []byte, valCurs *config.ValCurs) error {
	decoder := xml.NewDecoder(bytes.NewReader(inputFile))

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.ToLower(charset) == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%s", charset)
	}

	err := decoder.Decode(&valCurs)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
