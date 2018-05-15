package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	SID     string
	HSID    string
	SSID    string
	APISID  string
	SAPISID string
	CONSENT string
	NID     string
	JAR     string
}

// LoadConfig Title says it all :)
func LoadConfig(configFile string) (*Config, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, errors.New("Config File doesn not exist")

	} else if err != nil {
		return nil, err
	}

	var config Config

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil

}

func main() {
	config, err := LoadConfig("config.toml")
	if err != nil {
		fmt.Println(err)
	}

	// For now I hardcode a data so I can feth while looping
	start, err := time.Parse("2006-1-2", "2018-1-1")
	if err != nil {
		// handle err
	}

	m := start.Month() - 1 // Because javascript

	month := strconv.Itoa(int(m))

	for date := start; date.Month() == start.Month(); date = date.AddDate(0, 0, 1) {
		d := date.Day()
		day := strconv.Itoa(int(d))

		fmt.Println("Start fetching timeline")
		fmt.Println("Month:" + month + "-- Day:" + day)
		fmt.Println("Reading Cookie info:")
		fmt.Printf("SID: %s\nHSID: %s\nSSID: %s\nAPISID: %s\nSAPISID: %s\nNID: %s\nJAR: %s\n",
			config.SID, config.HSID, config.SSID, config.APISID, config.SAPISID,
			config.NID, config.JAR)
		fmt.Println("Setup curl like fetch:")

		req, err := http.NewRequest("GET", "https://www.google.be/maps/timeline/kml?authuser=0&pb=!1m8!1m3!1i2018!2i"+month+"!3i"+day+"!2m3!1i2018!2i3!3i21", nil)
		if err != nil {
			// handle err
		}
		req.Header.Set("Authority", "www.google.be")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Mobile Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		req.Header.Set("X-Client-Data", "CJe2yQEIprbJAQjEtskBCKmdygEIqKPKAQ==")
		req.Header.Set("Referer", "https://www.google.be/")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9,nl;q=0.8")
		req.Header.Set("Cookie", "SID="+config.SID+"; HSID="+config.HSID+"; SSID="+config.SSID+"; APISID="+config.APISID+"; SAPISID="+config.SAPISID+"; CONSENT="+config.CONSENT+"; NID="+config.NID+"; 1P_JAR="+config.JAR)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		errr := ioutil.WriteFile("output-2018-"+month+"-"+day+".kml", body, 0644)
		if errr != nil {
			fmt.Println(errr)
		}

		fmt.Println("Let's see what we got Will Roger:")
		fmt.Println(string(body))
	}
}
