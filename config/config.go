// Package config creates, structures, and imports gonc configuration files.
//
// Config files are named `conf.json` by default.
// JSON formatting is used due to portability, readability, and ease of use.
// Future versions will allow file specification with commandline arguments.

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type config struct {
	nick    string   `json:"nick"`
	servers []string `json:"servers"`
}

func Import(filePath ...string) (config, error) {
	loadedJSON, err := ioutil.ReadFile("./conf.json")
	if err != nil {
		fmt.Printf("err loading conf: %s\n", err)
		return config{}, errors.New("No config")
	}

	loadedConf := config{}
	err = json.Unmarshal(loadedJSON, &loadedConf)
	if err != nil {
		fmt.Printf("err mapping JSON to config: %s\n", err)
	}

	return loadedConf, err
}
