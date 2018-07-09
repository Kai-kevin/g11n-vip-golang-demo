package util

import (
	"testing"
)

var (
	cacheDTO = CacheDTO{
		Locale:    Locales[0],
		Component: Components[0],
		ProductID: productID,
		Version:   version,
	}
	params = map[string]string{"general.download": "Test updateCacheByComponent", "=========": "********"}
)

func TestLoadCached(t *testing.T) {
	LoadCached()
	cacheMap := *GetCacheMap()
	t.Logf("Result\"%v\"", cacheMap)
	if len(cacheMap) != len(Locales)*len(Components) {
		t.Fatal("Test LoadCached failed!!!")
	}
	LoadAllCached4Paral()
}

func TestRemoveCacheByComponent(t *testing.T) {
	t.Logf("The result before the removeCache is \"%v\"", cachedMap)
	if !GetTranslationCacheManagerInstance().RemoveCacheByComponent(cacheDTO) {
		t.Fatal("Remove cache failed!!!")
	} else if _, exist := cachedMap[cacheDTO]; exist {
		t.Fatal("The result of the removeCacheByComponent is not the excepted!!!")
	}
	t.Logf("The result after the removeCache is \"%v\"", cachedMap)
}

func TestUpdateCacheByComponent(t *testing.T) {
	t.Logf("The result before the updateCache is \"%s\"", cachedMap[cacheDTO])
	if !GetTranslationCacheManagerInstance().UpdateCacheByComponent(cacheDTO, params) {
		t.Fatal("UpdateCacheByComponent failed!!!")
	}
	t.Logf("The result after the updateCache is \"%s\"", cachedMap[cacheDTO])

	for k, v := range params {
		if cachedMap[cacheDTO][k] != v {
			t.Fatalf("The result of the updateCacheByComponent is \"%v\"; Not the excepted value:\"%v\"", cachedMap[cacheDTO][k], v)
		}
	}
}

func TestAddCacheByComponent(t *testing.T) {
	t.Logf("The result before the addCache is \"%s\"", cachedMap[cacheDTO])
	GetTranslationCacheManagerInstance().maxNumOfComponentInCache = 30
	if !GetTranslationCacheManagerInstance().AddCacheByComponent(cacheDTO, params) {
		t.Fatal("AddCacheByComponent failed!!!")
	}
	t.Logf("The result after the addCache is \"%s\"", cachedMap[cacheDTO])

	for k, v := range params {
		if cachedMap[cacheDTO][k] != v {
			t.Fatalf("The result of the addCacheByComponent is \"%v\"; Not the excepted value:\"%v\"", cachedMap[cacheDTO][k], v)
		}
	}
}

func TestLookForTranslationlnCache(t *testing.T) {
	value := GetTranslationCacheManagerInstance().LookForTranslationlnCache(Key, cacheDTO)
	t.Logf("The result of the lookForTranslationlnCache is \"%v\"", value)
	if value == "" {
		t.Fatalf("cacheDTO do not contain \"%s\"!!!", Key)
	} else if cachedMap[cacheDTO][Key] != value {
		t.Fatalf("The result of the lookForTranslationlnCache is \"%v\"; Not the excepted value:\"%v\"", value, cachedMap[cacheDTO][Key])
	}
}

func BenchmarkLoadCached(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LoadCached()
		LoadAllCached4Paral()
	}
}

func BenchmarkRemoveCacheByComponent(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTranslationCacheManagerInstance().RemoveCacheByComponent(cacheDTO)
	}
}

func BenchmarkUpdateCacheByComponent(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTranslationCacheManagerInstance().UpdateCacheByComponent(cacheDTO, params)
	}
}

func BenchmarkAddCacheByComponent(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTranslationCacheManagerInstance().maxNumOfComponentInCache = len(cachedMap) + 1
		GetTranslationCacheManagerInstance().AddCacheByComponent(cacheDTO, params)
	}
}

func BenchmarkLookForTranslationlnCache(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTranslationCacheManagerInstance().LookForTranslationlnCache(Key, cacheDTO)
	}
}
