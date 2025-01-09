package cache

// import (
// 	"sync"

// 	"github.com/sirupsen/logrus"
// )

// type Iterator <-chan string

// type OrderedCache interface {
// 	add(key string, data interface{})
// }

// type Cache interface {
// 	Save(key string, nestedKey string, data interface{})
// }

// type cache struct {
// 	sync.RWMutex
// 	logger       *logrus.Entry
// 	data         map[string]OrderedCache
// 	isOrderedMap bool
// }

// func NewCache(isOrdered ...bool) Cache {
// 	isOrderedMap := true
// 	if len(isOrdered) > 0 {
// 		isOrderedMap = isOrdered[0]
// 	}
// 	logger := logrus.New()
// 	logger.WithField("cache", "local")
// 	data := make(map[string]OrderedCache)
// 	return &cache{logger: logrus.NewEntry(logger), data: data, isOrderedMap: isOrderedMap}
// }

// func (c *cache) Save(key, nestedKey string, data interface{}) {
// }
