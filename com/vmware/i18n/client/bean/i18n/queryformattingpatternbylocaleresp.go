package bean

type QueryFormattingPatternByLocaleRespEvent struct {
	Response  Response                      `json:"response"`
	Signature string                        `json:"signature"`
	Data      QueryFormattingPatternByLocaleRespData `json:"data"`
}

type QueryFormattingPatternByLocaleRespData struct {
	ProductName string  `json:"productName"`
	Version     string  `json:"version"`
	Pseudo      bool    `json:"pseudo"`
	Component   string  `json:"component"`
	Messages    Message `json:"messages"`
	Locale      string  `json:"locale"`
	Status      string  `json:status`
}

type Message struct {
	LocaleID             string               `json:"localeID"`
	DayPeriodsFormat     DayPeriodsFormat     `json:"dayPeriodsFormat"`
	DayPeriodsStandalone DayPeriodsStandalone `json:"dayPeriodsStandalone"`
	DaysFormat           DaysFormat           `json:"daysFormat"`
	DaysStandalone       DaysStandalone       `json:"daysStandalone"`
	MonthsFormat         MonthsFormat         `json:"monthsFormat"`
	MonthsStandalone     MonthsStandalone     `json:"monthsStandalone"`
	Eras                 Eras                 `json:"eras"`
	FirstDayOfWeek       int                  `json:"firstDayOfWeek"`
	WeekendRange         []int                `json:"weekendRange"`
	DateFormat           DateFormat           `json:"dateFormat"`
	TimeFormat           TimeFormat           `json:"timeFormat"`
	DateTimeFormat       DateTimeFormat       `json:"dateTimeFormat"`
	NumberSymbols        NumberSymbols        `json:"NumberSymbols"`
	NumberFormats        NumberFormats        `json:"NumberFormats"`
	CurrencySymbol       string               `json:"currencySymbol"`
	currencyName         string               `json:"CurrencyName"`
	PluralRules          PluralRules          `json:"PluralRules"`
}

type DayPeriodsFormat struct {
	Narrow      []string `json:"narrow"`
	Abbreviated []string `json:"abbreviated"`
	Wide        []string `json:wide`
}

type DayPeriodsStandalone struct {
	Narrow      []string `json:"narrow"`
	Abbreviated []string `json:"abbreviated"`
	Wide        []string `json:wide`
}

type Abbreviated struct {
	Thu string `json:"thu"`
	Tue string `json:"tue"`
	Wed string `json:"wed"`
	Sat string `json:"sat"`
	Fri string `json:"fri"`
	Sun string `json:"sun"`
	Mon string `json:"mon"`
}

type DaysFormat struct {
	narrow      []string    `json:"narrow"`
	Abbreviated Abbreviated `json:"abbreviated"`
	Wide        []string    `json:"wide"`
	Short       []string    `json:"short"`
}

type DaysStandalone struct {
	narrow      []string    `json:"narrow"`
	Abbreviated []string `json:"abbreviated"`
	Wide        []string    `json:"wide"`
	Short       []string    `json:"short"`
}

type DateFormat struct {
	Short  string `json:"short"`
	Medium string `json:"medium"`
	Long   string `json:"long"`
	Full   string `json:"full"`
}

type MonthsFormat struct {
	Narrow      []string `json:"narrow"`
	Abbreviated []string `json:"narrow"`
	Wide        []string `json:"narrow"`
}

type Eras struct {
	Narrow      []string `json:"narrow"`
	Abbreviated []string `json:"abbreviated"`
	Wide        []string `json:"wide"`
}

type MonthsStandalone struct {
	Narrow      []string `json:"narrow"`
	Abbreviated []string `json:"abbreviated"`
	Wide        []string `json:"wide"`
}

type TimeFormat struct {
	Short  string `json:"short"`
	Medium string `json:"medium"`
	Long   string `json:"long"`
	Full   string `json:"full"`
}

type DateTimeFormat struct {
	Short  string `json:"short"`
	Medium string `json:"medium"`
	Long   string `json:"long"`
	Full   string `json:"full"`
}

type NumberSymbols struct {
	Decimal                string `json:"decimal"`
	Group                  string `json:"group"`
	List                   string `json:"list"`
	PercentSign            string `json:"percentSign"`
	PlusSign               string `json:"plusSign"`
	MinusSign              string `json:"minusSign"`
	Exponential            string `json:"exponential"`
	SuperscriptingExponent string `json:"superscriptingExponent"`
	PerMille               string `json:"perMille"`
	Infinity               string `json:"infinity"`
	Nan                    string `json:"nan"`
	TimeSeparator          string `json:"timeSeparator"`
}

type NumberFormats struct {
	DecimalFormats    string `json:"decimalFormats"`
	PercentFormats    string `json:"percentFormats"`
	CurrencyFormats   string `json:"currencyFormats"`
	ScientificFormats string `json:"scientificFormats"`
}

type PluralRules struct {
	PluralRule_count_other string `json:"pluralRule-count-other"`
}
