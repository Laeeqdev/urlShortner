package repository

import (
	"fmt"
	"sync"
)

type LongUrlToShortUrlMap map[string]string
type ShortUrlToLongUrlMap map[string]string

func ProvideLongUrlToShortUrlMap() LongUrlToShortUrlMap {
	return LongUrlToShortUrlMap(make(map[string]string))
}

func ProvideShortUrlToLongUrlMap() ShortUrlToLongUrlMap {
	return ShortUrlToLongUrlMap(make(map[string]string))
}

func ProvideMutex() *sync.RWMutex {
	return &sync.RWMutex{}
}

type ShortUrlRepository interface {
	GetShortUrlByLongUrl(longUrl string) (string, bool)
	GetLongUrlByShortUrl(shortUrl string) (string, bool)
	AddUrl(longUrl string, shortUrl string) bool
	LogAllUrls()
}
type ShortUrlRepositoryImpl struct {
	longUrlToShortUrlMap LongUrlToShortUrlMap
	shortUrlToLongUrlMap ShortUrlToLongUrlMap
	mutex                *sync.RWMutex
}

func NewShortUrlRepositoryImpl(longUrlToShortUrlMap LongUrlToShortUrlMap, shortUrlToLongUrlMap ShortUrlToLongUrlMap, mutex *sync.RWMutex) *ShortUrlRepositoryImpl {
	return &ShortUrlRepositoryImpl{
		longUrlToShortUrlMap: longUrlToShortUrlMap,
		shortUrlToLongUrlMap: shortUrlToLongUrlMap,
		mutex:                mutex,
	}
}

func (impl *ShortUrlRepositoryImpl) GetShortUrlByLongUrl(longUrl string) (string, bool) {
	impl.mutex.RLock()
	shortUrl, ok := impl.longUrlToShortUrlMap[longUrl]
	impl.mutex.RUnlock()
	return shortUrl, ok
}

func (impl *ShortUrlRepositoryImpl) GetLongUrlByShortUrl(shortUrl string) (string, bool) {
	impl.mutex.RLock()
	longUrl, ok := impl.shortUrlToLongUrlMap[shortUrl]
	impl.mutex.RUnlock()
	return longUrl, ok
}

func (impl *ShortUrlRepositoryImpl) AddUrl(longUrl string, shortUrl string) bool {
	impl.mutex.Lock()
	impl.longUrlToShortUrlMap[longUrl] = shortUrl
	impl.shortUrlToLongUrlMap[shortUrl] = longUrl
	impl.mutex.Unlock()
	return true
}

func (impl *ShortUrlRepositoryImpl) LogAllUrls() {
	for longUrl, shortUrl := range impl.longUrlToShortUrlMap {
		fmt.Println(longUrl, shortUrl)
	}
	fmt.Println(len(impl.longUrlToShortUrlMap), len(impl.shortUrlToLongUrlMap))
}
