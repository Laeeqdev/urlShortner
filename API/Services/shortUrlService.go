package service

import (
	"fmt"
	constants "github.com/Laeeqdev/urlShortner/API/Constants"
	repository "github.com/Laeeqdev/urlShortner/API/Repository"
)

type ShortUrlService interface {
	GetShortUrlByLongUrl(longUrl string) (string, bool)
	GetLongUrlByShortUrl(shortUrl string) (string, bool)
	AddUrl(longUrl string, shortUrl string) (string, bool)
	LogAllUrls()
	CheckIfShortUrlExists(shortUrl string) (string, bool)
	CheckIfLongUrlExists(shortUrl string) (string, bool)
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
	return shortUrl, nil
}

func (impl *ShortUrlServiceImpl) AddUrl(longUrl string, shortUrl string) error {
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

func (impl *ShortUrlServiceImpl) CheckIfShortUrlExists(shortUrl string) (string, error) {
	return impl.GetShortUrlByLongUrl(shortUrl)
}

func (impl *ShortUrlServiceImpl) CheckIfLongUrlExists(longUrl string) (string, error) {
	return impl.GetLongUrlByShortUrl(longUrl)
}
