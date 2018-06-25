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

var localsChan chan string

func LoadAllCached4Paral() {

		defer func() {
			if err := recover();err != nil{
				fmt.Println("Load cached failed:",err)
			}
		}()

		initializeCache := conf.GetVipConfigInstance().InitializeCache

		if initializeCache == "true" {

			now := time.Now()

			//default separator is ','
			locales := strings.Split(conf.GetVipConfigInstance().Locales, ",")

			var wg sync.WaitGroup

			localsChan = make(chan string,len(locales))

			LoadCached4Paral(locales,&wg)

			for _,locale := range locales {

				localsChan <- locale
				wg.Add(1)
			}

			close(localsChan)

			wg.Wait()

			fmt.Println("Load cached4Paral successfully, cost ",time.Since(now))
		}

}

//loading cache information
func LoadCached4Paral(locals []string,wg *sync.WaitGroup) {

	components := strings.Split(conf.GetVipConfigInstance().Components, ",")

	for i:= 0; i < len(locals); i++ {

		go func() {
			defer wg.Done()

			locale4Tmp := <- localsChan

			fmt.Println(locale4Tmp)

			for _, component := range components {

				cacheDTO := CacheDTO{
					Locale:    locale4Tmp,
					Component: component,
					ProductID: productID,
					Version:   version,
				}

				//get the component translations
				respEvent := dao.GetTranslationByComponent(locale4Tmp, component)

				//get cache messages
				cachedMap4Paral.Store(cacheDTO,respEvent.Data.Messages)
			}

			//get the format patterns cache
			patternData := dao.GetFormattingPatternsByLocal(locale4Tmp)

			cacheFormatMap4Paral.Store(locale4Tmp,&patternData.Data)

			fmt.Println(cacheFormatMap4Paral)

		}()

	}

	}



func GetCache4ParalMap() *sync.Map {
	return &cachedMap4Paral
}

func GetFormat4ParalMap() *sync.Map{
	return &cacheFormatMap4Paral
}

