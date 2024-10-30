package cache

import (
	"github.com/showyquasar88/proj-combine/cache_demo/utils"
	"log"
	"sync"
	"time"
)

type Cache interface {
	SetMaxMemory(size string) bool // size: 1KB 100KB 1MB 2MB 1GB
	Set(key string, val interface{}, expire time.Duration) bool
	Get(key string) (interface{}, bool)
	Del(key string) bool
	Exists(key string) bool
	Flush() bool
	Keys() int64 // 获取缓存中的key
}

type MemCache struct {
	maxMemorySize    int64
	maxMemorySizeStr string // 冗余字段，因为设置max的时候用的是字符串
	curMemorySize    int64
	val              map[string]valWithExpire // 存储缓存的值
	lock             sync.RWMutex
	cleanInterval    time.Duration
}

type valWithExpire struct {
	val        interface{}
	expireTime *time.Time
	size       int64
}

func NewMemCache() Cache {
	result := &MemCache{
		val:           make(map[string]valWithExpire),
		cleanInterval: 5 * time.Minute,
	}
	go result.Clean()
	return result
}

func (m *MemCache) SetMaxMemory(size string) bool {
	m.maxMemorySizeStr, m.maxMemorySize = utils.ParseSize(size)
	log.Printf("set success: maxMemorySize %d\r\n", m.maxMemorySize)
	return true
}

func (m *MemCache) Set(key string, val interface{}, expire time.Duration) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	v := valWithExpire{
		val:        val,
		expireTime: nil,
		size:       utils.GetValSize(val),
	}
	if expire != 0 {
		*(v.expireTime) = time.Now().Add(expire)
	}
	existsVal, ok := m.val[key]
	if ok {
		if m.curMemorySize-existsVal.size+v.size > m.maxMemorySize {
			return false
		}
		m.Del(key)
		m.curMemorySize -= existsVal.size
	} else {
		if m.curMemorySize+v.size > m.maxMemorySize {
			return false
		}
	}
	m.val[key] = v
	m.curMemorySize += v.size
	return true
}

func (m *MemCache) Get(key string) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	existsVal, ok := m.val[key]
	if !ok {
		return nil, false
	}
	if existsVal.expireTime != nil && existsVal.expireTime.Before(time.Now()) {
		m.Del(key)
		return nil, false
	}
	return existsVal, true
}

func (m *MemCache) Del(key string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	existsVal, ok := m.val[key]
	if !ok {
		return true
	}
	delete(m.val, key)
	m.curMemorySize -= existsVal.size
	return true
}

func (m *MemCache) Exists(key string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.val[key]
	return ok
}

func (m *MemCache) Flush() bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.curMemorySize = 0
	m.val = make(map[string]valWithExpire)
	return true
}

func (m *MemCache) Keys() int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return int64(len(m.val))
}

func (m *MemCache) Clean() {
	ticker := time.NewTicker(m.cleanInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for key, val := range m.val {
				if val.expireTime != nil && val.expireTime.Before(time.Now()) {
					m.Del(key)
				}
			}
		}
	}
}
