package dao

import (
	"vipgoclient/com/vmware/i18n/client/conf"
	"vipgoclient/com/vmware/i18n/client/util"
	"errors"
)

var enableCache = conf.GetVipConfigInstance().EnableCache
var productID = conf.GetVipConfigInstance().ProductId
var version = conf.GetVipConfigInstance().Version

//获取翻译的值
func GetTranslationByKey(key ,locale, component string) (string,error){

	if(enableCache == "false"){
		respEvent := util.GetTranslationByKey(key ,locale, component)

		if(respEvent.Response.Code == 200){

			return respEvent.Data.Translation,nil
		}else{
			return "",errors.New("could't find the key:" + key + "value")
		}

	}else{
		cacheDTO := util.CacheDTO{
			Locale: locale,
			Component:component,
			ProductID: productID,
			Version:version,
		}

		cacheMap := *util.GetCacheMap()

		return cacheMap[cacheDTO][key],nil
		}
}
