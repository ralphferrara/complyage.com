package keys

import (
	"api/db"
	"api/models"
	"fmt"
	"sync"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Concurrent safe
//||------------------------------------------------------------------------------------------------||

var (
	keys      = make(map[string]models.API)
	keysMutex sync.RWMutex
)

//||------------------------------------------------------------------------------------------------||
//|| Load Keys
//||------------------------------------------------------------------------------------------------||

func LoadKeys() error {
	var apis []models.API
	if err := db.DB.Find(&apis).Error; err != nil {
		return err
	}

	temp := make(map[string]models.API)
	for _, api := range apis {
		temp[api.APIPublic] = api
	}

	keysMutex.Lock()
	keys = temp
	keysMutex.Unlock()

	fmt.Printf("Loaded %d API keys into memory\n", len(keys))
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Lookup Keys
//||------------------------------------------------------------------------------------------------||

func GetKey(apiPublic string) (models.API, bool) {
	keysMutex.RLock()
	defer keysMutex.RUnlock()

	api, ok := keys[apiPublic]
	return api, ok
}

//||------------------------------------------------------------------------------------------------||
//|| Start AutoReload
//||------------------------------------------------------------------------------------------------||

func StartAutoReload(interval time.Duration) {
	if err := LoadKeys(); err != nil {
		fmt.Println("Failed to load keys:", err)
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			if err := LoadKeys(); err != nil {
				fmt.Println("Failed to reload keys:", err)
			}
		}
	}()
}
