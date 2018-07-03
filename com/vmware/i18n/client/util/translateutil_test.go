package util

import (
	"testing"
	"strings"
	"vipgoclient/com/vmware/i18n/client/conf"
	"math/rand"
	"time"
	"vipgoclient/com/vmware/i18n/client/constants"
)

var (
	Locales = strings.Split(conf.GetVipConfigInstance().Locales, ",")
 	Components = strings.Split(conf.GetVipConfigInstance().Components, ",")
 	Key = "general.download"
 	param = ""
 	LocaleNumber,ConstantNumber,ComponentNumber int
)

func init()  {
	rand.Seed(time.Now().Unix())
	LocaleNumber = rand.Intn(len(Locales))
	ConstantNumber = rand.Intn(constants.FULL)
	ComponentNumber = rand.Intn(len(Components))
}

func TestGetTranslationByKey(t *testing.T) {
	t.Log("Given the need to test getTranslationByKey")
	{
		for _,local := range Locales {
			for _,component := range Components {
				value, err := getTranslationByKey(Key, local, component)

				if err != nil {
					t.Fatalf("GetTranslationByKey failed:\"%s\"\n",err)
				} else {
					t.Logf("The key:\"%s\" convert to \"%s\" in component:\"%s\",local:\"%s\"",Key,value,component,local)
				}
			}
		}
	}
}

func TestGetTranslationByKeyWithParams(t *testing.T) {
	t.Log("Given the need to test getTranslationByKeyWithParams")
	{
		for _, local := range Locales {
			for _, component := range Components {
				value, err := GetTranslationCacheManagerInstance().GetTranslationByKeyWithParams(Key, local, component, param, param)

				if err != nil {
					t.Fatalf("GetTranslationByKeyWithParams failed:\"%s\"\n", err)
				} else {
					t.Logf("The key:\"%s\" convert to \"%s\" in component:\"%s\",local:\"%s\"", Key, value, component, local)
				}
			}
		}
	}
}

func TestGetFormatMap(t *testing.T) {
	cacheFormatMap := *GetFormatMap()

	for _,local := range Locales {
		if cacheFormatMap[local].Messages.DateTimeFormat.Full == "" {
			t.Fatal("cacheFormat failed!!!")
		}
		t.Log(cacheFormatMap[local])
	}
}

func BenchmarkGetTranslationByKey(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := getTranslationByKey(Key, Locales[LocaleNumber], Components[ComponentNumber])
		if err != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkGetTranslationByKey",success,fail)
}

func BenchmarkGetTranslationByKeyWithParams(b *testing.B) {
	success,fail := 0,0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetTranslationCacheManagerInstance().GetTranslationByKeyWithParams(Key, Locales[LocaleNumber], Components[LocaleNumber],param,param)
		if err != nil{
			fail++
		} else {
			success++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t","BenchmarkGetTranslationByKeyWithParams",success,fail)
}
