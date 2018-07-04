package sample

import (
	"strings"
	"vipgoclient/com/vmware/i18n/client/conf"
	"testing"
	"vipgoclient/com/vmware/i18n/client/constants"
	"time"
	"math/rand"
)

var (
	Locales = strings.Split(conf.GetVipConfigInstance().Locales, ",")
	number = -1234.5679
	dateTimeString = "Jan 2, 2006 at 3:04:05pm"
	initialLocation = "Asia/Shanghai"
	cstSh, _ = time.LoadLocation(initialLocation)
	dateTime,_ = time.ParseInLocation(dateTimeString,dateTimeString,cstSh)
	dateResult = map[string]map[int]string{
		"zh_CN" : {constants.SHORTDATE:"2006/1/2",constants.MEDIUMDATE:"2006年1月2日",constants.FULLDATE:"2006年1月2日星期一",constants.LONGDATE:"2006年1月2日",constants.SHORTTIME:"下午3:04",constants.MEDIUMTIME:"下午3:04:05",constants.LONGTIME:"+08:0000 下午3:04:05",constants.FULLTIME:"CST+08:00 下午3:04:05",constants.SHORT:"下午3:04 2006/1/2", constants.MEDIUM:"下午3:04:05 2006年1月2日",constants.LONG:"+08:0000 下午3:04:05 2006年1月2日",constants.FULL:"CST+08:00 下午3:04:05 2006年1月2日星期一"},
		"de" : {constants.SHORTDATE:"02.01.06",constants.MEDIUMDATE:"02.01.2006",constants.FULLDATE:"Montag, 2. Januar 2006",constants.LONGDATE:"2. Januar 2006",constants.SHORTTIME:"15:04",constants.MEDIUMTIME:"15:04:05",constants.LONGTIME:"15:04:05 +08:0000",constants.FULLTIME:"15:04:05 CST+08:00",constants.SHORT:"15:04, 02.01.06", constants.MEDIUM:"15:04:05, 02.01.2006",constants.LONG:"15:04:05 +08:0000 'u4' 2. Januar 2006",constants.FULL:"15:04:05 CST+08:00 'u4' Montag, 2. Januar 2006"},
		"es" : {constants.SHORTDATE:"2/1/06",constants.MEDIUMDATE:"2 ene. 2006",constants.FULLDATE:"lunes, 2 '2e' enero '2e' 2006",constants.LONGDATE:"2 '2e' enero '2e' 2006",constants.SHORTTIME:"15:04",constants.MEDIUMTIME:"15:04:05",constants.LONGTIME:"15:04:05 +08:0000",constants.FULLTIME:"15:04:05 (CST+08:00)",constants.SHORT:"15:04 2/1/06", constants.MEDIUM:"15:04:05 2 ene. 2006",constants.LONG:"15:04:05 +08:0000, 2 '2e' enero '2e' 2006",constants.FULL:"15:04:05 (CST+08:00), lunes, 2 '2e' enero '2e' 2006"},
		"fr" : {constants.SHORTDATE:"02/01/2006",constants.MEDIUMDATE:"2 janv. 2006",constants.FULLDATE:"lundi 2 janvier 2006",constants.LONGDATE:"2 janvier 2006",constants.SHORTTIME:"15:04",constants.MEDIUMTIME:"15:04:05",constants.LONGTIME:"15:04:05 +08:0000",constants.FULLTIME:"15:04:05 CST+08:00",constants.SHORT:"15:04 02/01/2006", constants.MEDIUM:"15:04:05 'à' 2 janv. 2006",constants.LONG:"15:04:05 +08:0000 'à' 2 janvier 2006",constants.FULL:"15:04:05 CST+08:00 'à' lundi 2 janvier 2006"},
		"ja" : {constants.SHORTDATE:"2006/01/02",constants.MEDIUMDATE:"2006/01/02",constants.FULLDATE:"2006年1月2日月曜日",constants.LONGDATE:"2006年1月2日",constants.SHORTTIME:"15:04",constants.MEDIUMTIME:"15:04:05",constants.LONGTIME:"15:04:05 +08:0000",constants.FULLTIME:"15時04分05秒 CST+08:00",constants.SHORT:"15:04 2006/01/02", constants.MEDIUM:"15:04:05 2006/01/02",constants.LONG:"15:04:05 +08:0000 2006年1月2日",constants.FULL:"15時04分05秒 CST+08:00 2006年1月2日月曜日"},
		"ko" : {constants.SHORTDATE:"06. 1. 2.",constants.MEDIUMDATE:"2006. 1. 2.",constants.FULLDATE:"2006년 1월 2일 월요일",constants.LONGDATE:"2006년 1월 2일",constants.SHORTTIME:"오후 3:04",constants.MEDIUMTIME:"오후 3:04:05",constants.LONGTIME:"오후 3시 4분 5초 +08:0000",constants.FULLTIME:"오후 3시 4분 5초 CST+08:00",constants.SHORT:"오후 3:04 06. 1. 2.", constants.MEDIUM:"오후 3:04:05 2006. 1. 2.",constants.LONG:"오후 3시 4분 5초 +08:0000 2006년 1월 2일",constants.FULL:"오후 3시 4분 5초 CST+08:00 2006년 1월 2일 월요일"},
		"zh_TW" : {constants.SHORTDATE:"2006/1/2",constants.MEDIUMDATE:"2006年1月2日",constants.FULLDATE:"2006年1月2日 星期一",constants.LONGDATE:"2006年1月2日",constants.SHORTTIME:"下午3:04",constants.MEDIUMTIME:"下午3:04:05",constants.LONGTIME:"下午3:04:05 [+08:0000]",constants.FULLTIME:"下午3:04:05 [CST+08:00]",constants.SHORT:"下午3:04 2006/1/2", constants.MEDIUM:"下午3:04:05 2006年1月2日",constants.LONG:"下午3:04:05 [+08:0000] 2006年1月2日",constants.FULL:"下午3:04:05 [CST+08:00] 2006年1月2日 星期一"},
	}
	exceptedNumberValue = map[string]string{
		"zh_CN":"-1,234.568",
		"de":"-1.234,568",
		"es":"-1.234,568",
		"fr":"-1 234,568",
		"ja":"-1,234.568",
		"ko":"-1,234.568",
		"zh_TW":"-1,234.568",
	}
	exceptedPercentValue = map[string]string{
		"zh_CN":"-123,456.79",
		"de":"-123.456,79",
		"es":"-123.456,79",
		"fr":"-123 456,79",
		"ja":"-123,456.79",
		"ko":"-123,456.79",
		"zh_TW":"-123,456.79",
	}
	exceptedCurrencyValue = map[string]string{
		"zh_CN":"￥-1,234.57",
		"de":"-1.234,57 €",
		"es":"-1.234,57 €",
		"fr":"-1 234,57 €",
		"ja":"￥-1,234.57",
		"ko":"₩-1,234.57",
		"zh_TW":"CN¥-1,234.57",
	}
	count int
)

func init()  {
	rand.Seed(time.Now().Unix())
	count = rand.Intn(len(Locales))
}

func TestGetNumber(t *testing.T) {
	t.Log("Get the need to test getNumber")
	{
		for _,local := range Locales{
			format := Formator{local}
			result,error := format.GetNumber(number)
			if error != nil {
				t.Log("GetNumberFormatByLocal failed!!!")
			}

			_,exist := exceptedNumberValue[local]
			if exist && exceptedNumberValue[local] != result {
				t.Fatalf("The result of the getNumber is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("oldNumber:\"%v\" convert to newNumber:\"%s\" in locale:\"%s\"",number,result,local)
		}
	}
}

func TestGetPercentNumber(t *testing.T) {
	t.Log("Get the need to test getPercentNumber")
	{
		for _,local := range Locales{
			format := Formator{local}
			result,error := format.GetPercentNumber(number)
			if error != nil {
				t.Log("GetPercentNumber failed!!!")
			}

			_,exist := exceptedPercentValue[local]
			if exist && exceptedPercentValue[local] != result {
				t.Fatalf("The result of the getPercentNumber is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("oldNumber:\"%v\" convert to newNumber:\"%s\" in locale:\"%s\"",number,result,local)
		}
	}
}

func TestGetCurrencyNumber(t *testing.T) {
	t.Log("Get the need to test getCurrencyNumber")
	{
		for _,local := range Locales{
			format := Formator{local}
			result,error := format.GetCurrencyNumber(number)
			if error != nil {
				t.Log("GetCurrencyNumber failed!!!")
			}

			_,exist := exceptedCurrencyValue[local]
			if exist && exceptedCurrencyValue[local] != result {
				t.Fatalf("The result of the getCurrencyNumber is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("oldNumber:\"%v\" convert to newNumber:\"%s\" in locale:\"%s\"",number,result,local)
		}
	}
}

func TestGetDate(t *testing.T) {
	t.Log("Get the need to test getDate")
	{
		for _,local := range Locales{
			format := Formator{local}
			rand.Seed(time.Now().Unix())
			number := rand.Intn(constants.FULL)
			result,error := format.GetDate(dateTime,number)

			if error != nil {
				t.Log("GetDate failed!!!")
			}

			exceptedValue := dateResult[local][number]
			if exceptedValue != result {
				t.Fatalf("The result of the getDate is not the excepted value in local:\"%s\" and pattern:\"%v\"!!!",local,number)
			}

			t.Logf("oldTime:\"%v\" convert to newTime:\"%s\" in locale:\"%s\"",dateTime,result,local)
		}
	}
}

func TestGetFullDateTime(t *testing.T) {
	t.Log("Get the need to test getFullDateTime")
	{
		for _,local := range Locales{
			format := Formator{local}
			result,error := format.GetFullDateTime(dateTime)

			if error != nil {
				t.Log("GetFullDateTime failed!!!")
			}

			exceptedValue := dateResult[local][constants.FULL]
			if exceptedValue != result {
				t.Fatalf("The result of the getFullDateTime is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("oldTime:\"%v\" convert to newTime:\"%s\" in locale:\"%s\"",dateTime,result,local)
		}
	}
}

func TestGetFullDate(t *testing.T) {
	t.Log("Get the need to test getFullDate")
	{
		for _,local := range Locales{
			format := Formator{local}
			result,error := format.GetFullDate(dateTime)

			if error != nil {
				t.Log("GetFullDateTime failed!!!")
			}

			exceptedValue := dateResult[local][constants.FULLDATE]
			if exceptedValue != result {
				t.Fatalf("The result of the getFullDate is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("oldTime:\"%v\" convert to newTime:\"%s\" in locale:\"%s\"",dateTime,result,local,)
		}
	}
}

func BenchmarkFormator_GetNumber(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format := Formator{Locales[count]}
		_,error := format.GetNumber(number)
		if error != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormator_GetNumber",success,fail)
}

func BenchmarkFormator_GetPercentNumber(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format := Formator{Locales[count]}
		_,error := format.GetPercentNumber(number)
		if error != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormator_GetPercentNumber",success,fail)
}

func BenchmarkFormator_GetCurrencyNumber(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format := Formator{Locales[count]}
		_,error := format.GetCurrencyNumber(number)
		if error != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormator_GetCurrencyNumber",success,fail)
}

func BenchmarkFormator_GetDate(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format := Formator{Locales[count]}
		_,error := format.GetDate(dateTime,constants.SHORTTIME)
		if error != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormator_GetDate",success,fail)
}

func BenchmarkFormator_GetFullDateTime(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format := Formator{Locales[count]}
		_,error := format.GetFullDateTime(dateTime)
		if error != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormator_GetFullDateTime",success,fail)
}

func BenchmarkFormator_GetFullDate(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format := Formator{Locales[count]}
		_,error := format.GetFullDate(dateTime)
		if error != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormator_GetFullDate",success,fail)
}