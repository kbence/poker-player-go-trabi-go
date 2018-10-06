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
	Version  string
	NewLogic bool
	Curves   struct {
		StackCurve float64
		RaiseCurve float64
		HandCurve  float64
	}
	ConfidenceLevels struct {
		AllIn float64
		Raise float64
		Call  float64
	}
}

func NewConfig() *Config {
	config := &Config{
		Version:  "Default inline version",
		NewLogic: false,
	}
	config.ConfidenceLevels.AllIn = 0.9
	config.ConfidenceLevels.Raise = 0.8
	config.ConfidenceLevels.Call = 0.5

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
	return fmt.Sprintf("Curves={StackCurve=%f, RaiseCurve=%f}, "+
		"ConfidenceLevels={AllIn=%f, Raise=%f, Call=%f}"+
		"NewLogic=%b, Version=%s",
		cfg.Curves.StackCurve,
		cfg.Curves.RaiseCurve,
		cfg.ConfidenceLevels.AllIn,
		cfg.ConfidenceLevels.Raise,
		cfg.ConfidenceLevels.Call,
		cfg.NewLogic,
		cfg.Version)
}
