package cache

import (
	"time"

	"github.com/levigross/grequests"
)

type Cache struct {
	Response     *grequests.Response
	ResponseBody []byte
	CreatedAt    time.Time
}

func NewCache(response *grequests.Response, body []byte, cAt time.Time) *Cache {
	return &Cache{
		Response:     response,
		ResponseBody: body,
		CreatedAt:    cAt,
	}
}
