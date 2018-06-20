package format

type NumberFormat struct {
	PositivePrefix    string
	PositiveSuffix    string
	NegativePrefix    string
	NegativeSuffix    string
	Multiplier        int
	MinDecimalDigits  int
	MaxDecimalDigits  int
	MinIntegerDigits  int
	GroupSizeFinal    int // only the right-most (least significant) group
	GroupSizeMain     int // all other groups
	Percent           string
	Permille          string
	DecimalSymbol     string
	GroupSymbol       string
	ExponentialSymbol string
	CurrencySymbol	  string
}
