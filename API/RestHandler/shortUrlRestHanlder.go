package resthandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	constants "github.com/Laeeqdev/urlShortner/API/Constants"
	models "github.com/Laeeqdev/urlShortner/API/Models"
	service "github.com/Laeeqdev/urlShortner/API/Services"
	"github.com/gorilla/mux"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	domain string
)

func init() {
	_, ok := os.LookupEnv(constants.APP_DOMAIN)
	if !ok {
		domain = constants.DOMAIN
		return
	}
	domain = os.Getenv(constants.APP_DOMAIN)
}

type ShortUrlHandler interface {
	ShortenUrl(w http.ResponseWriter, r *http.Request)
	LengthenUrl(w http.ResponseWriter, r *http.Request)
	Redirect(w http.ResponseWriter, r *http.Request)
	LogAllUrls(w http.ResponseWriter, r *http.Request)
	parseAndValidateRequest(data interface{}, w http.ResponseWriter, r *http.Request) bool
	GetTopDomains(w http.ResponseWriter, r *http.Request)
}
type ShortUrlHandlerImpl struct {
	shortUrlService service.ShortUrlService
}

func NewShortUrlHandlerImpl(shortUrlService service.ShortUrlService) *ShortUrlHandlerImpl {
	return &ShortUrlHandlerImpl{shortUrlService: shortUrlService}
}

func (impl *ShortUrlHandlerImpl) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var requestBody models.ShortUrlRequest

	if ok := impl.parseAndValidateRequest(&requestBody, w, r); !ok {
		return
	}

	shortUrl, err := impl.shortUrlService.GenerateShortUrl(requestBody.LongUrl)
	shortUrl = domain + shortUrl
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	urlResponse := models.UrlResponse{LongUrl: requestBody.LongUrl, ShortUrl: shortUrl}
	response, _ := json.Marshal(urlResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// handler for POST /lengthen
func (impl *ShortUrlHandlerImpl) LengthenUrl(w http.ResponseWriter, r *http.Request) {
	var requestBody models.LongUrlRequest
	if ok := impl.parseAndValidateRequest(&requestBody, w, r); !ok {
		return
	}

	longUrl, err := impl.shortUrlService.GetLongUrlByShortUrl(requestBody.ShortUrl)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	urlResponse := models.UrlResponse{LongUrl: longUrl, ShortUrl: requestBody.ShortUrl}
	response, _ := json.Marshal(urlResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (impl *ShortUrlHandlerImpl) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars[constants.SHORT_URL]
	longUrl, err := impl.shortUrlService.GetLongUrlByShortUrl(shortUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, longUrl, http.StatusSeeOther)
}

func (impl *ShortUrlHandlerImpl) LogAllUrls(w http.ResponseWriter, r *http.Request) {
	impl.shortUrlService.LogAllUrls()
}

func (impl *ShortUrlHandlerImpl) parseAndValidateRequest(requestBody interface{}, w http.ResponseWriter, r *http.Request) bool {
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println("error while parsing request data err : ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	validateStruct := validator.New()
	err = validateStruct.Struct(requestBody)
	if err != nil {
		fmt.Println("error while validating struct err : ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func (impl *ShortUrlHandlerImpl) GetTopDomains(w http.ResponseWriter, r *http.Request) {
	topDomains := impl.shortUrlService.GetTopDomains(3)
	json.NewEncoder(w).Encode(topDomains)
}
