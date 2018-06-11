package util

import (
	"testing"
	"fmt"
)

func TestFormatNumber(t *testing.T) {

	locale,err := GetNumberFormatByLocal("fr")

	if err != nil{
		t.Fatal("Test formatNumber failed!!!")
	}

	i  := FormatNumber(&locale,-1212.1211)

	fmt.Println(i)

}

func TestFormatPercent(t *testing.T) {
	locale,err := GetNumberFormatByLocal("fr")

	if err != nil{
		t.Fatal("Test formatNumber failed!!!")
	}

	i  := FormatPercent(&locale,1212.1211)

	fmt.Println(i)

}


