package pokecache

import (
	"io"
	"net/http"
	"testing"
	"time"
)

func TestCacheSpeed(t *testing.T){
	testCache := NewCache(time.Second * 5)
	testUrl := "https://pokeapi.co/api/v2/location-area/"
	currTime := time.Now()
	res, err := http.Get(testUrl)
	if err != nil{
		t.Fatal(err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil{
		t.Fatal(err)
	}
	getTime := time.Now().Sub(currTime)
	testCache.Add(testUrl, data)
	currTime = time.Now()
	data, ok := testCache.Get(testUrl)
	if !ok{
		t.Fatal("data not stored in cache")
	}
	cacheTime := time.Now().Sub(currTime)
	if getTime < cacheTime{
		t.Errorf("%v should be less than %v", cacheTime, getTime)
	}
}