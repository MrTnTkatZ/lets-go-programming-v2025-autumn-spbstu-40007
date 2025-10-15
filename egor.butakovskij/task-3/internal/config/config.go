package config

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

type Valute struct {
	ID        string `xml:"ID,attr"`
	NumCode   int    `xml:"NumCode"`
	CharCode  string `xml:"CharCode"`
	Nominal   string `xml:"Nominal"`
	Name      string `xml:"Name"`
	Value     string `xml:"Value"`
	VunitRate string `xml:"VunitRate"`
}

type ProcessedValute struct {
	ID        string  `xml:"ID,attr"`
	NumCode   int     `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   string  `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     string  `xml:"Value"`
	VunitRate string  `xml:"VunitRate"`
	SortValue float64 `xml:"SortValue"`
}

type ValCurs struct {
	Date   string   `xml:"Date,attr"`
	Name   string   `xml:"name,attr"`
	Valute []Valute `xml:"Valute"`
}

type CurrencyResult struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type ByValue []ProcessedValute

func (a ByValue) Len() int {
	return len(a)
}

func (a ByValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByValue) Less(i, j int) bool {
	return a[i].SortValue > a[j].SortValue
}
