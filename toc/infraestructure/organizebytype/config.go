package organizebytype

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Rules []RuleConfig
}

type RuleConfig struct {
	Expresion string `json:"expresion"`
	Folder    string `json:"folder"`
}

func ReadConfig(configFile string) Config {

	jsonFile, err := os.Open(configFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var conf Config

	json.Unmarshal(byteValue, &conf)

	return conf
}
