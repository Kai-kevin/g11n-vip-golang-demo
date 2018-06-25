package util

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"vipgoclient/com/vmware/i18n/client/format"
	"vipgoclient/com/vmware/i18n/client/bean/i18n"
)

func GetCurrencyFormatByLocal(locale string) (numberFormat format.NumberFormat, err error) {
	cacheformatMap := *GetFormatMap()
	cacheFormat := cacheformatMap[locale]

	if cacheFormat == nil {
		err = errors.New("Can not get cacheCurrencyFormat from locale " + locale)
		return
	}

	return GetNumberFormatByPattern(locale, cacheFormat.Messages.NumberFormats.CurrencyFormats)
}

func GetPercentFormatByLocal(locale string) (numberFormat format.NumberFormat, err error) {
	cacheformatMap := *GetFormatMap()
	cacheFormat := cacheformatMap[locale]

	if cacheFormat == nil {
		err = errors.New("Can not get cachePercentFormat from locale " + locale)
		return
	}

	return GetNumberFormatByPattern(locale, cacheFormat.Messages.NumberFormats.PercentFormats)
}

func GetNumberFormatByLocal(locale string) (numberFormat format.NumberFormat, err error) {
	//cacheformatMap := *GetFormatMap()
	//cacheFormat := cacheformatMap[locale]

	cacheformatValue,_ := GetFormat4ParalMap().Load(locale)
	cacheFormat := cacheformatValue.(*bean.QueryFormattingPatternByLocaleRespData)

	if cacheFormat == nil {
		err = errors.New("Can not get cacheNumberFormat from locale " + locale)
		return
	}

	return GetNumberFormatByPattern(locale, cacheFormat.Messages.NumberFormats.DecimalFormats)
}

func GetNumberFormatByPattern(locale string, formatPattern string) (numberFormat format.NumberFormat, err error) {

	numberFormat = *new(format.NumberFormat)

	numberFormat.Pattern = formatPattern

	pos := strings.Index(formatPattern, ".")

	if pos != -1 {
		pos2 := strings.LastIndex(formatPattern, "0")
		if pos2 > pos {
			numberFormat.MinDecimalDigits = pos2 - pos
		}

		pos3 := strings.LastIndex(formatPattern, "#")

		if pos3 > pos2 {
			numberFormat.MaxDecimalDigits = pos3 - pos
		} else {
			numberFormat.MaxDecimalDigits = numberFormat.MinDecimalDigits
		}

		formatPattern = formatPattern[0:pos]
	}

	p := strings.Replace(formatPattern, ",", "", -1)

	pos = strings.Index(p, "0")
	if pos != -1 {
		numberFormat.MinIntegerDigits = strings.LastIndex(p, "0") - pos + 1
	}

	p = strings.Replace(formatPattern, "#", "0", -1)
	pos = strings.LastIndex(formatPattern, ",")
	if pos != -1 {
		numberFormat.GroupSizeFinal = strings.LastIndex(p, "0") - pos
		pos2 := strings.LastIndex(p[0:pos], ",")
		if pos2 != -1 {
			numberFormat.GroupSizeMain = pos - pos2 - 1
		} else {
			numberFormat.GroupSizeMain = numberFormat.GroupSizeFinal
		}
	}

	cacheformatMap := *GetFormatMap()
	cacheFormat := cacheformatMap[locale]

	//init some params
	numberFormat.NegativePrefix = cacheFormat.Messages.NumberSymbols.MinusSign
	numberFormat.DecimalSymbol = cacheFormat.Messages.NumberSymbols.Decimal
	numberFormat.ExponentialSymbol = cacheFormat.Messages.NumberSymbols.Exponential
	numberFormat.CurrencySymbol = cacheFormat.Messages.CurrencySymbol
	numberFormat.GroupSymbol = cacheFormat.Messages.NumberSymbols.Group

	//percentPattern := cacheFormat.Messages.NumberFormats.PercentFormats
	//
	//if strings.Index(percentPattern, "%") != -1 {
	//	numberFormat.Multiplier = 100
	//	numberFormat.Percent = "%"
	//} else if strings.Index(percentPattern, "‰") != -1 {
	//	numberFormat.Multiplier = 1000
	//} else {
	//	numberFormat.Multiplier = 1
	//	numberFormat.Permille = "‰"
	//}
	numberFormat.Multiplier = 1

	return
}

//format number
//func FormatNumberDefault(obj interface{},local string) (str string,err error){
//	format := "###,##0.###"
//
//	fmt.Sprintf(format)
//
//	return "",nil
//}

//get percent format
func FormatPercent(format *format.NumberFormat, number float64) string {
	format.Multiplier = 100

	format.Percent = "%"

	r, _ := regexp.Compile("[\\.\\,#0]+")

	strs := r.FindStringSubmatch(format.Pattern)

	if len(strs) == 0 {
		return ""
	}

	ret := strings.Replace(format.Pattern, strs[0], FormatNumber(format, number), -1)

	return strings.Replace(ret, "%", format.Percent, -1)
}

// get currency number format
func FormatCurrency(format *format.NumberFormat, number float64) string {

	r, _ := regexp.Compile("[\\.\\,#0]+")

	strs := r.FindStringSubmatch(format.Pattern)

	if len(strs) == 0 {
		return ""
	}

	ret := strings.Replace(format.Pattern, strs[0], FormatNumber(format, number), -1)

	return strings.Replace(ret, "¤", format.CurrencySymbol, -1)
}

func FormatNumber(format *format.NumberFormat, number float64) string {
	negative := number < 0

	value := math.Abs(number * float64(format.Multiplier))
	stringValue := ""

	if format.MaxDecimalDigits >= 0 {
		stringValue = numberRound(value, format.MaxDecimalDigits)
	} else {
		stringValue = fmt.Sprintf("%f", value)
	}

	pos := strings.Index(stringValue, ".")

	integer := stringValue
	decimal := ""
	if pos != -1 {
		integer = stringValue[:pos]
		decimal = stringValue[pos+1:]
	}

	for len(decimal) < format.MinDecimalDigits {
		decimal = decimal + "0"
	}

	for len(integer) < format.MinIntegerDigits {
		integer = "0" + integer
	}

	if len(decimal) > 0 {
		decimal = format.DecimalSymbol + decimal
	}

	// put the integer portion into properly sized groups
	if format.GroupSizeFinal > 0 && len(integer) > format.GroupSizeFinal {
		if len(integer) > format.GroupSizeMain {
			groupFinal := integer[len(integer)-format.GroupSizeFinal:]
			groupFirst := integer[:len(integer)-format.GroupSizeFinal]
			integer = strings.Join(chunkString(groupFirst, format.GroupSizeMain), format.GroupSymbol) + format.GroupSymbol + groupFinal
		}
	}

	// append/prepend negative/positive prefix/suffix
	formatted := ""
	if negative {
		formatted = format.NegativePrefix + integer + decimal + format.NegativeSuffix
	} else {
		formatted = format.PositivePrefix + integer + decimal + format.PositiveSuffix
	}

	return formatted
}

func numberRound(number float64, decimals int) string {

	//Decimal
	if number == float64(int64(number)) {
		return strconv.FormatInt(int64(number), 10)
	}

	str := fmt.Sprintf("%f", number)
	pos := strings.Index(str, ".")

	if pos != -1 && len(str) > (pos+decimals) {

		num, _ := strconv.ParseFloat(str, 64)

		num = num * (math.Pow(10, float64(decimals)))

		convNum := math.Round(num)

		value := convNum / (math.Pow(10, float64(decimals)))

		return strconv.FormatFloat(value, 'f', -1, 64)
	} else {
		return strconv.FormatFloat(number, 'f', -1, 64)
	}
}

// chunkString takes a string and chunks it into size-sized pieces in a slice.
// If the length of the string is not divisible by the size, then the first
// chunk in the slice will be padded to compensate.
func chunkString(str string, size int) []string {
	if str == "" {
		return []string{}
	}

	if size == 0 {
		return []string{str}
	}

	chunks := make([]string, int64(math.Ceil(float64(len(str))/float64(size))))

	for len(str) < len(chunks)*size {
		str = " " + str
	}

	for i := 0; i < len(chunks); i++ {
		start := i * size
		stop := int64(math.Min(float64(start+size), float64(len(str))))
		chunks[i] = str[start:stop]
	}

	chunks[0] = strings.TrimLeft(chunks[0], " ")

	return chunks
}
