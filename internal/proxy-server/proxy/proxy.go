package proxy

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Venukishore-R/caching-proxy/internal/proxy-server/cache"
	"github.com/levigross/grequests"
)

type Proxy struct {
	Origin string
	Cache  map[string]*cache.Cache
	M      sync.RWMutex // Use RWMutex for better concurrency
}

func NewProxy(origin string, clear bool) *Proxy {
	return &Proxy{
		Origin: origin,
		Cache:  make(map[string]*cache.Cache),
		M:      sync.RWMutex{},
	}
}

func (p *Proxy) ClearCache() {
	p.M.Lock() // Lock for writing
	defer p.M.Unlock()
	p.Cache = make(map[string]*cache.Cache)
	fmt.Println("Cache cleared successfully")
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path

	if r.URL.Path == "/clear-cache" {
		fmt.Println("Clearing cache...")
		p.ClearCache()
		w.Write([]byte("Cache cleared successfully"))
		return
	}

	CACHE_KEY := r.Method + ":" + urlPath

	p.M.RLock() // Lock for reading
	c, ok := p.Cache[CACHE_KEY]
	p.M.RUnlock()

	if ok {
		WriteResponseFromOrigin(w, c.Response, c.ResponseBody, "HIT", CACHE_KEY)
		return
	}

	response, err := grequests.Get(p.Origin+urlPath, nil)
	if err != nil {
		http.Error(w, "Error fetching from origin: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, "Error from origin", response.StatusCode)
		return
	}

	body := response.Bytes()

	cachedResponse := cache.NewCache(response, body, time.Now())

	p.M.Lock()
	p.Cache[CACHE_KEY] = cachedResponse
	p.M.Unlock()

	WriteResponseFromOrigin(w, response, body, "MISS", CACHE_KEY)
}

func WriteResponseFromOrigin(w http.ResponseWriter, resp *grequests.Response, body []byte, cacheHeader, key string) {
	fmt.Printf("X-Cache: %s %s\n", cacheHeader, key)
	w.WriteHeader(resp.StatusCode)
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.Write(body)
}
