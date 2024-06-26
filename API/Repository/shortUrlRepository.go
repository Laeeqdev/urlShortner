package repository

import (
	"fmt"
	"sync"
)

type ShortUrlRepository interface {
	GetShortUrlByLongUrl(longUrl string) (string, bool)
	GetLongUrlByShortUrl(shortUrl string) (string, bool)
	AddUrl(longUrl string, shortUrl string) bool
	LogAllUrls()
}
type ShortUrlRepositoryImpl struct {
	longUrlToShortUrlMap map[string]string
	shortUrlToLongUrlMap map[string]string
	mutex                *sync.RWMutex
}

func NewShortUrlRepositoryImpl(longUrlToShortUrlMap map[string]string, shortUrlToLongUrlMap map[string]string, mutex *sync.RWMutex) *ShortUrlRepositoryImpl {
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
