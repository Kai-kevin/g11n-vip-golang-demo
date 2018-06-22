package util

import (
	"fmt"
	"strings"
	"sync"
	"vipgoclient/com/vmware/i18n/client/conf"
	"vipgoclient/com/vmware/i18n/client/bean/i18n"
	"vipgoclient/com/vmware/i18n/client/dao"
	"time"
)

type CacheDTO struct {
	ProductID string
	Version   string
	Locale    string
	Component string
}

//tranlation cache map
var cachedMap = make(map[CacheDTO]map[string]string)

//format cache map
var cacheFormatMap = make(map[string]*bean.QueryFormattingPatternByLocaleRespData)

var mux sync.RWMutex

var once sync.Once

type translationCacheManager struct {
	maxNumOfComponentInCache       int
	maxNumOfTranslationInComponent int
	cacheEnabled                   bool
}

var translationCacheManagerInstance = translationCacheManager{
	maxNumOfComponentInCache:       10,
	maxNumOfTranslationInComponent: 10,
	cacheEnabled:                   true,
}

func GetTranslationCacheManagerInstance() *translationCacheManager {
	return &translationCacheManagerInstance
}

func init() {
	once.Do(func() {
		initializeCache := conf.GetVipConfigInstance().InitializeCache

		if initializeCache == "true" {
			LoadCached()
		}
	})
}

//loading cache information
func LoadCached() {

	now := time.Now();

	//default separator is ','
	locales := strings.Split(conf.GetVipConfigInstance().Locales, ",")
	components := strings.Split(conf.GetVipConfigInstance().Components, ",")
	productID := conf.GetVipConfigInstance().ProductId
	version := conf.GetVipConfigInstance().Version

	for _, locale := range locales {

		for _, component := range components {

			cacheDTO := CacheDTO{
				Locale:    locale,
				Component: component,
				ProductID: productID,
				Version:   version,
			}

			//get the component translations
			respEvent := dao.GetTranslationByComponent(locale, component)

			//get cache messages
			cachedMap[cacheDTO] = respEvent.Data.Messages
		}

		//get the format patterns cache
		patternData := dao.GetFormattingPatternsByLocal(locale)
		cacheFormatMap[locale] = &
			patternData.Data
	}

	fmt.Println("Load cached successfully, cost ",time.Since(now))
}

func GetCacheMap() *map[CacheDTO]map[string]string {
	return &cachedMap
}

func GetFormatMap() *map[string]*bean.QueryFormattingPatternByLocaleRespData{
	return &cacheFormatMap
}

func (*translationCacheManager) LookForTranslationlnCache(key string, dto CacheDTO) string {
	mux.RLock()
	defer mux.RUnlock()

	translationCache := cachedMap[dto]

	value, ok := translationCache[key]

	if ok {
		return value
	}
	return ""
}

func (manager *translationCacheManager) AddCacheByComponent(dto CacheDTO, object interface{}) bool {

	//compare to the maximum number of components
	if len(cachedMap) >= manager.maxNumOfComponentInCache {
		fmt.Println("CachedMap has alreay exceeded!")
		return false
	}

	translationCache := cachedMap[dto]

	mux.Lock()

	defer mux.Unlock()

	//compare to the maximum number of translations in a component
	if len(translationCache) < manager.maxNumOfTranslationInComponent {
		return false
	} else {
		//update cache
		translationRets := object.(map[string]string)

		for k, v := range translationRets {
			translationCache[k] = v
		}
	}

	return true
}

//delete cache
func (*translationCacheManager) RemoveCacheByComponent(dto CacheDTO) bool {
	mux.Lock()
	defer mux.Unlock()

	delete(cachedMap, dto)

	return true
}


func (*translationCacheManager) UpdateCacheByComponent(dto CacheDTO, object interface{}) bool {
	mux.Lock()
	defer mux.Unlock()

	translationCache := cachedMap[dto]
	translationRets := object.(map[string]string)

	for k, v := range translationRets {
		translationCache[k] = v
	}

	return true
}
