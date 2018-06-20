package util

import (
	"testing"
	"fmt"
)

func TestFormatNumber(t *testing.T) {
	format,err := GetNumberFormatByLocal("fr")

	if err != nil{
		t.Fatal("Test formatNumber failed!!!")
	}

	i  := FormatNumber(&format,-1212.1211)

	fmt.Println(i)

}

func TestFormatPercent(t *testing.T) {
	format,err := GetNumberFormatByLocal("fr")

	if err != nil{
		t.Fatal("Test formatNumber failed!!!")
	}

	i  := FormatPercent(&format,1212.1211)

	fmt.Println(i)

}

func TestFormatCurrency(t *testing.T) {

	format,err := GetCurrencyFormatByLocal("fr")

	if err != nil{
		t.Fatal("Test formatNumber failed!!!")
	}

	i  := FormatCurrency(&format,1212.1211)
	j  := FormatCurrency(&format,-1212.1211)

	fmt.Println(i)
	fmt.Println(j)
}


