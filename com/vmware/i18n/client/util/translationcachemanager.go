package util

import (
	"fmt"
	"strings"
	"sync"
	"vipgoclient/com/vmware/i18n/client/conf"
	"vipgoclient/com/vmware/i18n/client/bean/i18n"
)

type CacheDTO struct {
	ProductID string
	Version   string
	Locale    string
	Component string
}

//缓存MAP
var cachedMap = make(map[CacheDTO]map[string]string)

//缓存format相关信息
var cacheFormatMap = make(map[string]*bean.QueryFormattingPatternByLocaleRespData)

//读写锁
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

//加载缓存信息，缓存的加载方式和逻辑
func LoadCached() {
	//获取Local信息和Component信息
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

			//通过component获取翻译信息
			respEvent := GetTranslationByComponent(locale, component)

			//获取缓存的信息
			cachedMap[cacheDTO] = respEvent.Data.Messages

			//获取Format缓存信息
			patternData := GetFormattingPatternsByLocal(locale)
			cacheFormatMap[locale] = &
				patternData.Data
		}
	}
}

//获取缓存Map
func GetCacheMap() *map[CacheDTO]map[string]string {
	return &cachedMap
}

//获取当前缓存FormatMap
func GetFormatMap() *map[string]*bean.QueryFormattingPatternByLocaleRespData{
	return &cacheFormatMap
}

//查询缓存信息
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

//增加缓存
func (manager *translationCacheManager) AddCacheByComponent(dto CacheDTO, object interface{}) bool {

	if len(cachedMap) >= manager.maxNumOfComponentInCache {
		fmt.Println("CachedMap has alreay exceeded!")
		return false
	}

	translationCache := cachedMap[dto]

	mux.Lock()

	defer mux.Unlock()

	//获取组件最大数值
	if len(translationCache) < manager.maxNumOfTranslationInComponent {
		return false
	} else {
		//数组更新
		translationRets := object.(map[string]string)

		for k, v := range translationRets {
			translationCache[k] = v
		}
	}

	return true
}

//删除缓存
func (*translationCacheManager) RemoveCacheByComponent(dto CacheDTO) bool {
	mux.Lock()

	defer mux.Unlock()

	//删除参数
	delete(cachedMap, dto)

	return true
}

//更新缓存
func (*translationCacheManager) UpdateCacheByComponent(dto CacheDTO, object interface{}) bool {
	return false
}
