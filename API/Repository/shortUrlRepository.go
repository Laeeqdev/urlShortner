package repository

import (
	"fmt"
	"sync"
)

type ShortUrlRepository interface {
	GetShortUrlByLongUrl(longUrl string) (string, bool)
	GetLongUrlByShortUrl(shortUrl string) (string, bool)
	AddUrl(longUrl string, shortUrl string) (string, bool)
	LogAllUrls()
}
type ShortUrlRepositoryImpl struct {
	LongUrlToShortUrlMap map[string]string
	ShortUrlToLongUrlMap map[string]string
	Mutexx               *sync.RWMutex
}

func NewShortUrlRepositoryImpl(LongUrlToShortUrlMap map[string]string, ShortUrlToLongUrlMap map[string]string, Mutexx *sync.RWMutex) *ShortUrlRepositoryImpl {
	return &ShortUrlRepositoryImpl{
		LongUrlToShortUrlMap: LongUrlToShortUrlMap,
		ShortUrlToLongUrlMap: ShortUrlToLongUrlMap,
		Mutexx:               Mutexx,
	}
}

func (impl *ShortUrlRepositoryImpl) GetShortUrlByLongUrl(longUrl string) (string, bool) {
	impl.Mutexx.RLock()
	shortUrl, ok := impl.LongUrlToShortUrlMap[longUrl]
	impl.Mutexx.RUnlock()
	return shortUrl, ok
}

func (impl *ShortUrlRepositoryImpl) GetLongUrlByShortUrl(shortUrl string) (string, bool) {
	impl.Mutexx.RLock()
	longUrl, ok := impl.ShortUrlToLongUrlMap[shortUrl]
	impl.Mutexx.RUnlock()
	return longUrl, ok
}

func (impl *ShortUrlRepositoryImpl) AddUrl(longUrl string, shortUrl string) bool {
	impl.Mutexx.Lock()
	impl.LongUrlToShortUrlMap[longUrl] = shortUrl
	impl.ShortUrlToLongUrlMap[shortUrl] = longUrl
	impl.Mutexx.Unlock()
	return true
}

func (impl *ShortUrlRepositoryImpl) LogAllUrls() {
	for longUrl, shortUrl := range impl.LongUrlToShortUrlMap {
		fmt.Println(longUrl, shortUrl)
	}
	fmt.Println(len(impl.LongUrlToShortUrlMap), len(impl.ShortUrlToLongUrlMap))
}
