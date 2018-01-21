//Hazelcast client wrapper to access maps using memcached client
package hazelcast

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

type HzClient struct {
	client *memcache.Client
}

func NewHzClient(address ...string) *HzClient {
	mc := memcache.New(address...)
	return &HzClient{client: mc}
}

// Puts the key in Hazelcast map with expiry in Seconds
func (c *HzClient) PutWithExpiry(mapName string, key string, value []byte, expiry int32) (bool, error) {
	return c.put(&memcache.Item{Key: mapName + ":" + key, Value: value, Expiration: expiry})
}

func (c *HzClient) Put(mapName string, key string, value []byte) (bool, error) {
	return c.put(&memcache.Item{Key: mapName + ":" + key, Value: value})

}

func (c *HzClient) put(item *memcache.Item) (bool, error) {
	err := c.client.Set(item)
	if err != nil {
		log.Printf("error storing key: %s, err: %s", item.Key, err)
		return false, err
	}
	return true, nil
}

func (c *HzClient) Get(mapName string, key string) ([]byte, error) {
	value, err := c.client.Get(mapName + ":" + key)
	if err != nil {
		log.Printf("error retrieving key: %s, err: %s", key, err)
		return nil, err
	}
	return value.Value, nil

}
