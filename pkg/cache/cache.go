package cache

type Cache interface {
	Get(key string) (response []byte, err error)
	Put(key string, value []byte) (success bool, err error)
	Exists(key string) (found bool, err error)
	PutWithExpiry(key string, value []byte, expiry int32) (success bool, err error)
}
