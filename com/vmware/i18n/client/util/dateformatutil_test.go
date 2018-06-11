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

	dateTime,_ := time.Parse(dateTimeString,dateTimeString)

	dFull,err := FormatDateTime(constants.FULL,dateTime,"fr")

	if err != nil {
		t.Fatal("TestFormatDateTime failed!!!")
	}

	fmt.Println(dFull)
}