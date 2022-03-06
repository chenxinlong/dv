package configs

import (
	_ "embed"
	"encoding/json"
)

var Conf Config

type Config struct {
	Display []string `json:"display"`
	Deepl   Deepl    `json:"deepl"`
	Weather Weather  `json:"weather"`
}

type Deepl struct {
	AuthKey string `json:"auth_key"`
}

type Weather struct {
	Key string `json:"key"`
}

//go:embed config.json
var ConfigBytes []byte

func LoadConfig() (err error) {
	err = json.Unmarshal(ConfigBytes, &Conf)
	return
}
