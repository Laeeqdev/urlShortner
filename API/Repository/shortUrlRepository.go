package repository

import (
	"fmt"
	"sort"
	"sync"

	models "github.com/Laeeqdev/urlShortner/API/Models"
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
	IncrementDomainCount(domain string)
	GetTopDomains(n int) []models.DomainCount
	LogAllUrls()
}
type ShortUrlRepositoryImpl struct {
	longUrlToShortUrlMap LongUrlToShortUrlMap
	shortUrlToLongUrlMap ShortUrlToLongUrlMap
	domainStats          map[string]int
	mutex                *sync.RWMutex
}

func NewShortUrlRepositoryImpl(longUrlToShortUrlMap LongUrlToShortUrlMap, shortUrlToLongUrlMap ShortUrlToLongUrlMap, mutex *sync.RWMutex) *ShortUrlRepositoryImpl {
	return &ShortUrlRepositoryImpl{
		longUrlToShortUrlMap: longUrlToShortUrlMap,
		shortUrlToLongUrlMap: shortUrlToLongUrlMap,
		domainStats:          make(map[string]int),
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
		fmt.Printf("long url : %s short url : %s \n", longUrl, shortUrl)
	}
	fmt.Println(len(impl.longUrlToShortUrlMap), len(impl.shortUrlToLongUrlMap))
}

func (impl *ShortUrlRepositoryImpl) IncrementDomainCount(domain string) {
	impl.mutex.Lock()
	defer impl.mutex.Unlock()
	impl.domainStats[domain]++
}

func (impl *ShortUrlRepositoryImpl) GetTopDomains(n int) []models.DomainCount {
	impl.mutex.RLock()
	defer impl.mutex.RUnlock()

	var domainCounts []models.DomainCount
	for domain, count := range impl.domainStats {
		domainCounts = append(domainCounts, models.DomainCount{Domain: domain, Count: count})
	}

	sort.Slice(domainCounts, func(i, j int) bool {
		return domainCounts[i].Count > domainCounts[j].Count
	})

	topDomains := make([]models.DomainCount, 0)
	for i := 0; i < n && i < len(domainCounts); i++ {
		topDomains = append(topDomains, models.DomainCount{Domain: domainCounts[i].Domain, Count: domainCounts[i].Count})
	}
	return topDomains
}
