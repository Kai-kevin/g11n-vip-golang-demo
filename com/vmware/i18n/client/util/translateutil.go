package util

import (
	"strings"
	"strconv"
	"vipgoclient/com/vmware/i18n/client/conf"
	"vipgoclient/com/vmware/i18n/client/dao"
	"errors"
)

type Translator struct{
	locale string
}

var enableCache = conf.GetVipConfigInstance().EnableCache
var productID = conf.GetVipConfigInstance().ProductId
var version = conf.GetVipConfigInstance().Version

//get the translated value by key
func getTranslationByKey(key, locale, component string) (string,error){

	if(enableCache == "false"){
		respEvent := dao.GetTranslationByKey(key ,locale, component)

		if(respEvent.Response.Code == 200){

			return respEvent.Data.Translation,nil
		}else{
			return "",errors.New("could't find the key:" + key + "value")
		}

	}else{
		cacheDTO := CacheDTO{
			Locale: locale,
			Component:component,
			ProductID: productID,
			Version:version,
		}

		cacheMap := *GetCacheMap()

		if cacheMap[cacheDTO][key] != ""{
			return cacheMap[cacheDTO][key],nil
		}else{
			respEvent := dao.GetTranslationByKey(key ,locale, component)

			if(respEvent.Response.Code == 200){

				//update cache
				defer GetTranslationCacheManagerInstance().UpdateCacheByComponent(cacheDTO,map[string]string{key:respEvent.Data.Translation})

				return respEvent.Data.Translation,nil
			}else{
				return "",errors.New("could't find the key:" + key + "value")
			}
		}

	}
}

//get the translated value by key and values
func (*translationCacheManager) GetTranslationByKeyWithParams(key, locale, component string,params ...interface{}) (string,error){

	value,err := getTranslationByKey(key,locale,component)

	if err != nil {
		return value,err
	}

	for index,param := range params{
		value = strings.Replace(value,"{"+strconv.Itoa(index)+"}",param.(string),-1)
	}

	return value,nil

}
