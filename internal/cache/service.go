package cache

import (
	"log"

	"github.com/dgraph-io/ristretto"
)

func NewCacheService() ICache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000, // For tracking access frequency
		MaxCost:     1024, // Max size of cache
		BufferItems: 64,   // How many keys we process at once
	})
	if err != nil {
		log.Fatal(err)
	}
	return &SCache{
		cache: cache,
	}
}
