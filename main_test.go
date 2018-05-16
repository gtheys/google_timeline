package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// a dummy config file to pass in as a slice of bytes
var testConfig = []byte(`
SID = "SIDvalue"
HSID = "HSIDvalue"
SSID = "SSIDvalue"
APISID = "APISIDvalue"
SAPISID = "SAPISIDvalue"
CONSENT = "CONSENTvalue"
NID = "NIDvalue"  
JAR = "JARVAlue"
`)

func TestLoadConfig(t *testing.T) {
	// write config to current folder (0664 denotes the permissions in octal notation)
	ioutil.WriteFile("config.toml", testConfig, os.FileMode(int(0664)))
	// clean up config.json after tests are completed
	//defer os.Remove("config.toml")

	config, err := LoadConfig("config.toml")
	if err != nil {
		t.Error(err)
	}
	t.Logf("config successfully created: \n %v", config)
}
