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
	dateResult = map[int]map[string]string{
		constants.SHORTDATE : {"zh_CN":"2006/1/2","de":"02.01.06","es":"2/1/06","fr":"02/01/2006","ja":"2006/01/02","ko":"06. 1. 2.","zh_TW":"2006/1/2",},
		constants.MEDIUMDATE : {"zh_CN":"2006年1月2日","de":"02.01.2006","es":"2 ene. 2006","fr":"2 janv. 2006","ja":"2006/01/02","ko":"2006. 1. 2.","zh_TW":"2006年1月2日",},
		constants.FULLDATE : {"zh_CN":"2006年1月2日星期一","de":"Montag, 2. Januar 2006","es":"lunes, 2 '2e' enero '2e' 2006","fr":"lundi 2 janvier 2006","ja":"2006年1月2日月曜日","ko":"2006년 1월 2일 월요일","zh_TW":"2006年1月2日 星期一",},
		constants.LONGDATE : {"zh_CN":"2006年1月2日","de":"2. Januar 2006","es":"2 '2e' enero '2e' 2006","fr":"2 janvier 2006","ja":"2006年1月2日","ko":"2006년 1월 2일","zh_TW":"2006年1月2日",},
		constants.SHORTTIME : {"zh_CN":"下午3:04","de":"15:04","es":"15:04","fr":"15:04","ja":"15:04","ko":"오후 3:04","zh_TW":"下午3:04",},
		constants.MEDIUMTIME : {"zh_CN":"下午3:04:05","de":"15:04:05","es":"15:04:05","fr":"15:04:05","ja":"15:04:05","ko":"오후 3:04:05","zh_TW":"下午3:04:05",},
		constants.LONGTIME : {"zh_CN":"+08:0000 下午3:04:05","de":"15:04:05 +08:0000","es":"15:04:05 +08:0000","fr":"15:04:05 +08:0000","ja":"15:04:05 +08:0000","ko":"오후 3시 4분 5초 +08:0000","zh_TW":"下午3:04:05 [+08:0000]",},
		constants.FULLTIME : {"zh_CN":"CST+08:00 下午3:04:05","de":"15:04:05 CST+08:00","es":"15:04:05 (CST+08:00)","fr":"15:04:05 CST+08:00","ja":"15時04分05秒 CST+08:00","ko":"오후 3시 4분 5초 CST+08:00","zh_TW":"下午3:04:05 [CST+08:00]",},
		constants.SHORT : {"zh_CN":"下午3:04 2006/1/2","de":"15:04, 02.01.06","es":"15:04 2/1/06","fr":"15:04 02/01/2006","ja":"15:04 2006/01/02","ko":"오후 3:04 06. 1. 2.","zh_TW":"下午3:04 2006/1/2",},
		constants.MEDIUM : {"zh_CN":"下午3:04:05 2006年1月2日","de":"15:04:05, 02.01.2006","es":"15:04:05 2 ene. 2006","fr":"15:04:05 'à' 2 janv. 2006","ja":"15:04:05 2006/01/02","ko":"오후 3:04:05 2006. 1. 2.","zh_TW":"下午3:04:05 2006年1月2日",},
		constants.LONG : {"zh_CN":"+08:0000 下午3:04:05 2006年1月2日","de":"15:04:05 +08:0000 'u4' 2. Januar 2006","es":"15:04:05 +08:0000, 2 '2e' enero '2e' 2006","fr":"15:04:05 +08:0000 'à' 2 janvier 2006","ja":"15:04:05 +08:0000 2006年1月2日","ko":"오후 3시 4분 5초 +08:0000 2006년 1월 2일","zh_TW":"下午3:04:05 [+08:0000] 2006年1月2日",},
		constants.FULL : {"zh_CN":"CST+08:00 下午3:04:05 2006年1月2日星期一","de":"15:04:05 CST+08:00 'u4' Montag, 2. Januar 2006","es":"15:04:05 (CST+08:00), lunes, 2 '2e' enero '2e' 2006","fr":"15:04:05 CST+08:00 'à' lundi 2 janvier 2006","ja":"15時04分05秒 CST+08:00 2006年1月2日月曜日","ko":"오후 3시 4분 5초 CST+08:00 2006년 1월 2일 월요일","zh_TW":"下午3:04:05 [CST+08:00] 2006年1月2日 星期一",},
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

			exceptedValue := dateResult[number][local]
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

			exceptedValue := dateResult[constants.FULL][local]
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

			exceptedValue := dateResult[constants.FULLDATE][local]
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