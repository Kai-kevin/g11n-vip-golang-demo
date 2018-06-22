package util

import (
	"testing"
	"fmt"
)

func TestLoadCached(t *testing.T) {

	LoadCached()

	cacheMap := *GetCacheMap()

	fmt.Println(cacheMap)

	LoadAllCached4Paral()
}
