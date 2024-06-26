package router

import (
	resthandler "github.com/Laeeqdev/urlShortner/API/RestHandler"
	"github.com/gorilla/mux"
)

type Router interface {
	MyRouter() *mux.Router
}
type RouterImpl struct {
	shortUrlHandler resthandler.ShortUrlHandler
}

func NewRouterImpl(shortUrlHandler resthandler.ShortUrlHandler) *RouterImpl {
	return &RouterImpl{shortUrlHandler: shortUrlHandler}
}

func (impl *RouterImpl) MyRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", impl.shortUrlHandler.ShortenUrl).Methods("POST")
	r.HandleFunc("/lengthen", impl.shortUrlHandler.LengthenUrl).Methods("POST")
	r.HandleFunc("/{shortUrl}", impl.shortUrlHandler.Redirect).Methods("GET")
	r.HandleFunc("/", impl.shortUrlHandler.LogAllUrls).Methods("GET")
	r.HandleFunc("/metrics", impl.shortUrlHandler.GetTopDomains).Methods("POST")
	return r
}
