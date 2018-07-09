package util

import (
	"testing"
	"time"
	"vipgoclient/com/vmware/i18n/client/constants"
)

var (
	dateTimeString    = "Jan 2, 2006 at 3:04:05pm"
	initialLocation   = "Asia/Shanghai"
	dateTimeConstants = []int{
		constants.SHORTDATE,
		constants.MEDIUMDATE,
		constants.FULLDATE,
		constants.LONGDATE,
		constants.SHORTTIME,
		constants.MEDIUMTIME,
		constants.LONGTIME,
		constants.FULLTIME,
		constants.SHORT,
		constants.MEDIUM,
		constants.LONG,
		constants.FULL,
	}
	dateResult = map[string]map[int]string{
		"zh_CN": {constants.SHORTDATE: "2006/1/2", constants.MEDIUMDATE: "2006年1月2日", constants.FULLDATE: "2006年1月2日星期一", constants.LONGDATE: "2006年1月2日", constants.SHORTTIME: "下午3:04", constants.MEDIUMTIME: "下午3:04:05", constants.LONGTIME: "+08:0000 下午3:04:05", constants.FULLTIME: "CST+08:00 下午3:04:05", constants.SHORT: "下午3:04 2006/1/2", constants.MEDIUM: "下午3:04:05 2006年1月2日", constants.LONG: "+08:0000 下午3:04:05 2006年1月2日", constants.FULL: "CST+08:00 下午3:04:05 2006年1月2日星期一"},
		"de":    {constants.SHORTDATE: "02.01.06", constants.MEDIUMDATE: "02.01.2006", constants.FULLDATE: "Montag, 2. Januar 2006", constants.LONGDATE: "2. Januar 2006", constants.SHORTTIME: "15:04", constants.MEDIUMTIME: "15:04:05", constants.LONGTIME: "15:04:05 +08:0000", constants.FULLTIME: "15:04:05 CST+08:00", constants.SHORT: "15:04, 02.01.06", constants.MEDIUM: "15:04:05, 02.01.2006", constants.LONG: "15:04:05 +08:0000 'u4' 2. Januar 2006", constants.FULL: "15:04:05 CST+08:00 'u4' Montag, 2. Januar 2006"},
		"es":    {constants.SHORTDATE: "2/1/06", constants.MEDIUMDATE: "2 ene. 2006", constants.FULLDATE: "lunes, 2 '2e' enero '2e' 2006", constants.LONGDATE: "2 '2e' enero '2e' 2006", constants.SHORTTIME: "15:04", constants.MEDIUMTIME: "15:04:05", constants.LONGTIME: "15:04:05 +08:0000", constants.FULLTIME: "15:04:05 (CST+08:00)", constants.SHORT: "15:04 2/1/06", constants.MEDIUM: "15:04:05 2 ene. 2006", constants.LONG: "15:04:05 +08:0000, 2 '2e' enero '2e' 2006", constants.FULL: "15:04:05 (CST+08:00), lunes, 2 '2e' enero '2e' 2006"},
		"fr":    {constants.SHORTDATE: "02/01/2006", constants.MEDIUMDATE: "2 janv. 2006", constants.FULLDATE: "lundi 2 janvier 2006", constants.LONGDATE: "2 janvier 2006", constants.SHORTTIME: "15:04", constants.MEDIUMTIME: "15:04:05", constants.LONGTIME: "15:04:05 +08:0000", constants.FULLTIME: "15:04:05 CST+08:00", constants.SHORT: "15:04 02/01/2006", constants.MEDIUM: "15:04:05 'à' 2 janv. 2006", constants.LONG: "15:04:05 +08:0000 'à' 2 janvier 2006", constants.FULL: "15:04:05 CST+08:00 'à' lundi 2 janvier 2006"},
		"ja":    {constants.SHORTDATE: "2006/01/02", constants.MEDIUMDATE: "2006/01/02", constants.FULLDATE: "2006年1月2日月曜日", constants.LONGDATE: "2006年1月2日", constants.SHORTTIME: "15:04", constants.MEDIUMTIME: "15:04:05", constants.LONGTIME: "15:04:05 +08:0000", constants.FULLTIME: "15時04分05秒 CST+08:00", constants.SHORT: "15:04 2006/01/02", constants.MEDIUM: "15:04:05 2006/01/02", constants.LONG: "15:04:05 +08:0000 2006年1月2日", constants.FULL: "15時04分05秒 CST+08:00 2006年1月2日月曜日"},
		"ko":    {constants.SHORTDATE: "06. 1. 2.", constants.MEDIUMDATE: "2006. 1. 2.", constants.FULLDATE: "2006년 1월 2일 월요일", constants.LONGDATE: "2006년 1월 2일", constants.SHORTTIME: "오후 3:04", constants.MEDIUMTIME: "오후 3:04:05", constants.LONGTIME: "오후 3시 4분 5초 +08:0000", constants.FULLTIME: "오후 3시 4분 5초 CST+08:00", constants.SHORT: "오후 3:04 06. 1. 2.", constants.MEDIUM: "오후 3:04:05 2006. 1. 2.", constants.LONG: "오후 3시 4분 5초 +08:0000 2006년 1월 2일", constants.FULL: "오후 3시 4분 5초 CST+08:00 2006년 1월 2일 월요일"},
		"zh_TW": {constants.SHORTDATE: "2006/1/2", constants.MEDIUMDATE: "2006年1月2日", constants.FULLDATE: "2006年1月2日 星期一", constants.LONGDATE: "2006年1月2日", constants.SHORTTIME: "下午3:04", constants.MEDIUMTIME: "下午3:04:05", constants.LONGTIME: "下午3:04:05 [+08:0000]", constants.FULLTIME: "下午3:04:05 [CST+08:00]", constants.SHORT: "下午3:04 2006/1/2", constants.MEDIUM: "下午3:04:05 2006年1月2日", constants.LONG: "下午3:04:05 [+08:0000] 2006年1月2日", constants.FULL: "下午3:04:05 [CST+08:00] 2006年1月2日 星期一"},
	}
	cstSh, _    = time.LoadLocation(initialLocation)
	dateTime, _ = time.ParseInLocation(dateTimeString, dateTimeString, cstSh)
)

func TestFormatDateTime(t *testing.T) {
	t.Log("Get the need to test DateTimeFormat")
	{
		for _, local := range Locales {
			for _, dateTimeConstant := range dateTimeConstants {
				dFull, err := FormatDateTime(dateTimeConstant, dateTime, local)
				if err != nil {
					t.Fatal("TestFormatDateTime failed!!!")
				}

				exceptedValue := dateResult[local][dateTimeConstant]
				if exceptedValue != dFull {
					t.Fatalf("The result of the getNumber is not the excepted value in local:\"%s\" and pattern:\"%v\"!!!", local, dateTimeConstant)
				}

				t.Logf("local:\"%s\";pattern:\"%v\";oldTime:\"%s\" convert to newTime:\"%s\"", local, dateTimeConstant, dateTime, dFull)
			}
		}
	}
}

func BenchmarkFormatDateTime(b *testing.B) {
	success, fail := 0, 0
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := FormatDateTime(ConstantNumber, dateTime, Locales[LocaleNumber])
		if err != nil {
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t", "BenchmarkFormatDateTime", success, fail)
}
