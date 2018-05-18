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
	ioutil.WriteFile("test_config.toml", testConfig, os.FileMode(int(0664)))
	// clean up config.json after tests are completed
	defer os.Remove("test_config.toml")

	config, err := LoadConfig("test_config.toml")
	if err != nil {
		t.Error(err)
	}
	t.Logf("config successfully created: \n %v", config)
}

func testFetchKML(t *testing.T) {
	// Test if request returns 200

	// Test if request return kml File

}
