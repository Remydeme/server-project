package config

import (
	"github.com/jinzhu/configor"
	"os"
	"testing"
)

var testTable = []Configuration{

	{
		Database{
			Dialect:  "postgres",
			Username: "remy",
			Password: "1234567",
			Host:     "exemple.com",
			Port:     5652,
			SSLMode:  true,
			DBname:   "iurgence_test",
		},
		JWT{
			Secret: "Remyistheboss",
		},
		Server{
			Port: ":8080",
		},
	},
}

func testConfigLoading(t *testing.T) {

	var conf Configuration
	configPath := os.Getenv("CONFIG_PATH")
	err := configor.Load(&conf, configPath)

	if err != nil {
		t.Error("Failed to load configuration")
	}
}

func testConfigValues(t *testing.T) {
	var conf Configuration

	configPath := os.Getenv("CONFIG_TEST_PATH")

	err := configor.Load(&conf, configPath)

	if err != nil {
		t.Error("Failed to load configuration")
	}

	if conf == testTable[0] {
		t.Errorf(" Error the two config are diferent :  %v , get %v", testTable[0], conf)
	}
}
