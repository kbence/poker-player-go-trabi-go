package fuzzy

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const URL = "http://geci.herokuapp.com/gotrabigokiakiraly/trabi_v1.json"

type Config struct {
}

func NewConfig() *Config {
	config := &Config{}

	go config.Update()

	return config
}

func (d *Config) Update() {
	log.Printf("Start scraping URL \"%s\"\n", URL)

	for {
		for {
			res, err := http.Get(URL)
			if err != nil {
				log.Printf("Error querying config: %s\n", err)
				break
			}

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Printf("Error reading config: %s\n", err)
				break
			}

			err = json.Unmarshal(data, d)
			if err != nil {
				log.Printf("Error parsing config: %s\n", err)
				break
			}

			log.Println("Downloaded config")
			break
		}

		time.Sleep(5 * time.Second)
	}
}
