// wire.go

package main

import (
	"sync"

	repository "github.com/Laeeqdev/urlShortner/API/Repository"
	resthandler "github.com/Laeeqdev/urlShortner/API/RestHandler"
	router "github.com/Laeeqdev/urlShortner/API/Router"
	service "github.com/Laeeqdev/urlShortner/API/Services"
	"github.com/google/wire"
)

func InitializeApp(map1 map[string]string, map2 map[string]string, mutex *sync.RWMutex) *router.RouterImpl {
	wire.Build(repository.NewShortUrlRepositoryImpl, wire.Bind(new(repository.ShortUrlRepository), new(*repository.ShortUrlRepositoryImpl)),
		service.NewShortUrlServiceImpl, wire.Bind(new(service.ShortUrlService), new(*service.ShortUrlServiceImpl)),
		resthandler.NewShortUrlHandlerImpl, wire.Bind(new(resthandler.ShortUrlHandler), new(*resthandler.ShortUrlHandlerImpl)),
		router.NewRouterImpl,
	)
	return &router.RouterImpl{}
}
