package dao

import (
	"testing"
	"vipgoclient/com/vmware/i18n/client/util"
)

func TestGetFormattingPatternsByLocal(t *testing.T) {
	resp := GetFormattingPatternsByLocal("fr")

	if resp.Response.Code != 200{
		t.Fatal("GetFormattingPatternsByLocal failed!!!")
	}

	t.Log(resp.Data.Messages)

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

func TestGetFormatMap(t *testing.T){
	cacheFormatMap := *util.GetFormatMap()

	if cacheFormatMap["fr"].Messages.DateTimeFormat.Full == ""{
		t.Fatal("cacheFormat failed!!!")
	}

	t.Log(cacheFormatMap["fr"])
}