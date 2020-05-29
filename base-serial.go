package conf

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

func EncodeJson(v interface{}) string {
	bs, err := json.Marshal(v)
	if err != nil {
		return ""
	} else {
		return string(bs)
	}
}

func DecodeJson(js string, pv interface{}) error {
	return json.Unmarshal([]byte(js), pv)
}

func EncodeYaml(v interface{}) string {
	bs, err := yaml.Marshal(v)
	if err != nil {
		return ""
	} else {
		return string(bs)
	}
}

func DecodeYaml(yml string, pv interface{}) error {
	return yaml.Unmarshal([]byte(yml), pv)
}

