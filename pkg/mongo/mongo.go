package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Mongo       struct{}
	MongoConfig struct {
		Addrs              []string
		Timeout            time.Duration
		Username, Password string
	}
	CollectionClient struct {
		db, collection string
		*mgo.Session
	}
)

func NewMongoCache(mConfig MongoConfig) (*mgo.Session, error) {
	info := &mgo.DialInfo{
		Addrs:    mConfig.Addrs,
		Timeout:  60 * time.Second,
		Username: mConfig.Username,
		Password: mConfig.Password,
	}
	return mgo.DialWithInfo(info)
}

func NewCollectionClient(db, collection string, session *mgo.Session) *CollectionClient {
	return &CollectionClient{db: db, collection: collection, Session: session}
}
func (m *CollectionClient) Get(key string) (response []byte, err error) {
	session := m.Session.Copy()
	defer session.Close()
	collection := session.DB(m.db).C(m.collection)
	content := struct {
		Id       string
		Response []byte
	}{}
	err = collection.Find(bson.M{"id": key}).One(&content)
	if err != nil {
		return nil, err
	}
	return content.Response, nil
}

/*
 *
 */
func (m *CollectionClient) Put(key string, value []byte) (success bool, err error) {
	session := m.Session.Copy()
	defer session.Close()
	collection := session.DB(m.db).C(m.collection)
	_, err = collection.Upsert(bson.M{"id": key}, bson.M{"id": key, "response": value})
	return err == nil, err
}
func (m *CollectionClient) Exists(key string) (found bool, err error) {
	session := m.Session.Copy()
	defer session.Close()
	collection := session.DB(m.db).C(m.collection)
	if n, err := collection.Find(bson.M{"id": key}).Count(); n > 0 {
		return true, nil
	} else {
		return false, err
	}
}
func (m *CollectionClient) PutWithExpiry(key string, value []byte, expiry int32) (success bool, err error) {
	return true, nil
}
