package util

import (
	"time"
	"fmt"
	"errors"
	"strings"
	"vipgoclient/com/vmware/i18n/client/constants"
)

// Characters with special meaning in a datetime string:
// Technically, all a-z,A-Z characters should be treated as if they represent a
// datetime unit - but not all actually do. Any a-z,A-Z character that is
// intended to be rendered as a literal a-z,A-Z character should be surrounded
// by single quotes. There is currently no support for rendering a single quote
// literal.
const (
	datetimeFormatUnitEra       = 'G'
	datetimeFormatUnitYear      = 'y'
	datetimeFormatUnitMonth     = 'M'
	datetimeFormatUnitDayOfWeek = 'E'
	datetimeFormatUnitDay       = 'd'

	datetimeFormatUnitHour12    = 'h'
	datetimeFormatUnitHour24    = 'H'
	datetimeFormatUnitMinute    = 'm'
	datetimeFormatUnitSecond    = 's'
	datetimeFormatUnitPeriod    = 'a'
	datetimeForamtUnitQuarter   = 'Q'
	datetimeFormatUnitTimeZone1 = 'z'
	datetimeFormatUnitTimeZone2 = 'v'

)

var dateUnits = []string{string(datetimeFormatUnitEra),string(datetimeFormatUnitYear),string(datetimeFormatUnitMonth),string(datetimeFormatUnitDayOfWeek),
string(datetimeFormatUnitDay),string(datetimeFormatUnitHour12),string(datetimeFormatUnitHour24),string(datetimeFormatUnitMinute),
string(datetimeFormatUnitSecond),string(datetimeFormatUnitPeriod),string(datetimeForamtUnitQuarter),string(datetimeFormatUnitTimeZone1),
string(datetimeFormatUnitTimeZone2),}

// The sequence length of datetime unit characters indicates how they should be
// rendered.
const (
	datetimeFormatLength1Plus       = 1
	datetimeFormatLength2Plus       = 2
	datetimeFormatLengthAbbreviated = 3
	datetimeFormatLengthWide        = 4
	datetimeFormatLengthNarrow      = 5
)

// datetime formats are a sequences off datetime components and string literals
const (
	datetimePatternComponentUnit = iota
	datetimePatternComponentLiteral
)

const (
	stdTZ                    = iota                // "MST"
	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
	stdISO8601SecondsTZ                            // "Z070000"
	stdISO8601ShortTZ                              // "Z07"
	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
	stdNumTZ                                       // "-0700"  // always numeric
	stdNumSecondsTz                                // "-070000"
	stdNumShortTZ                                  // "-07"    // always numeric
	stdNumColonTZ                                  // "-07:00" // always numeric
	stdNumColonSecondsTZ                           // "-07:00:00"
)

// A list of currently unsupported units:
// These still need to be implemented. For now they are ignored.
var (
	datetimeFormatUnitCutset = []rune{
		datetimeFormatUnitEra,
		datetimeForamtUnitQuarter,
		datetimeFormatUnitTimeZone2,
	}
)

type datetimePatternComponent struct {
	pattern       string
	componentType int
}

// FormatDateTime takes a time struct and a format and returns a formatted
// string. Callers should use a DateFormat, TimeFormat, or DateTimeFormat
// constant.
func FormatDateTime(format int, datetime time.Time, locale string) (string, error) {
	pattern := ""
	switch format {
	case constants.FULLDATE:
		pattern = cacheFormatMap[locale].Messages.DateFormat.Full
	case constants.LONGDATE:
		pattern = cacheFormatMap[locale].Messages.DateFormat.Long
	case constants.MEDIUMDATE:
		pattern = cacheFormatMap[locale].Messages.DateFormat.Medium
	case constants.SHORTDATE:
		pattern = cacheFormatMap[locale].Messages.DateFormat.Short
	case constants.FULLTIME:
		pattern = cacheFormatMap[locale].Messages.TimeFormat.Full
	case constants.LONGTIME:
		pattern = cacheFormatMap[locale].Messages.TimeFormat.Long
	case constants.MEDIUMTIME:
		pattern = cacheFormatMap[locale].Messages.TimeFormat.Medium
	case constants.SHORTTIME:
		pattern = cacheFormatMap[locale].Messages.TimeFormat.Short
	case constants.FULL:
		tempPattern := strings.Replace(cacheFormatMap[locale].Messages.DateTimeFormat.Full, "{0}",cacheFormatMap[locale].Messages.DateFormat.Full,-1)
		pattern = strings.Replace(tempPattern, "{1}",cacheFormatMap[locale].Messages.TimeFormat.Full,-1)
	case constants.LONG:
		tempPattern := strings.Replace(cacheFormatMap[locale].Messages.DateTimeFormat.Long, "{0}",cacheFormatMap[locale].Messages.DateFormat.Long,-1)
		pattern = strings.Replace(tempPattern, "{1}",cacheFormatMap[locale].Messages.TimeFormat.Long,-1)
	case constants.MEDIUM:
		tempPattern := strings.Replace(cacheFormatMap[locale].Messages.DateTimeFormat.Medium, "{0}",cacheFormatMap[locale].Messages.DateFormat.Medium,-1)
		pattern = strings.Replace(tempPattern, "{1}",cacheFormatMap[locale].Messages.TimeFormat.Medium,-1)
	case constants.SHORT:
		tempPattern := strings.Replace(cacheFormatMap[locale].Messages.DateTimeFormat.Short, "{0}",cacheFormatMap[locale].Messages.DateFormat.Short,-1)
		pattern = strings.Replace(tempPattern, "{1}",cacheFormatMap[locale].Messages.TimeFormat.Short,-1)
	default:
		return "", errors.New("unknown datetime format" + pattern[0:1])
	}

	parsed, err := parseDateTimeFormat(pattern)
	if err != nil {
		return "", err
	}

	return formatDateTime(datetime, parsed,locale)
}

// formatDateTime takes a time.Time and a sequence of parsed pattern components
// and returns an internationalized string representation.
func formatDateTime(datetime time.Time, pattern []*datetimePatternComponent,locale string) (string, error) {
	formatted := ""
	for _, component := range pattern {
		if component.componentType == datetimePatternComponentLiteral {
			formatted += component.pattern
		} else {
			f, err := formatDateTimeComponent(datetime, component.pattern,locale)
			if err != nil {
				return "", err
			}
			formatted += f
		}
	}

	return strings.Trim(formatted, " ,"), nil
}

// formatDateTimeComponent renders a single component of a datetime format
// pattern.
func formatDateTimeComponent(datetime time.Time, pattern string,locale string) (string, error) {

	switch pattern[0:1] {
	case string(datetimeFormatUnitEra):
		//TODO
		fallthrough
	case string(datetimeFormatUnitYear):
		return formatDateTimeComponentYear(datetime, len(pattern))
	case string(datetimeFormatUnitMonth):
		return formatDateTimeComponentMonth(datetime, len(pattern),locale)
	case string(datetimeFormatUnitDayOfWeek):
		return formatDateTimeComponentDayOfWeek(datetime, len(pattern),locale)
	case string(datetimeFormatUnitDay):
		return formatDateTimeComponentDay(datetime, len(pattern))
	case string(datetimeFormatUnitHour12):
		return formatDateTimeComponentHour12(datetime, len(pattern))
	case string(datetimeFormatUnitHour24):
		return formatDateTimeComponentHour24(datetime, len(pattern))
	case string(datetimeFormatUnitMinute):
		return formatDateTimeComponentMinute(datetime, len(pattern))
	case string(datetimeFormatUnitSecond):
		return formatDateTimeComponentSecond(datetime, len(pattern))
	case string(datetimeFormatUnitPeriod):
		return formatDateTimeComponentPeriod(datetime, len(pattern),locale)
	case string(datetimeForamtUnitQuarter):
		//TODO
		fallthrough
	case string(datetimeFormatUnitTimeZone1):
		return formatDateTimeComponentLocationZone(datetime,len(pattern))
	case string(datetimeFormatUnitTimeZone2):
		//TODO
	}

	return "", errors.New("unknown datetime format unit: " + pattern[0:1])
}

func isFormatLiteral(char string) bool{
	for _,value := range dateUnits{
		if char == value{
			return false
	}
	}

	return true
}

// parseDateTimeFormat takes a format pattern string and returns a sequence of
// components.
func parseDateTimeFormat(pattern string) ([]*datetimePatternComponent, error) {
	// every thing between single quotes should become a literal
	// all non a-z, A-Z characters become a literal
	// everything else, repeat character sequences become a component
	format := []*datetimePatternComponent{}
	for i := 0; i < len(pattern); {
		char := pattern[i : i+1]

		skip := false
		// for units we don't support yet, just skip over them
		for _, r := range datetimeFormatUnitCutset {
			if char == string(r) {
				skip = true
				break
			}
		}

		if skip {
			i++
			continue
		}

		if isFormatLiteral(char) {

			component := &datetimePatternComponent{
				pattern:       char,
				componentType: datetimePatternComponentLiteral,
			}

			format = append(format, component)
			i++
			continue

		}
		if (char >= "a" && char <= "z") || (char >= "A" && char <= "Z") {
			// this represents a format unit
			// find the entire sequence of the same character
			endChar := lastSequenceIndex(pattern[i:]) + i

			component := &datetimePatternComponent{
				pattern:       pattern[i : endChar+1],
				componentType: datetimePatternComponentUnit,
			}

			format = append(format, component)
			i = endChar + 1
			continue

		}

		component := &datetimePatternComponent{
			pattern:       char,
			componentType: datetimePatternComponentLiteral,
		}

		format = append(format, component)
		i++
		continue

	}

	return format, nil
}

func formatDateTimeComponentYear(datetime time.Time, length int) (string, error) {
	year := datetime.Year()
	switch length {
	case datetimeFormatLength1Plus:
		return formatDateTimeComponentYearLengthWide(year), nil
	case datetimeFormatLength2Plus:
		return formatDateTimeComponentYearLength2Plus(year), nil
	case datetimeFormatLengthWide:
		return formatDateTimeComponentYearLengthWide(year), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported year length: %d", length))
}

// formatDateTimeComponentMonth renders a month component.
func formatDateTimeComponentMonth(datetime time.Time, length int,locale string) (string, error) {

	month := int(datetime.Month())

	switch length {
	case datetimeFormatLength1Plus:
		return formatDateTimeComponentMonth1Plus(month), nil
	case datetimeFormatLength2Plus:
		return formatDateTimeComponentMonth2Plus(month), nil
	case datetimeFormatLengthAbbreviated:
		return formatDateTimeComponentMonthAbbreviated(month,locale), nil
	case datetimeFormatLengthWide:
		return formatDateTimeComponentMonthWide(month,locale), nil
	case datetimeFormatLengthNarrow:
		return formatDateTimeComponentMonthNarrow(month,locale), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported month length: %d", length))
}

// formatDateTimeComponentDayOfWeek renders a day-of-week component.
func formatDateTimeComponentDayOfWeek(datetime time.Time, length int,locale string) (string, error) {
	switch length {
	case datetimeFormatLength1Plus:
		return formatDateTimeComponentDayOfWeekWide(datetime.Weekday(),locale), nil
	case datetimeFormatLength2Plus:
		return formatDateTimeComponentDayOfWeekShort(datetime.Weekday(),locale), nil
	case datetimeFormatLengthAbbreviated:
		return formatDateTimeComponentDayOfWeekAbbreviated(datetime.Weekday(),locale), nil
	case datetimeFormatLengthWide:
		return formatDateTimeComponentDayOfWeekWide(datetime.Weekday(),locale), nil
	case datetimeFormatLengthNarrow:
		return formatDateTimeComponentDayOfWeekNarrow(datetime.Weekday(),locale), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported year day-of-week: %d", length))
}

// formatDateTimeComponentDay renders a day-of-year component.
func formatDateTimeComponentDay(datetime time.Time, length int) (string, error) {
	day := datetime.Day()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", day), nil
	case datetimeFormatLength2Plus:
		if day < 10 {
			return fmt.Sprintf("0%d", day), nil
		}
		return fmt.Sprintf("%d", day), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported day-of-year: %d", length))
}

// formatDateTimeComponentHour12 renders an hour-component using a 12-hour
// clock.
func formatDateTimeComponentHour12(datetime time.Time, length int) (string, error) {
	hour := datetime.Hour()
	if hour > 12 {
		hour = hour - 12
	}

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", hour), nil
	case datetimeFormatLength2Plus:
		if hour < 10 {
			return fmt.Sprintf("0%d", hour), nil
		}
		return fmt.Sprintf("%d", hour), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported hour-12: %d", length))
}

// formatDateTimeComponentHour24 renders an hour-component using a 24-hour
// clock.
func formatDateTimeComponentHour24(datetime time.Time, length int) (string, error) {
	hour := datetime.Hour()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", hour), nil
	case datetimeFormatLength2Plus:
		if hour < 10 {
			return fmt.Sprintf("0%d", hour), nil
		}
		return fmt.Sprintf("%d", hour), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported hour-24: %d", length))
}

// formatDateTimeComponentMinute renders a minute component.
func formatDateTimeComponentMinute(datetime time.Time, length int) (string, error) {
	minute := datetime.Minute()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", minute), nil
	case datetimeFormatLength2Plus:
		if minute < 10 {
			return fmt.Sprintf("0%d", minute), nil
		}
		return fmt.Sprintf("%d", minute), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported minute: %d", length))
}

// formatDateTimeComponentSecond renders a second component
func formatDateTimeComponentSecond(datetime time.Time, length int) (string, error) {
	second := datetime.Second()

	switch length {
	case datetimeFormatLength1Plus:
		return fmt.Sprintf("%d", second), nil
	case datetimeFormatLength2Plus:
		if second < 10 {
			return fmt.Sprintf("0%d", second), nil
		}
		return fmt.Sprintf("%d", second), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported second: %d", length))
}

// formatDateTimeComponentPeriod renders a period component (AM/PM).
func formatDateTimeComponentPeriod(datetime time.Time, length int,locale string) (string, error) {
	hour := datetime.Hour()

	switch length {
	case datetimeFormatLength1Plus:
		return formatDateTimeComponentPeriodWide(hour,locale), nil
	case datetimeFormatLengthAbbreviated:
		return formatDateTimeComponentPeriodAbbreviated(hour,locale), nil
	case datetimeFormatLengthWide:
		return formatDateTimeComponentPeriodWide(hour,locale), nil
	case datetimeFormatLengthNarrow:
		return formatDateTimeComponentPeriodNarrow(hour,locale), nil
	}

	return "", errors.New(fmt.Sprintf("unsupported day-period: %d", length))
}

// formatDateTimeComponentYearLength2Plus renders a full-year component - for
// all modern dates, that's four digits.
func formatDateTimeComponentYearLengthWide(year int) string {
	return fmt.Sprintf("%d", year)
}

// formatDateTimeComponentYearLength2Plus renders a 2-digit year component.
func formatDateTimeComponentYearLength2Plus(year int) string {
	yearShort := year % 100

	if yearShort < 10 {
		return fmt.Sprintf("0%d", yearShort)
	}

	return fmt.Sprintf("%d", yearShort)
}

// formatDateTimeComponentMonth1Plus renders a numeric month component with 1 or
// 2 digits depending on value.
func formatDateTimeComponentMonth1Plus(month int) string {
	return fmt.Sprintf("%d", month)
}

// formatDateTimeComponentMonth2Plus renders a numeric month component always
// with 2 digits.
func formatDateTimeComponentMonth2Plus(month int) string {
	if month < 10 {
		return fmt.Sprintf("0%d", month)
	}
	return fmt.Sprintf("%d", month)
}

// formatDateTimeComponentMonthAbbreviated renders an abbreviated text month
// component.
func formatDateTimeComponentMonthAbbreviated(month int,locale string) string {
	switch month {
	case 1:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[0]
	case 2:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[1]
	case 3:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[2]
	case 4:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[3]
	case 5:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[4]
	case 6:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[5]
	case 7:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[6]
	case 8:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[7]
	case 9:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[8]
	case 10:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[9]
	case 11:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[10]
	case 12:
		return cacheFormatMap[locale].Messages.MonthsFormat.Abbreviated[11]
	}

	return ""
}

// formatDateTimeComponentMonthWide renders a full text month component.
func formatDateTimeComponentMonthWide(month int,locale string) string {
	switch month {
	case 1:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[0]
	case 2:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[1]
	case 3:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[2]
	case 4:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[3]
	case 5:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[4]
	case 6:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[5]
	case 7:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[6]
	case 8:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[7]
	case 9:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[8]
	case 10:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[9]
	case 11:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[10]
	case 12:
		return cacheFormatMap[locale].Messages.MonthsFormat.Wide[11]
	}

	return ""
}

// formatDateTimeComponentMonthNarrow renders a super-short month compontent -
// not guaranteed to be unique for different months.
func formatDateTimeComponentMonthNarrow(month int,locale string) string {
	switch month {
	case 1:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[0]
	case 2:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[1]
	case 3:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[2]
	case 4:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[3]
	case 5:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[4]
	case 6:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[5]
	case 7:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[6]
	case 8:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[7]
	case 9:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[8]
	case 10:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[9]
	case 11:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[10]
	case 12:
		return cacheFormatMap[locale].Messages.MonthsFormat.Narrow[11]
	}

	return ""
}

// formatDateTimeComponentDayOfWeekWide renders a full text day-of-week
// component.
func formatDateTimeComponentDayOfWeekWide(dayOfWeek time.Weekday,locale string) string {
	switch dayOfWeek {
	case time.Sunday:
		return cacheFormatMap[locale].Messages.DaysFormat.Wide[0]
	case time.Monday:
		return cacheFormatMap[locale].Messages.DaysFormat.Wide[1]
	case time.Tuesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Wide[2]
	case time.Wednesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Wide[3]
	case time.Thursday:
		return cacheFormatMap[locale].Messages.DaysFormat.Wide[4]
	case time.Friday:
		return cacheFormatMap[locale].Messages.DaysFormat.Wide[5]
	case time.Saturday:
		return cacheFormatMap[locale].Messages.DaysFormat.Wide[6]
	}

	return ""
}

// formatDateTimeComponentDayOfWeekNarrow renders a super-short day-of-week
// compontent - not guaranteed to be unique for different days.
func formatDateTimeComponentDayOfWeekNarrow(dayOfWeek time.Weekday,locale string) string {
	switch dayOfWeek {
	case time.Sunday:
		return cacheFormatMap[locale].Messages.DaysFormat.Narrow[0]
	case time.Monday:
		return cacheFormatMap[locale].Messages.DaysFormat.Narrow[1]
	case time.Tuesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Narrow[2]
	case time.Wednesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Narrow[3]
	case time.Thursday:
		return cacheFormatMap[locale].Messages.DaysFormat.Narrow[4]
	case time.Friday:
		return cacheFormatMap[locale].Messages.DaysFormat.Narrow[5]
	case time.Saturday:
		return cacheFormatMap[locale].Messages.DaysFormat.Narrow[6]
	}

	return ""
}

// formatDateTimeComponentDayOfWeekAbbreviated renders an abbreviated text
// day-of-week component.
func formatDateTimeComponentDayOfWeekAbbreviated(dayOfWeek time.Weekday,locale string) string {
	switch dayOfWeek {
	case time.Sunday:
		return cacheFormatMap[locale].Messages.DaysFormat.Abbreviated.Sun
	case time.Monday:
		return cacheFormatMap[locale].Messages.DaysFormat.Abbreviated.Mon
	case time.Tuesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Abbreviated.Tue
	case time.Wednesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Abbreviated.Wed
	case time.Thursday:
		return cacheFormatMap[locale].Messages.DaysFormat.Abbreviated.Thu
	case time.Friday:
		return cacheFormatMap[locale].Messages.DaysFormat.Abbreviated.Fri
	case time.Saturday:
		return cacheFormatMap[locale].Messages.DaysFormat.Abbreviated.Sat
	}

	return ""
}

// formatDateTimeComponentDayOfWeekAbbreviated renders a
// shorter-then-abbreviated but still unique text day-of-week component.
func formatDateTimeComponentDayOfWeekShort(dayOfWeek time.Weekday,locale string) string {
	switch dayOfWeek {
	case time.Sunday:
		return cacheFormatMap[locale].Messages.DaysFormat.Short[0]
	case time.Monday:
		return cacheFormatMap[locale].Messages.DaysFormat.Short[1]
	case time.Tuesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Short[2]
	case time.Wednesday:
		return cacheFormatMap[locale].Messages.DaysFormat.Short[3]
	case time.Thursday:
		return cacheFormatMap[locale].Messages.DaysFormat.Short[4]
	case time.Friday:
		return cacheFormatMap[locale].Messages.DaysFormat.Short[5]
	case time.Saturday:
		return cacheFormatMap[locale].Messages.DaysFormat.Short[6]
	}

	return ""
}

// formatDateTimeComponentPeriodAbbreviated renders an abbreviated period
// component.
func formatDateTimeComponentPeriodAbbreviated(hour int,locale string) string {
	if hour < 12 {
		return cacheFormatMap[locale].Messages.DayPeriodsFormat.Abbreviated[0]
	}

	return cacheFormatMap[locale].Messages.DayPeriodsFormat.Abbreviated[1]
}

// formatDateTimeComponentPeriodWide renders a full period component.
func formatDateTimeComponentPeriodWide(hour int,locale string) string {
	if hour < 12 {
		return cacheFormatMap[locale].Messages.DayPeriodsFormat.Wide[0]
	}

	return cacheFormatMap[locale].Messages.DayPeriodsFormat.Wide[1]
}

// formatDateTimeComponentPeriodNarrow renders a super-short period component.
func formatDateTimeComponentPeriodNarrow(hour int,locale string) string {
	if hour < 12 {
		return cacheFormatMap[locale].Messages.DayPeriodsFormat.Narrow[0]
	}

	return cacheFormatMap[locale].Messages.DayPeriodsFormat.Narrow[1]
}

func formatDateTimeComponentLocationZone(datetime time.Time, length int) (string, error){
	var b []byte
	name,offset := datetime.Zone()

	zone := offset / 60 // convert to minutes
	absoffset := offset
	if zone < 0 {
		b = append(b, '-')
		zone = -zone
		absoffset = -absoffset
	} else {
		b = append(b, '+')
	}
	b = appendInt(b, zone/60, 2)

	switch length {
	case datetimeFormatLength1Plus:
		b = append(b, ':')
		b = appendInt(b, zone%60, 2)
		b = appendInt(b, absoffset%60, 2)
		return string(b),nil
	case datetimeFormatLengthWide:
		b = append(b, ':')
		b = appendInt(b, zone%60, 2)
		return name + string(b),nil
	}

	return "",errors.New(fmt.Sprintf("unsupported zone-period: %d", length))
}

// lastSequenceIndex looks at the first character in a string and returns the
// last digits of the first sequence of that character. For example:
//  - ABC: 0
//  - AAB: 1
//  - ABA: 0
//  - AAA: 2
func lastSequenceIndex(str string) int {
	if len(str) == 0 {
		return -1
	}

	if len(str) == 1 {
		return 0
	}

	sequenceChar := str[0:1]
	lastPos := 0
	for i := 1; i < len(str); i++ {
		if str[i:i+1] != sequenceChar {
			break
		}

		lastPos = i
	}

	return lastPos
}

// appendInt appends the decimal form of x to b and returns the result.
// If the decimal form (excluding sign) is shorter than width, the result is padded with leading 0's.
// Duplicates functionality in strconv, but avoids dependency.
func appendInt(b []byte, x int, width int) []byte {
	u := uint(x)
	if x < 0 {
		b = append(b, '-')
		u = uint(-x)
	}

	// Assemble decimal in reverse order.
	var buf [20]byte
	i := len(buf)
	for u >= 10 {
		i--
		q := u / 10
		buf[i] = byte('0' + u - q*10)
		u = q
	}
	i--
	buf[i] = byte('0' + u)

	// Add 0-padding.
	for w := len(buf) - i; w < width; w++ {
		b = append(b, '0')
	}

	return append(b, buf[i:]...)
}




