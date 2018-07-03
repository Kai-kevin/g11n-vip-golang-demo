package dao

import (
	"testing"
	"vipgoclient/com/vmware/i18n/client/conf"
	"strings"
)

var locales = strings.Split(conf.GetVipConfigInstance().Locales, ",")
var components = strings.Split(conf.GetVipConfigInstance().Components, ",")

func TestGetFormattingPatternsByLocal(t *testing.T) {
	t.Log("Given the need to test GetFormattingPatternsByLocal")
	{
		for _,local := range locales {
			resp := GetFormattingPatternsByLocal(local)

			if resp.Response.Code != 200{
				t.Fatal("GetFormattingPatternsByLocal failed!!!")
			}

			t.Log(resp.Data.Messages)
		}
	}

	/*resp := GetFormattingPatternsByLocal("fr")

	if resp.Response.Code != 200{
		t.Fatal("GetFormattingPatternsByLocal failed!!!")
	}

	t.Log(resp.Data.Messages)*/

	//fmt.Println(time.Now().Format("06-01-02 15:04:05"))
	//
	//now := time.Now()
	//local1, err1 := time.LoadLocation("") //equals "UTC"
	//
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//local2, err2 := time.LoadLocation("Local")//set location
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//
	//fmt.Println(local1.String())
	//fmt.Println(time.Local.String())
	//
	//local3, err3 := time.LoadLocation("America/Los_Angeles")
	//if err3 != nil {
	//	fmt.Println(err3)
	//}
	//
	//fmt.Println(local3.String())
	//
	//fmt.Println(now.In(local1))
	//fmt.Println(now.In(local2))
	//fmt.Println(now.In(local3))

}

/*func TestGetFormatMap(t *testing.T){
	cacheFormatMap := *util.GetFormatMap()

	if cacheFormatMap["fr"].Messages.DateTimeFormat.Full == ""{
		t.Fatal("cacheFormat failed!!!")
	}

	t.Log(cacheFormatMap["fr"])
}*/

func TestGetTranslationByComponent(t *testing.T)  {
	t.Log("Given the need to test TestGetTranslationByComponent")
	{
		for _,local := range locales {
			for _,component := range components {
				resp := GetTranslationByComponent(local,component)

				if resp.Response.Code != 200{
					t.Fatal("TestGetTranslationByComponent failed!!!")
				}

				t.Log(resp.Data.Messages)
			}
		}
	}
}

func TestGetTranslationByKey(t *testing.T)  {
	t.Log("Given the need to test GetTranslationByKey")
	{
		for _,local := range locales {
			for _,component := range components {
				resp := GetTranslationByKey("Payment",local,component)

				if resp.Response.Code != 200{
					t.Fatal("GetTranslationByKey failed!!!")
				}

				t.Log(resp.Data)
			}
		}
	}
}

func BenchmarkGetFormattingPatternsByLocal(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resp := GetFormattingPatternsByLocal(locales[0])
		if resp.Response.Code == 200{
			success++
		} else {
			fail++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkGetFormattingPatternsByLocal",success,fail)
}

func BenchmarkGetTranslationByComponent(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp := GetTranslationByComponent(locales[0],components[0])
		if resp.Response.Code == 200{
			success++
		} else {
			fail++
		}
	}
	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkGetTranslationByComponent",success,fail)
}

func BenchmarkGetTranslationByKey(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp := GetTranslationByKey("Payment",locales[0],components[0])
		if resp.Response.Code == 200{
			success++
		} else {
			fail++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkFormator",success,fail)
}