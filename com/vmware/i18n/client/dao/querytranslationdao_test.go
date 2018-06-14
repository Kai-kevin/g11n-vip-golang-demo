package dao

import (
	"testing"
	"log"
	"fmt"
)

func TestGetTranslationByKey(t *testing.T) {
	value,err := GetTranslationByKey("general.download","zh_CN","JS")

	if(err != nil){
		log.Fatal(err)
	}else{
		fmt.Println(value)
	}
}

func TestGetTranslationByKeyWithParams(t *testing.T) {
	value,err := GetTranslationByKeyWithParams("general.download1","zh_CN","JS","prefix","suffix")

	if(err != nil){
		log.Fatal(err)
	}else{
		fmt.Println(value)
	}

	value1,err1 := GetTranslationByKeyWithParams("general.download1","zh_CN","JS","prefix","suffix")
	if(err1 != nil){
		log.Fatal(err1)
	}else{
		fmt.Println(value1)
	}
}
