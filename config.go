package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const URL = "http://geci.herokuapp.com/gotrabigokiakiraly/trabi_v1.json"

type Config struct {
	StackCurve float64
	RaiseCurve float64
}

func NewConfig() *Config {
	config := &Config{}

	go config.Update()

	return config
}

func (cfg *Config) Update() {
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

			err = json.Unmarshal(data, cfg)
			if err != nil {
				log.Printf("Error parsing config: %s\n", err)
				break
			}

			log.Println("Downloaded config")
			log.Printf("New configuration: %s\n", cfg.String())
			break
		}

		time.Sleep(5 * time.Second)
	}
}

func (cfg *Config) String() string {
	return fmt.Sprintf("StackCurve=%f, RaiseCurve=%f", cfg.StackCurve, cfg.RaiseCurve)
}
