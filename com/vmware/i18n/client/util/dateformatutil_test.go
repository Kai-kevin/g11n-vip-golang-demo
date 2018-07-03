package util

import (
	"testing"
	"time"
	"vipgoclient/com/vmware/i18n/client/constants"
)

var(
	dateTimeString = "Jan 2, 2006 at 3:04:05pm"
	initialLocation = "Asia/Shanghai"
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
	cstSh, _ = time.LoadLocation(initialLocation)
	dateTime,_ = time.ParseInLocation(dateTimeString,dateTimeString,cstSh)
)

func TestFormatDateTime(t *testing.T) {
	t.Log("Get the need to test DateTimeFormat")
	{
		for _,local := range Locales{
			for _,dateTimeConstant := range dateTimeConstants{
				dFull,err := FormatDateTime(dateTimeConstant,dateTime,local)

				if err != nil {
					t.Fatal("TestFormatDateTime failed!!!")
				}

				exceptedValue := dateResult[dateTimeConstant][local]

				if exceptedValue != dFull {
					t.Fatalf("The result of the getNumber is not the excepted value in local:\"%s\" and pattern:\"%v\"!!!",local,dateTimeConstant)
				}

				t.Logf("local:\"%s\";pattern:\"%v\";oldTime:\"%s\" convert to newTime:\"%s\"",local,dateTimeConstant,dateTime,dFull)
			}
		}
	}

}

func BenchmarkFormatDateTime(b *testing.B) {
	success,fail := 0,0
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_,err := FormatDateTime(ConstantNumber,dateTime,Locales[LocaleNumber])
		if err != nil {
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormatDateTime",success,fail)
}