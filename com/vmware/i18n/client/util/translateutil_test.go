package util

import (
	"fmt"
	"log"
	"testing"
)

func TestGetTranslationByKey(t *testing.T) {
	value, err := getTranslationByKey("general.download", "zh_CN", "JS")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}
}

func TestGetTranslationByKeyWithParams(t *testing.T) {
	value, err := GetTranslationCacheManagerInstance().GetTranslationByKeyWithParams("general.download1", "zh_CN", "JS", "prefix", "suffix")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(value)
	}

	value1, err1 := GetTranslationCacheManagerInstance().GetTranslationByKeyWithParams("general.download1", "zh_CN", "JS", "prefix", "suffix")
	if err1 != nil {
		log.Fatal(err1)
	} else {
		fmt.Println(value1)
	}
}

func TestGetFormatMap(t *testing.T) {
	cacheFormatMap := *GetFormatMap()

	if cacheFormatMap["fr"].Messages.DateTimeFormat.Full == "" {
		t.Fatal("cacheFormat failed!!!")
	}

	t.Log(cacheFormatMap["fr"])
}
