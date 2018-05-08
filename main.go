package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type cookie struct {
	SID     string
	HSID    string
	SSID    string
	APISID  string
	SAPISID string
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

}
