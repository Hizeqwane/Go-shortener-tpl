package service

import "github.com/Hizeqwane/Go-shortener-tpl/internal/service/generator"

type IShortenerService interface {
	GetShortUri(string) string
	TryGetLongUri(string) (string, bool)
}

func NewShortenerService() IShortenerService {
	return &shortenerService{
		gen:      generator.ShortUriGenerator{},
		shortMap: make(map[string]string),
		longMap:  make(map[string]string),
	}
}

type shortenerService struct {
	gen      generator.ShortUriGenerator
	shortMap map[string]string
	longMap  map[string]string
}

func (service *shortenerService) GetShortUri(longUri string) string {
	if i, ok := service.longMap[longUri]; ok {
		return i
	}

	shortUri := service.gen.GenerateShortUri(longUri)
	service.shortMap[shortUri] = longUri
	service.longMap[longUri] = shortUri

	return shortUri
}

func (service *shortenerService) TryGetLongUri(shortUri string) (string, bool) {
	i, ok := service.shortMap[shortUri]

	return i, ok
}

var _ IShortenerService = (*shortenerService)(nil)
