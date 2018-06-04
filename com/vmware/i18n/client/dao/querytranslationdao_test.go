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
