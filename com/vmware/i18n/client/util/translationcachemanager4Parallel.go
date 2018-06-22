package util

import (
	"fmt"
	"strings"
	"vipgoclient/com/vmware/i18n/client/conf"
	"vipgoclient/com/vmware/i18n/client/dao"
	"time"
	"sync"
)


//tranlation cache map
var cachedMap4Paral sync.Map

//format cache map
var cacheFormatMap4Paral sync.Map




func LoadAllCached4Paral() {

		initializeCache := conf.GetVipConfigInstance().InitializeCache

		if initializeCache == "true" {

			now := time.Now()

			var wg *sync.WaitGroup

			//default separator is ','
			locales := strings.Split(conf.GetVipConfigInstance().Locales, ",")


				wg.Add(len(locales))


			LoadCached4Paral(locales,wg)

			wg.Wait()

			fmt.Println("Load cached4Paral successfully, cost ",time.Since(now))
		}

}

//loading cache information
func LoadCached4Paral(locals []string,wg *sync.WaitGroup) {

	components := strings.Split(conf.GetVipConfigInstance().Components, ",")

	for _,locale := range locals{

		go func() {
			defer wg.Done()

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

				cachedMap4Paral.Store(cacheDTO,respEvent.Data.Messages)
			}

			//get the format patterns cache
			patternData := dao.GetFormattingPatternsByLocal(locale)

			cacheFormatMap4Paral.Store(locale,&patternData.Data)

			}()
	}

	}

