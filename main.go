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

type tomlConfig struct {
	Cookie cookie
	Dates  dates
}

type cookie struct {
	SID     string
	HSID    string
	SSID    string
	APISID  string
	SAPISID string
	CONSENT string
	NID     string
	JAR     string
}

type dates struct {
	Startdate string
	Enddate   string
}

// LoadConfig Title says it all :)
func LoadConfig(configFile string) (*tomlConfig, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, errors.New("Config File doesn not exist")

	} else if err != nil {
		return nil, err
	}

	var config tomlConfig

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		return nil, err
	}

	return &config, nil

}

// FetchKML takes a date and downloads the kml file from said date
func FetchKML(day, month, year string, config *tomlConfig) ([]byte, error) {
	fmt.Println("Start fetching timeline")
	fmt.Println("Month:" + month + "-- Day:" + day)
	fmt.Println("Reading Cookie info:")
	fmt.Printf("SID: %s\nHSID: %s\nSSID: %s\nAPISID: %s\nSAPISID: %s\nNID: %s\nJAR: %s\n",
		config.Cookie.SID, config.Cookie.HSID, config.Cookie.SSID, config.Cookie.APISID, config.Cookie.SAPISID,
		config.Cookie.NID, config.Cookie.JAR)
	fmt.Println("Setup curl like fetch:")

	req, err := http.NewRequest("GET", "https://www.google.be/maps/timeline/kml?authuser=0&pb=!1m8!1m3!1i"+year+"!2i"+month+"!3i"+day+"!2m3!1i2018!2i3!3i21", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authority", "www.google.be")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Mobile Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("X-Client-Data", "CJe2yQEIprbJAQjEtskBCKmdygEIqKPKAQ==")
	req.Header.Set("Referer", "https://www.google.be/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,nl;q=0.8")
	req.Header.Set("Cookie", "SID="+config.Cookie.SID+"; HSID="+config.Cookie.HSID+"; SSID="+config.Cookie.SSID+"; APISID="+config.Cookie.APISID+"; SAPISID="+config.Cookie.SAPISID+"; CONSENT="+config.Cookie.CONSENT+"; NID="+config.Cookie.NID+"; 1P_JAR="+config.Cookie.JAR)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else {
		fmt.Println(resp)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	config, err := LoadConfig("config.toml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Startdate: %s\n", config.Dates.Startdate)

	// For now I hardcode a data so I can feth while looping
	start, err := time.Parse("2006-1-2", config.Dates.Startdate)
	end, err := time.Parse("2006-1-2", config.Dates.Enddate)
	if err != nil {
		fmt.Println(err)
	}

	m := start.Month() - 1 // Because javascript
	y := start.Year()      // Because javascript

	month := strconv.Itoa(int(m))
	year := strconv.Itoa(int(y))

	for date := start; date != end; date = date.AddDate(0, 0, 1) {
		d := date.Day()
		day := strconv.Itoa(int(d))

		body, err := FetchKML(day, month, year, config)
		if err != nil {
			fmt.Println(err)
		}

		errr := ioutil.WriteFile("output-"+year+"-"+month+"-"+day+".kml", body, 0644)
		if errr != nil {
			fmt.Println(errr)
		}

		fmt.Println("Let's see what we got Will Roger:")
		fmt.Println(string(body))
	}
}
