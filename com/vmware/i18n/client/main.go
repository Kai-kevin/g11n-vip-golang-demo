package main

import (
	"vipgoclient/com/vmware/i18n/client/util"
	"vipgoclient/com/vmware/i18n/client/bean/i18n"
	"fmt"
	"encoding/json"
)

func main(){

	respData := bean.QueryTranslationByCompRespData{
		ProductName:"1111",
		Version:"1.0",
		Pseudo:true,
		Messages: map[string]string{
			"key1":"value1",
			"key2":"value2",
		},
	}

	respEvent := bean.QueryTranslationByCompRespEvent{
		Data: respData,
		Signature: "sig11111",
	}

	file,_ := json.Marshal(respEvent)

	fmt.Println(string(file))

	util.GetTranslationByComponent()
}
