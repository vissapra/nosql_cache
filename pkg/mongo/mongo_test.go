package mongo

import (
	"log"
	"testing"
)

func TestCollectionClient_Put(t *testing.T) {
	mongoCache, e := NewMongoCache(MongoConfig{Addrs: []string{"localhost:27017"}})
	if e != nil {
		t.Error(e)
	}
	arsReqClient := NewCollectionClient("super_session", "ars_request", mongoCache)
	success, err := arsReqClient.Put("208", []byte("resolveAvailability"))
	if err != nil {
		t.Fatal(err)
	}
	if success {
		if response, err := arsReqClient.Get("208"); err != nil {
			t.Error(err)

		} else {
			log.Println("Response;", string(response))
		}

	}
}
