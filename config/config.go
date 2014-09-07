// Package config creates, structures, and imports gonc configuration files.
//
// Config files are named `conf.json` by default.
// JSON formatting is used due to portability, readability, and ease of use.
// Future versions will allow file specification with commandline arguments.

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Nick    string   `json:"nick"`
	Servers []string `json:"servers"`
}

func Import(filePath ...string) Config {
	loadedJSON, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		fmt.Printf("err loading conf: %s\n", err)
		create()
		loadedJSON, err = ioutil.ReadFile("./conf.json")
	}

	loadedConf := Config{}
	err = json.Unmarshal(loadedJSON, &loadedConf)
	if err != nil {
		fmt.Printf("err mapping JSON to config: %s\n", err)
	}

	return loadedConf
}
