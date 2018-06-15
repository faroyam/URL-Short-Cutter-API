package urlshortcutter

import (
	"log"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const symbols = "ABCDEFGHIJKLMNO5PQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

type url struct {
	LongURL  string
	ShortURL string
}

func randomizer() (shortURL string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		shortURL += string(symbols[r.Intn(62)])
	}
	return shortURL
}

// Converter converts an user URL to short version
func Converter(URL, mongoURL string) (shortURL string) {

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	c := session.DB("url-short-cutter").C("URLs")

	result := url{}
	err = c.Find(bson.M{"longurl": URL}).One(&result)
	if err != nil {
		data := &url{LongURL: URL, ShortURL: randomizer()}
		err = c.Insert(data)
		if err != nil {
			log.Println(err)
		}
		return data.ShortURL
	}
	return result.ShortURL
}

// ReConverter returns long url from short
func ReConverter(shortURL, mongoURL string) (longURL string) {

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	c := session.DB("url-short-cutter").C("URLs")

	result := url{}
	err = c.Find(bson.M{"shorturl": shortURL}).One(&result)
	if err != nil {
		return ""
	}
	return result.LongURL
}
