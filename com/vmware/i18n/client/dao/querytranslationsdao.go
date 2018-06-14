package dao

import (
	"vipgoclient/com/vmware/i18n/client/conf"
	"vipgoclient/com/vmware/i18n/client/util"
	"errors"
	"strings"
	"strconv"
)

var enableCache = conf.GetVipConfigInstance().EnableCache
var productID = conf.GetVipConfigInstance().ProductId
var version = conf.GetVipConfigInstance().Version

//get the translated value by key
func GetTranslationByKey(key, locale, component string) (string,error){

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

		if cacheMap[cacheDTO][key] != ""{
			return cacheMap[cacheDTO][key],nil
		}else{
			respEvent := util.GetTranslationByKey(key ,locale, component)

			if(respEvent.Response.Code == 200){

				//update cache
				defer util.GetTranslationCacheManagerInstance().UpdateCacheByComponent(cacheDTO,map[string]string{key:respEvent.Data.Translation})

				return respEvent.Data.Translation,nil
			}else{
				return "",errors.New("could't find the key:" + key + "value")
			}
		}


		}
}

func GetTranslationByKeyWithParams(key, locale, component string,params ...interface{}) (string,error){

	value,err := GetTranslationByKey(key,locale,component)

	if err != nil {
		return value,err
	}

	for index,param := range params{
		value = strings.Replace(value,"{"+strconv.Itoa(index)+"}",param.(string),-1)
	}

	return value,nil

}