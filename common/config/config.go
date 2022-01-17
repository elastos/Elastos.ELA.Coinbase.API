// Copyright (c) 2017-2022 The Elastos Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	// DefaultConfigFilename indicates the file name of config.
	DefaultConfigFilename = "./config.json"
)

var (
	Version    string
	Parameters ConfigParams
)

type RpcConfiguration struct {
	User        string   `json:"User"`
	Pass        string   `json:"Pass"`
	WhiteIPList []string `json:"WhiteIPList"`
}

type Configuration struct {
	ActiveNet   string     `json:"ActiveNet"`
	Version     uint32     `json:"Version"`
	ServerPort  int        `json:"ServerPort"`
	MainNodeRPC *RpcConfig `json:"MainNodeRPC"`
}

type RpcConfig struct {
	IpAddress    string `json:"IpAddress"`
	HttpJsonPort int    `json:"HttpJsonPort"`
	User         string `json:"User"`
	Pass         string `json:"Pass"`
}

type ConfigFile struct {
	ConfigFile Configuration `json:"Configuration"`
}

type ConfigParams struct {
	*Configuration
}

func Initialize() {
	file, e := ioutil.ReadFile(DefaultConfigFilename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return
	}
	i := ConfigFile{}
	// Remove the UTF-8 Byte Order Mark
	file = bytes.TrimPrefix(file, []byte("\xef\xbb\xbf"))

	e = json.Unmarshal(file, &i)
	var config ConfigFile
	switch strings.ToLower(i.ConfigFile.ActiveNet) {
	case "testnet", "test":
		config = testnet
	case "regnet", "reg":
		config = regnet
	default:
		config = mainnet
	}

	Parameters.Configuration = &(config.ConfigFile)

	e = json.Unmarshal(file, &config)
	if e != nil {
		fmt.Printf("Unmarshal json file erro %v", e)
		os.Exit(1)
	}

	e = json.Unmarshal(file, &config)
	if e != nil {
		fmt.Printf("Unmarshal json file erro %v", e)
		os.Exit(1)
	}

	var out bytes.Buffer
	err := json.Indent(&out, file, "", "")
	if err != nil {
		fmt.Printf("Config file error: %v\n", e)
		os.Exit(1)
	}

	if Parameters.Configuration.MainNodeRPC == nil {
		fmt.Printf("Need to set main node in config file\n")
		return
	}
	printJson(Parameters)
}

func printJson(data interface{}) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := new(bytes.Buffer)
	json.Indent(buf, dataBytes, "", "    ")
	fmt.Println(string(buf.Bytes()))
}
