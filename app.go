//go:build wireinject
// +build wireinject

// wire.go
package main

import (
	repository "github.com/Laeeqdev/urlShortner/API/Repository"
	resthandler "github.com/Laeeqdev/urlShortner/API/RestHandler"
	router "github.com/Laeeqdev/urlShortner/API/Router"
	service "github.com/Laeeqdev/urlShortner/API/Services"
	"github.com/google/wire"
)

func InitializeApp() *router.RouterImpl {
	wire.Build(repository.NewShortUrlRepositoryImpl, wire.Bind(new(repository.ShortUrlRepository), new(*repository.ShortUrlRepositoryImpl)),
		service.NewShortUrlServiceImpl, wire.Bind(new(service.ShortUrlService), new(*service.ShortUrlServiceImpl)),
		resthandler.NewShortUrlHandlerImpl, wire.Bind(new(resthandler.ShortUrlHandler), new(*resthandler.ShortUrlHandlerImpl)),
		router.NewRouterImpl,
		repository.ProvideLongUrlToShortUrlMap,
		repository.ProvideMutex,
		repository.ProvideShortUrlToLongUrlMap,
	)
	return &router.RouterImpl{}
}
