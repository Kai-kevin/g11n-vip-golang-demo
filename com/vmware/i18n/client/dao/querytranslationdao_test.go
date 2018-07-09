package dao

import (
	"testing"
	"vipgoclient/com/vmware/i18n/client/conf"
	"strings"
)

var locales = strings.Split(conf.GetVipConfigInstance().Locales, ",")
var components = strings.Split(conf.GetVipConfigInstance().Components, ",")

func TestGetFormattingPatternsByLocal(t *testing.T) {
	for _, local := range locales {
		t.Logf("Request parameter:local:\"%v\"", local)
		resp := GetFormattingPatternsByLocal(local)
		t.Logf("The result is \"%v\"", resp)
		if resp.Response.Code != 200 {
			t.Fatal("GetFormattingPatternsByLocal failed!!!")
		}
	}
}

func TestGetTranslationByComponent(t *testing.T) {
	for _, local := range locales {
		for _, component := range components {

			t.Logf("Request parameters:local:\"%v\",component:\"%v\"", local, component)
			resp := GetTranslationByComponent(local, component)

			t.Logf("The result is \"%v\"", resp)
			if resp.Response.Code != 200 {
				t.Fatal("TestGetTranslationByComponent failed!!!")
			}
		}
	}
}

func TestGetTranslationByKey(t *testing.T) {
	for _, local := range locales {
		for _, component := range components {
			t.Logf("Request parameters:local:\"%v\",component:\"%v\"", local, component)
			resp := GetTranslationByKey("Payment", local, component)
			t.Logf("The result is \"%v\"", resp)
			if resp.Response.Code != 200 {
				t.Fatal("GetTranslationByKey failed!!!")
			}
		}
	}
}

func BenchmarkGetFormattingPatternsByLocal(b *testing.B) {
	success, fail := 0, 0

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resp := GetFormattingPatternsByLocal(locales[0])
		if resp.Response.Code == 200 {
			success++
		} else {
			fail++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t", "BenchmarkGetFormattingPatternsByLocal", success, fail)
}

func BenchmarkGetTranslationByComponent(b *testing.B) {
	success, fail := 0, 0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp := GetTranslationByComponent(locales[0], components[0])
		if resp.Response.Code == 200 {
			success++
		} else {
			fail++
		}
	}
	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t", "BenchmarkGetTranslationByComponent", success, fail)
}

func BenchmarkGetTranslationByKey(b *testing.B) {
	success, fail := 0, 0

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp := GetTranslationByKey("Payment", locales[0], components[0])
		if resp.Response.Code == 200 {
			success++
		} else {
			fail++
		}
	}

	b.Logf("Title:\"%s\"\tSuccess:\"%v\"\tfaile:\"%v\"\t", "BenchmarkFormator", success, fail)
}
