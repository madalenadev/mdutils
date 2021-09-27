package cache

import (
	"context"
	"math/rand"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type SetItemFn func(item interface{})
type GetItemFn func(key string) interface{}

type ObjDTO struct {
	Key  string
	Val1 int
	Val2 int
	Val3 int
}

var ctx = context.Background()
var payloads = make(chan interface{})
var keys = make([]string, 50)

func generateRandom(min int, max int) int {
	return rand.Intn(max-min) + min
}

func init() {
	for idx := 0; idx < 50; idx++ {
		keys[idx] = uuid.New().String()
	}
}

func runTest(iterations int, setItemFn SetItemFn, getItemFn GetItemFn) {
	// Initialize data generator
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			payloads <- &ObjDTO{
				Key:  keys[generateRandom(0, 50)],
				Val1: rand.Int(),
				Val2: rand.Int(),
				Val3: rand.Int(),
			}
		}
	}(ctx)
	// /Initialize data generator

	wgWrite := &sync.WaitGroup{}
	wgRead := &sync.WaitGroup{}

	for i := 0; i < iterations; i++ {

		wgWrite.Add(1)
		go func(wg *sync.WaitGroup, idx int) {
			defer wg.Done()

			payload := <-payloads
			setItemFn(payload)
		}(wgWrite, i)

		wgRead.Add(1)
		go func(wg *sync.WaitGroup, idx int) {
			defer wg.Done()

			key := keys[generateRandom(0, 50)]
			getItemFn(key)
		}(wgRead, i)
	}

	wgRead.Wait()
	wgWrite.Wait()
	cancel()
}

func BenchmarkOnRedis(b *testing.B) {
	cacheImpl := NewRedis(Options{
		URL: "localhost:6379",
	})

	setItem := func(obj interface{}) {
		objDTO := obj.(*ObjDTO)
		cacheImpl.Set(ctx, objDTO.Key, objDTO)
	}

	getItem := func(key string) interface{} {
		objDTO := &ObjDTO{}
		err := cacheImpl.Get(ctx, key, objDTO)

		if err != nil {
			panic(err.Error())
		}

		return objDTO
	}

	runTest(b.N, setItem, getItem)
}

func BenchmarkOnMemcached(b *testing.B) {
	cacheImpl := NewMemcache(Options{
		URL: "localhost:11211",
	})

	setItem := func(obj interface{}) {
		objDTO := obj.(*ObjDTO)
		cacheImpl.Set(ctx, objDTO.Key, objDTO)
	}

	getItem := func(key string) interface{} {
		objDTO := &ObjDTO{}
		err := cacheImpl.Get(ctx, key, objDTO)

		if err != nil {
			panic(err.Error())
		}

		return objDTO
	}

	runTest(b.N, setItem, getItem)
}

func TestRedisFlow(t *testing.T) {
	redisCache := NewRedis(Options{
		URL: "localhost:6379",
	})

	objDTO := &ObjDTO{
		Val1: 1,
		Val2: 2,
		Val3: 3,
		Key:  "Key",
	}

	err := redisCache.Set(ctx, "teste1", objDTO)
	if err != nil {
		panic(err.Error())
	}

	newObj := ObjDTO{}

	err = redisCache.Get(ctx, "teste1", &newObj)

	if err != nil {
		panic(err.Error())
	}

	assert.Equal(t, objDTO.Key, newObj.Key)
	assert.Equal(t, objDTO.Val1, newObj.Val1)
	assert.Equal(t, objDTO.Val2, newObj.Val2)
	assert.Equal(t, objDTO.Val3, newObj.Val3)
}

func TestMemcacheFlow(t *testing.T) {
	memcachedCache := NewMemcache(Options{
		URL: "localhost:11211",
	})

	objDTO := &ObjDTO{
		Val1: 1,
		Val2: 2,
		Val3: 3,
		Key:  "Key",
	}

	err := memcachedCache.Set(ctx, "teste1", objDTO)
	if err != nil {
		panic(err.Error())
	}

	newObj := ObjDTO{}

	err = memcachedCache.Get(ctx, "teste1", &newObj)

	if err != nil {
		panic(err.Error())
	}

	assert.Equal(t, objDTO.Key, newObj.Key)
	assert.Equal(t, objDTO.Val1, newObj.Val1)
	assert.Equal(t, objDTO.Val2, newObj.Val2)
	assert.Equal(t, objDTO.Val3, newObj.Val3)
}
