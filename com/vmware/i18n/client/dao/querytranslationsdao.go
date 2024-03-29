package dao

import (
	"vipgoclient/com/vmware/i18n/client/conf"
	"vipgoclient/com/vmware/i18n/client/bean/i18n"
	"net/http"
	"encoding/json"
	"fmt"
)

var enableCache = conf.GetVipConfigInstance().EnableCache
var productID = conf.GetVipConfigInstance().ProductId
var version = conf.GetVipConfigInstance().Version

func GetTranslationByComponent(local, component string) *bean.QueryTranslationByCompRespEvent {

	//config := consulapi.DefaultConfig()
	//
	//client, err := consulapi.NewClient(config)

	//load balance by consul
	//servicesData, _, err := client.Health().Service("i18n", "", false,
	//	&consulapi.QueryOptions{})
	//
	//for _, entry := range servicesData {
	//
	//	fmt.Println(entry.Service.Address)
	//	fmt.Println(entry.Service.Port)
	//}
	//
	//fmt.Print(servicesData)
	//
	//fmt.Println(client.Agent().NodeName())
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	productName := conf.GetVipConfigInstance().ProductName
	//version := conf.GetVipConfigInstance().Version
	//
	//url := "http://" + servicesData[0].Service.Address + ":" + strconv.Itoa(servicesData[0].Service.Port) +
	//	"/i18n/api/v2/translation/products/" + productName + "/versions/" + version +
	//	"/locales/" + local + "/components/" + component

	url := "http://" + conf.GetVipConfigInstance().VipServer +
		"/i18n/api/v2/translation/products/" + productName + "/versions/" + version +
		"/locales/" + local + "/components/" + component

	fmt.Println(url)

	resp, _ := http.Get(url)

	respData := new(bean.QueryTranslationByCompRespEvent)

	error := json.NewDecoder(resp.Body).Decode(respData)

	if error != nil {
		fmt.Println(url + error.Error())
	}

	fmt.Println(respData.Data.Messages)

	return respData
}

//get translated value by key
func GetTranslationByKey(key, component, local string) *bean.QueryTranslationByKeyRespEvent {

	productName := conf.GetVipConfigInstance().ProductName
	version := conf.GetVipConfigInstance().Version

	url := "http://" + conf.GetVipConfigInstance().VipServer +
		"/i18n/api/v2/translation/products/" + productName + "/versions/" +
		version + "/locales/" + local + "/components/" + component + "/keys/" + key

	fmt.Println(url)

	resp, _ := http.Get(url)

	respData := new(bean.QueryTranslationByKeyRespEvent)

	error := json.NewDecoder(resp.Body).Decode(respData)

	if error != nil {
		fmt.Println(url + error.Error())
	}

	fmt.Println(respData)

	return respData
}

//get format information by local string
func GetFormattingPatternsByLocal(local string) *bean.QueryFormattingPatternByLocaleRespEvent {

	url := "http://" + conf.GetVipConfigInstance().VipServer + "/i18n/api/v2/formatting/patterns/locales/" + local

	fmt.Println(url)

	resp, _ := http.Get(url)

	respData := new(bean.QueryFormattingPatternByLocaleRespEvent)

	error := json.NewDecoder(resp.Body).Decode(respData)

	if error != nil {
		fmt.Println(url + error.Error())
	}

	fmt.Println(respData.Data)

	return respData
}


