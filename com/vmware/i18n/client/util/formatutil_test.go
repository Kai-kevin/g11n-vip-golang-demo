package util

import (
	"testing"
)

var (
	number              = -1234.5679
	exceptedNumberValue = map[string]string{
		"zh_CN": "-1,234.568",
		"de":    "-1.234,568",
		"es":    "-1.234,568",
		"fr":    "-1 234,568",
		"ja":    "-1,234.568",
		"ko":    "-1,234.568",
		"zh_TW": "-1,234.568",
	}
	exceptedPercentValue = map[string]string{
		"zh_CN": "-123,457%",
		"de":    "-123.457 %",
		"es":    "-123.457 %",
		"fr":    "-123 457 %",
		"ja":    "-123,457%",
		"ko":    "-123,457%",
		"zh_TW": "-123,457%",
	}
	exceptedCurrencyValue = map[string]string{
		"zh_CN": "￥-1,234.57",
		"de":    "-1.234,57 €",
		"es":    "-1.234,57 €",
		"fr":    "-1 234,57 €",
		"ja":    "￥-1,234.57",
		"ko":    "₩-1,234.57",
		"zh_TW": "CN¥-1,234.57",
	}
)

func TestFormatNumber(t *testing.T) {
	for _, local := range Locales {
		t.Logf("Request parameter:local:\"%v\"", local)
		format, err := GetNumberFormatByLocal(local)
		if err != nil {
			t.Fatal("Test formatNumber failed!!!")
		}

		t.Logf("NumberFormat:\"%v\"", format)
		i := FormatNumber(&format, number)
		t.Logf("Result:\"%v\" convert to \"%v\"", number, i)
		_, exist := exceptedNumberValue[local]
		if exist && exceptedNumberValue[local] != i {
			t.Fatalf("The result of the getNumber is not the excepted value in local:\"%s\"!!!", local)
		}
	}
}

func TestFormatPercent(t *testing.T) {
	for _, local := range Locales {
		t.Logf("Request parameter:local:\"%v\"", local)
		format, err := GetPercentFormatByLocal(local)
		if err != nil {
			t.Fatal("Test formatPercent failed!!!")
		}

		t.Logf("PercentFormat:\"%v\"", format)
		i := FormatPercent(&format, number)
		t.Logf("Result:\"%v\" convert to \"%v\"", number, i)
		_, exist := exceptedPercentValue[local]
		if exist && exceptedPercentValue[local] != i {
			t.Fatalf("The result of the formatPercent is not the excepted value in local:\"%s\"!!!", local)
		}
	}
}

func TestFormatCurrency(t *testing.T) {
	for _, local := range Locales {
		t.Logf("Request parameter:local:\"%v\"", local)
		format, err := GetCurrencyFormatByLocal(local)
		if err != nil {
			t.Fatal("Test formatCurrency failed!!!")
		}

		t.Logf("CurrencyFormat:\"%v\"", format)
		i := FormatCurrency(&format, number)
		t.Logf("Result:\"%v\" convert to \"%v\"", number, i)
		_, exist := exceptedCurrencyValue[local]
		if exist && exceptedCurrencyValue[local] != i {
			t.Fatalf("The result of the formatCurrency is not the excepted value in local:\"%s\"!!!", local)
		}
	}
}

func BenchmarkFormatNumber(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format, err := GetNumberFormatByLocal(Locales[LocaleNumber])

		if err != nil {
			b.Fatal("Benchmark formatNumber failed!!!")
		}

		FormatNumber(&format, number)
	}
}

func BenchmarkFormatPercent(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format, err := GetPercentFormatByLocal(Locales[LocaleNumber])

		if err != nil {
			b.Fatal("Benchmark formatPercent failed!!!")
		}

		FormatPercent(&format, number)
	}
}

func BenchmarkFormatCurrency(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format, err := GetCurrencyFormatByLocal(Locales[LocaleNumber])

		if err != nil {
			b.Fatal("Benchmark formatCurrency failed!!!")
		}

		FormatCurrency(&format, number)
	}
}
