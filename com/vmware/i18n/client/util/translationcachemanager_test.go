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
	params = map[string]string{"general.download": "Test updateCacheByComponent","=========":"********"}
)

func TestLoadCached(t *testing.T) {
	t.Log("Test loadCached")
	{
		LoadCached()
		LoadAllCached4Paral()
		cacheMap := *GetCacheMap()
		if len(cacheMap) != len(Locales)*len(Components) {
			t.Fatal("Test LoadCached failed!!!")
		}

		t.Logf("The result of loadCached is \"%v\"",cacheMap)
	}

}

func TestRemoveCacheByComponent(t *testing.T) {
	t.Log("Test removeCacheByComponent")
	{
		t.Log(len(cachedMap))
		if !GetTranslationCacheManagerInstance().RemoveCacheByComponent(cacheDTO) {
			t.Fatal("Remove cache failed!!!")
		} else if _,exist := cachedMap[cacheDTO]; exist {
			t.Fatal("The result of the removeCacheByComponent is not the excepted!!!")
		}
		t.Log(len(cachedMap))

	}
}

func TestUpdateCacheByComponent(t *testing.T) {
	t.Log("Test updateCacheByComponent")
	{
		t.Logf("The old cachedMap is \"%s\"", cachedMap[cacheDTO])
		if !GetTranslationCacheManagerInstance().UpdateCacheByComponent(cacheDTO, params) {
			t.Fatal("UpdateCacheByComponent failed!!!")
		}

		for k,v := range params {
			if cachedMap[cacheDTO][k] != v {
				t.Fatalf("The result of the updateCacheByComponent is \"%v\"; Not the excepted value:\"%v\"", cachedMap[cacheDTO][k], v)
			}
		}

		t.Logf("The Result of the updateCacheByComponent is \"%s\"", cachedMap[cacheDTO])
	}
}

func TestAddCacheByComponent(t *testing.T) {
	t.Log("Test addCacheByComponent")
	{
		t.Logf("The old cachedMap is \"%v\"", len(cachedMap))
		GetTranslationCacheManagerInstance().maxNumOfComponentInCache = 30
		if !GetTranslationCacheManagerInstance().AddCacheByComponent(cacheDTO, params) {
			t.Fatal("AddCacheByComponent failed!!!")
		}

		for k, v := range params {
			if cachedMap[cacheDTO][k] != v {
				t.Fatalf("The result of the addCacheByComponent is \"%v\"; Not the excepted value:\"%v\"", cachedMap[cacheDTO][k], v)
			}
		}
		t.Logf("The Result of the addCacheByComponent is \"%v\"", cachedMap[cacheDTO])

	}
}

func TestLookForTranslationlnCache(t *testing.T) {
	t.Log("Test lookForTranslationlnCache")
	{
		value := GetTranslationCacheManagerInstance().LookForTranslationlnCache(Key, cacheDTO)
		if value == "" {
			t.Fatalf("cacheDTO do not contain \"%s\"!!!", Key)
		} else if cachedMap[cacheDTO][Key] != value {
			t.Fatalf("The result of the lookForTranslationlnCache is \"%v\"; Not the excepted value:\"%v\"", value, cachedMap[cacheDTO][Key])
		}

		t.Logf("The result of the lookForTranslationlnCache is \"%v\"", value)

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
		GetTranslationCacheManagerInstance().maxNumOfComponentInCache = len(cachedMap) +1
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