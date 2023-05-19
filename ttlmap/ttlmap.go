package ttlmap

import (
	"log"
	"runtime"
	"sync"
	"time"
)

const (
	KEY_LEN = 32
)

type item struct {
	value  interface{}
	expire int64
}

type TTLMapCore struct {
	ticker *time.Ticker
	lock   *sync.RWMutex
	stop   chan int
	m      map[[KEY_LEN]byte]item
}

func (t *TTLMapCore) GetWithTime(key [KEY_LEN]byte, now time.Time) (interface{}, bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	item, ok := t.m[key]
	if !ok || now.UnixMilli() > item.expire {
		return nil, false
	}
	return item.value, true
}

func (t *TTLMapCore) Get(key [KEY_LEN]byte) (interface{}, bool) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	item, ok := t.m[key]
	if !ok || time.Now().UnixMilli() > item.expire {
		return nil, false
	}
	return item.value, true
}

func (t *TTLMapCore) SetWithTime(key [KEY_LEN]byte, value interface{}, now time.Time, exp time.Duration) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.m[key] = item{value: value, expire: now.Add(exp).UnixMilli()}
}

func (t *TTLMapCore) Set(key [KEY_LEN]byte, value interface{}, exp time.Duration) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.m[key] = item{value: value, expire: time.Now().Add(exp).UnixMilli()}
}

func (t *TTLMapCore) SetByUnSafeBytesKey(key []byte, value interface{}, exp time.Duration) {
	var _key [32]byte
	copy(_key[:], key[:32])
	t.lock.Lock()
	defer t.lock.Unlock()
	t.m[_key] = item{value: value, expire: time.Now().Add(exp).UnixMilli()}
}

func (t *TTLMapCore) GetByUnSafeBytesKey(key []byte) (interface{}, bool) {
	var _key [32]byte
	copy(_key[:], key[:32])
	t.lock.RLock()
	defer t.lock.RUnlock()
	item, ok := t.m[_key]
	if !ok || time.Now().UnixMilli() > item.expire {
		return nil, false
	}
	return item.value, true
}

func (t *TTLMapCore) clean() {
	defer t.ticker.Stop()
	for range t.ticker.C {
		select {
		case <-t.stop:
			log.Println("exit")
			return
		default:
			t.lock.Lock()
			timeNow := time.Now().UnixMilli()
			for key, value := range t.m {
				if timeNow >= value.expire {
					delete(t.m, key)
				}
			}
			t.lock.Unlock()
		}

	}
}

type TTLMap struct {
	*TTLMapCore
}

func NewTTLMap(cleanTick time.Duration) *TTLMap {
	ttl_map := &TTLMap{TTLMapCore: &TTLMapCore{m: make(map[[KEY_LEN]byte]item), lock: &sync.RWMutex{}, stop: make(chan int, 1), ticker: time.NewTicker(time.Second * 1)}}
	go ttl_map.clean()
	runtime.SetFinalizer(ttl_map, func(t *TTLMap) {
		log.Println("On GC")
		t.stop <- 1
	})
	return ttl_map
}
