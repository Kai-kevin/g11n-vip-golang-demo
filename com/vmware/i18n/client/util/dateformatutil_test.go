package util

import (
	"testing"
	"time"
	"vipgoclient/com/vmware/i18n/client/constants"
	"fmt"
)

var(
	dateTimeString = "Jan 2, 2006 at 3:04:05pm"
)

func TestFormatDateTime(t *testing.T) {

	var cstSh, _ = time.LoadLocation("Asia/Shanghai")

	dateTime,_ := time.ParseInLocation(dateTimeString,dateTimeString,cstSh)

	fmt.Println(dateTime.Zone())

	dFull,err := FormatDateTime(constants.FULL,dateTime,"fr")

	if err != nil {
		t.Fatal("TestFormatDateTime failed!!!")
	}

	fmt.Println(dFull)

	dShortTW,errTW := FormatDateTime(constants.LONG,dateTime,"zh_TW")

	if errTW != nil {
		t.Fatal("TestFormatDateTime failed!!!")
	}

	fmt.Println(dShortTW)
}