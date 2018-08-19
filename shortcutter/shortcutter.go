package shortcutter

import (
	"log"
	"math/rand"
	"time"

	"github.com/faroyam/url-short-cutter-API/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const symbols = "ABCDEFGHIJKLMNO5PQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

type url struct {
	LongURL  string
	ShortURL string
}
type mongo struct {
	MongoIP         string
	MongoDBName     string
	MongoCollection string
	Session         *mgo.Session
}

func (d *mongo) connect() error {
	var err error
	d.Session, err = mgo.Dial(d.MongoIP)
	return err
}

// Close mgo session
func (d *mongo) Close() {
	d.Session.Close()
}

func randomizer() (shortURL string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		shortURL += string(symbols[r.Intn(62)])
	}
	return shortURL
}

// Converter converts an user URL to short version
func Converter(URL string) (string, error) {

	c := DB.Session.DB(DB.MongoDBName).C(DB.MongoCollection)

	result := url{}
	err := c.Find(bson.M{"longurl": URL}).One(&result)
	if err != nil {
		data := &url{LongURL: URL, ShortURL: randomizer()}
		err = c.Insert(data)
		return data.ShortURL, err
	}
	return result.ShortURL, err
}

// ReConverter returns long url from short
func ReConverter(shortURL string) (string, error) {

	c := DB.Session.DB(DB.MongoDBName).C(DB.MongoCollection)

	result := url{}
	err := c.Find(bson.M{"shorturl": shortURL}).One(&result)
	return result.LongURL, err
}

// DB Data Access Object
var DB = mongo{}

func init() {
	DB.MongoIP = config.C.MongoIP
	DB.MongoDBName = config.C.MongoDBName
	DB.MongoCollection = config.C.MongoCollection
	err := DB.connect()
	if err != nil {
		log.Fatal("connetcing to db error")
	}
}
