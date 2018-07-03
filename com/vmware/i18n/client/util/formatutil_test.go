package util

import (
	"testing"
)

var (
	number = -1234.5679
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
		"zh_CN":"-123,457%",
		"de":"-123.457 %",
		"es":"-123.457 %",
		"fr":"-123 457 %",
		"ja":"-123,457%",
		"ko":"-123,457%",
		"zh_TW":"-123,457%",
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
)

func TestFormatNumber(t *testing.T) {
	t.Log("Given the need to Test Number Format")
	{
		for _, local := range Locales {
			format, err := GetNumberFormatByLocal(local)
			if err != nil {
				t.Fatal("Test formatNumber failed!!!")
			}

			t.Logf("the number format is \"%v\"",format)
			i := FormatNumber(&format, number)
			_,exist := exceptedNumberValue[local]
			if exist && exceptedNumberValue[local] != i {
				t.Fatalf("The result of the getNumber is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("The result of the number format in \"%s\" is \"%v\"",local,i)
		}
	}
}

func TestFormatPercent(t *testing.T) {
	t.Log("Given the need to Test Percent Format")
	{
		for _, local := range Locales {
			format,err := GetPercentFormatByLocal(local)
			if err != nil{
				t.Fatal("Test formatPercent failed!!!")
			}

			t.Logf("the percent format is \"%v\"",format)
			i := FormatPercent(&format,number)
			_,exist := exceptedPercentValue[local]
			if exist && exceptedPercentValue[local] != i {
				t.Fatalf("The result of the formatPercent is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("The result of the percent format in \"%s\" is \"%v\"",local,i)
		}
	}
}

func TestFormatCurrency(t *testing.T) {
	t.Log("Given the need to Test Currency Format")
	{
		for _, local := range Locales {
			format,err := GetCurrencyFormatByLocal(local)
			if err != nil{
				t.Fatal("Test formatCurrency failed!!!")
			}

			t.Logf("the currency format is \"%v\"",format)
			i := FormatCurrency(&format,number)
			_,exist := exceptedCurrencyValue[local]
			if exist && exceptedCurrencyValue[local] != i {
				t.Fatalf("The result of the formatCurrency is not the excepted value in local:\"%s\"!!!",local)
			}

			t.Logf("The result of the currency format in \"%s\" is \"%v\"",local,i)
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
		format,err := GetPercentFormatByLocal(Locales[LocaleNumber])

		if err != nil {
			b.Fatal("Benchmark formatPercent failed!!!")
		}

		FormatPercent(&format,number)
	}
}

func BenchmarkFormatCurrency(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		format,err := GetCurrencyFormatByLocal(Locales[LocaleNumber])

		if err != nil {
			b.Fatal("Benchmark formatCurrency failed!!!")
		}

		FormatCurrency(&format,number)
	}
}