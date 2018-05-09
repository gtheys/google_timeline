package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"net/http"
)

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

func main() {
	var config cookie
	if _, err := toml.DecodeFile("cookie.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Start fetching timeline")
	fmt.Println("Reading Cookie info:")
	fmt.Printf("SID: %s\nHSID: %s\nSSID: %s\nAPISID: %s\nSAPISID: %s\nNID: %s\nJAR: %s\n",
		config.SID, config.HSID, config.SSID, config.APISID, config.SAPISID,
		config.NID, config.JAR)
	fmt.Println("Setup curl like fetch:")

	req, err := http.NewRequest("GET", "https://www.google.be/maps/timeline/kml?authuser=0&pb=!1m8!1m3!1i2018!2i3!3i21!2m3!1i2018!2i3!3i21", nil)
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

}
