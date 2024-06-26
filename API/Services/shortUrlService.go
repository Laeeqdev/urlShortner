package service

import (
	"fmt"

	constants "github.com/Laeeqdev/urlShortner/API/Constants"
	models "github.com/Laeeqdev/urlShortner/API/Models"
	repository "github.com/Laeeqdev/urlShortner/API/Repository"
	utils "github.com/Laeeqdev/urlShortner/API/Utils"
)

type ShortUrlService interface {
	GetShortUrlByLongUrl(longUrl string) (string, error)
	GetLongUrlByShortUrl(shortUrl string) (string, error)
	AddUrl(longUrl string, shortUrl string) error
	LogAllUrls()
	CheckIfShortUrlExists(shortUrl string) (string, bool)
	CheckIfLongUrlExists(shortUrl string) (string, bool)
	GenerateShortUrl(longUrl string) (string, error)
	GetTopDomains(n int) []models.DomainCount
}

type ShortUrlServiceImpl struct {
	shortUrlRepository repository.ShortUrlRepository
}

func NewShortUrlServiceImpl(shortUrlRepository repository.ShortUrlRepository) *ShortUrlServiceImpl {
	return &ShortUrlServiceImpl{shortUrlRepository: shortUrlRepository}
}
func (impl *ShortUrlServiceImpl) GetShortUrlByLongUrl(longUrl string) (string, error) {
	shortUrl, ok := impl.shortUrlRepository.GetShortUrlByLongUrl(longUrl)
	if !ok {
		fmt.Printf("no short url found for long url : %s  \n", longUrl)
		return constants.EMPTY_STRING, fmt.Errorf(constants.NO_SHORT_URL_FOUND_FOR_GIVEN_URL)
	}
	fmt.Printf("successfully fetched short url. long url : %s short url : %s \n", longUrl, shortUrl)
	return shortUrl, nil
}

func (impl *ShortUrlServiceImpl) GetLongUrlByShortUrl(shortUrl string) (string, error) {
	longUrl, ok := impl.shortUrlRepository.GetLongUrlByShortUrl(shortUrl)
	if !ok {
		fmt.Printf("no long url found for short url : %s  \n", shortUrl)
		return constants.EMPTY_STRING, fmt.Errorf(constants.NO_LONG_URL_FOUND_FOR_GIVEN_URL)
	}
	fmt.Printf("successfully fetched long url. short url : %s long url : %s \n", shortUrl, longUrl)
	return longUrl, nil
}

func (impl *ShortUrlServiceImpl) AddUrl(longUrl string, shortUrl string) error {
	domain, err := utils.ExtractDomain(longUrl)
	if err != nil {
		fmt.Printf("Error extracting domain from URL %s: %v\n", longUrl, err)
	}
	impl.shortUrlRepository.IncrementDomainCount(domain)
	ok := impl.shortUrlRepository.AddUrl(longUrl, shortUrl)
	if !ok {
		fmt.Printf("unable to add urls short url : %s long url : %s \n", shortUrl, longUrl)
		return fmt.Errorf(constants.UNABLE_TO_ADD_URL)
	}
	fmt.Printf("urls added successfully short url : %s long url : %s \n", shortUrl, longUrl)
	return nil
}

func (impl *ShortUrlServiceImpl) LogAllUrls() {
	impl.shortUrlRepository.LogAllUrls()
}

func (impl *ShortUrlServiceImpl) CheckIfShortUrlExists(shortUrl string) (string, bool) {
	longUrl, err := impl.GetLongUrlByShortUrl(shortUrl)
	return longUrl, err == nil
}

func (impl *ShortUrlServiceImpl) CheckIfLongUrlExists(longUrl string) (string, bool) {
	shortUrl, err := impl.GetShortUrlByLongUrl(longUrl)
	return shortUrl, err == nil
}

func (impl *ShortUrlServiceImpl) GenerateShortUrl(longUrl string) (string, error) {
	// Check if long url already in datastore
	if shortUrl, ok := impl.CheckIfLongUrlExists(longUrl); ok {
		return shortUrl, nil
	}

	// limit max retries to 5
	retryLimit := 0
RETRY_DUE_TO_COLLISION:
	if retryLimit += 1; retryLimit <= 5 {
		// Genrating ShortUrl
		shortUrl := utils.GetRandomShortUrl()
		//3) Check for collision
		if _, ok := impl.CheckIfShortUrlExists(shortUrl); ok {
			goto RETRY_DUE_TO_COLLISION
		} else {
			// 4) insert url into map
			if err := impl.AddUrl(longUrl, shortUrl); err == nil {
				return shortUrl, nil
			} else {
				panic(constants.SOMETHING_WENT_WRONG)
			}
		}
	}
	return constants.EMPTY_STRING, fmt.Errorf(constants.RETRY_LIMIT_EXCEEDED)
}

func (impl *ShortUrlServiceImpl) GetTopDomains(n int) []models.DomainCount {
	return impl.shortUrlRepository.GetTopDomains(n)
}
