package conf

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/*
注意: 不支持动态加载. 修改conf.yml配置必须重启服务! 另外conf是其他obase的最底层依赖, 其日志输出采用fmt
*/

const (
	CONF_YAML_FILE string = "conf.yml"
	PATH_STEP_SEP  byte   = '.'
	CONF_YAML_ENV  string = "CONF_YAML"
)

var Values map[interface{}]interface{} = make(map[interface{}]interface{})

type ScanFunc func(val interface{}) interface{}

func ToString(val interface{}) string {
	switch val := val.(type) {
	case nil:
		return ""
	case string:
		return val
	default:
		return fmt.Sprintf("%v", val)
	}
}

func ToBool(val interface{}) bool {
	switch val := val.(type) {
	case nil:
		return false
	case bool:
		return val
	case string:
		return "true" == val
	}
	panic(fmt.Sprintf("invalid value to bool: %v", val))
}

func ToInt(val interface{}) int {
	switch val := val.(type) {
	case nil:
		return 0
	case int:
		return val
	case int64:
		return int(val)
	case string:
		if v, e := strconv.Atoi(val); e == nil {
			return v
		}
	}
	panic(fmt.Sprintf("invalid value to int: %v", val))
}

func ToInt64(val interface{}) int64 {
	switch val := val.(type) {
	case nil:
		return 0
	case int:
		return int64(val)
	case int64:
		return val
	case string:
		if v, e := strconv.ParseInt(val, 10, 64); e == nil {
			return v
		}
	}
	panic(fmt.Sprintf("invalid value to int: %v", val))
}

func ToFloat64(val interface{}) float64 {
	switch val := val.(type) {
	case nil:
		return 0
	case float64:
		return val
	case int:
		return float64(val)
	case int64:
		return float64(val)
	case string:
		if v, e := strconv.ParseFloat(val, 64); e == nil {
			return v
		}
	}
	panic(fmt.Sprintf("invalid value to float64: %v", val))
}

const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
	TIME_LENGTH = len(TIME_LAYOUT)
)

func ToTime(val interface{}) time.Time {
	switch val := val.(type) {
	case nil:
		return time.Unix(0, 0)
	case int:
		return time.Unix(int64(val), 0)
	case int64:
		return time.Unix(val, 0)
	case string:
		if len(val) == 0 {
			return ZERO_TIME
		}
		if vln := len(val); TIME_LENGTH == vln {
			if ret, err := time.Parse(TIME_LAYOUT, val); err == nil {
				return ret
			}
		} else {
			if ret, err := time.Parse(TIME_LAYOUT[0:vln], val); err == nil {
				return ret
			}
		}
	}
	panic(fmt.Sprintf("invalid value to time: %v", val))
}

func ToDuration(val interface{}) time.Duration {
	switch val := val.(type) {
	case nil:
		return 0
	case int:
		return time.Duration(val)
	case int64:
		return time.Duration(val)
	case string:
		if len(val) == 0 {
			return 0
		}
		if ret, err := time.ParseDuration(val); err == nil {
			return ret
		}
	}
	panic(fmt.Sprintf("invalid value to duration: %v", val))
}

func ToStringSlice(val interface{}) []string {
	switch val := val.(type) {
	case nil:
		return nil
	case []interface{}:
		ret := make([]string, len(val))
		for i, v := range val {
			ret[i] = ToString(v)
		}
		return ret
	case string:
		return strings.Split(val, ",")
	}
	panic(fmt.Sprintf("invalid value to strSlice: %v", val))
}

func ToSlice(val interface{}) []interface{} {
	switch val := val.(type) {
	case nil:
		return nil
	case []interface{}:
		return val
	}
	panic(fmt.Sprintf("invalid value to slice: %v", val))
}

func ToMap(val interface{}) map[string]interface{} {
	switch val := val.(type) {
	case nil:
		return nil
	case map[interface{}]interface{}:
		ret := make(map[string]interface{}, len(val))
		for k, v := range val {
			ret[ToString(k)] = v
		}
		return ret
	}
	panic(fmt.Sprintf("invalid value to map: %v", val))
}

func ToStringMap(val interface{}) map[string]string {
	switch val := val.(type) {
	case nil:
		return nil
	case map[interface{}]interface{}:
		ret := make(map[string]string, len(val))
		for k, v := range val {
			ret[ToString(k)] = ToString(v)
		}
		return ret
	}
	panic(fmt.Sprintf("invalid value to map: %v", val))
}

// it will panic if parse key failed
func Elem(val interface{}, key string) (ret interface{}, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			ret, ok = val[key]
		case map[string]interface{}:
			ret, ok = val[key]
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				ret, ok = val[idx], true
			}
		}
	}
	return
}

func ElemString(val interface{}, key string) (ret string, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToString(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToString(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToString(val[idx]), true
			}
		}
	}
	return
}

func ElemBool(val interface{}, key string) (ret bool, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToBool(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToBool(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToBool(val[idx]), true
			}
		}
	}
	return
}

func ElemInt(val interface{}, key string) (ret int, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToInt(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToInt(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToInt(val[idx]), true
			}
		}
	}
	return
}

func ElemInt64(val interface{}, key string) (ret int64, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToInt64(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToInt64(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToInt64(val[idx]), true
			}
		}
	}
	return
}

func ElemFloat64(val interface{}, key string) (ret float64, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToFloat64(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToFloat64(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToFloat64(val[idx]), true
			}
		}
	}
	return
}

func ElemTime(val interface{}, key string) (ret time.Time, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToTime(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToTime(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToTime(val[idx]), true
			}
		}
	}
	return
}

func ElemDuration(val interface{}, key string) (ret time.Duration, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToDuration(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToDuration(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToDuration(val[idx]), true
			}
		}
	}
	return
}

func ElemSlice(val interface{}, key string) (ret []interface{}, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToSlice(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToSlice(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToSlice(val[idx]), true
			}
		}
	}
	return
}

func ElemStringSlice(val interface{}, key string) (ret []string, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToStringSlice(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToStringSlice(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToStringSlice(val[idx]), true
			}
		}
	}
	return
}

func ElemMap(val interface{}, key string) (ret map[string]interface{}, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToMap(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToMap(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToMap(val[idx]), true
			}
		}
	}
	return
}

func ElemStringMap(val interface{}, key string) (ret map[string]string, ok bool) {
	if val != nil {
		switch val := val.(type) {
		case map[interface{}]interface{}:
			if kv, ok := val[key]; ok {
				return ToStringMap(kv), true
			}
		case map[string]interface{}:
			if kv, ok := val[key]; ok {
				return ToStringMap(kv), true
			}
		case []interface{}:
			if idx, err := strconv.Atoi(key); err == nil && idx >= 0 && idx < len(val) {
				return ToStringMap(val[idx]), true
			}
		}
	}
	return
}

func Get(keys string) (val interface{}, ok bool) {

	if keys == "" {
		return Values, true
	}

	val = Values
	for mk, ln := 0, len(keys); mk < ln; {
		ps := mk
		for ps < ln {
			if keys[ps] == PATH_STEP_SEP {
				break
			} else {
				ps++
			}
		}
		if mk < ps {
			if val, ok = Elem(val, keys[mk:ps]); !ok {
				// 非空需要重置,避免返回误解
				if val != nil {
					val = nil
				}
				return
			}
		}
		mk = ps + 1
	}
	return
}

func GetMap(keys string) (map[string]interface{}, bool) {
	if vl, ok := Get(keys); ok {
		return ToMap(vl), true
	}
	return nil, false
}

func GetStringMap(keys string) (map[string]string, bool) {
	if vl, ok := Get(keys); ok {
		return ToStringMap(vl), true
	}
	return nil, false
}

func GetBool(keys string) (bool, bool) {
	if vl, ok := Get(keys); ok {
		return ToBool(vl), true
	}
	return false, false
}

func OptiBool(keys string, def bool) bool {
	if vl, ok := GetBool(keys); ok {
		return vl
	}
	return def
}

func MustBool(keys string) bool {
	if vl, ok := GetBool(keys); ok {
		return vl
	}
	panic("missing bool config: " + keys)
}

func GetString(keys string) (string, bool) {
	if vl, ok := Get(keys); ok {
		return ToString(vl), true
	}
	return "", false
}

func OptiString(keys string, def string) string {
	if vl, ok := GetString(keys); ok {
		return vl
	}
	return def
}

func MustString(keys string) string {
	if vl, ok := GetString(keys); ok {
		return vl
	}
	panic("missing string config: " + keys)
}

func GetInt(keys string) (int, bool) {
	if vl, ok := Get(keys); ok {
		return ToInt(vl), true
	}
	return 0, false
}

func OptiInt(keys string, def int) int {
	if vl, ok := GetInt(keys); ok {
		return vl
	}
	return def
}

func MustInt(keys string) int {
	if vl, ok := GetInt(keys); ok {
		return vl
	}
	panic("missing int config: " + keys)
}

func GetInt64(keys string) (int64, bool) {
	if vl, ok := Get(keys); ok {
		return ToInt64(vl), true
	}
	return 0, false
}

func OptiInt64(keys string, def int64) int64 {
	if vl, ok := GetInt64(keys); ok {
		return vl
	}
	return def
}

func MustInt64(keys string) int64 {
	if vl, ok := GetInt64(keys); ok {
		return vl
	}
	panic("missing int64 config: " + keys)
}

func GetFloat64(keys string) (float64, bool) {
	if vl, ok := Get(keys); ok {
		return ToFloat64(vl), true
	}
	return 0, false
}

func OptiFloat64(keys string, def float64) float64 {
	if vl, ok := GetFloat64(keys); ok {
		return vl
	}
	return def
}

func MustFloat64(keys string) float64 {
	if vl, ok := GetFloat64(keys); ok {
		return vl
	}
	panic("missing float64 config: " + keys)
}

var ZERO_TIME = time.Unix(0, 0)

func GetTime(keys string) (time.Time, bool) {
	if vl, ok := Get(keys); ok {
		return ToTime(vl), true
	}
	return ZERO_TIME, false
}

func OptiTime(keys string, def time.Time) time.Time {
	if vl, ok := GetTime(keys); ok {
		return vl
	}
	return def
}

func MustTime(keys string) time.Time {
	if vl, ok := GetTime(keys); ok {
		return vl
	}
	panic("missing time config: " + keys)
}

func GetDuration(keys string) (time.Duration, bool) {
	if vl, ok := Get(keys); ok {
		return ToDuration(vl), true
	}
	return 0, false
}

func OptiDuration(keys string, def time.Duration) time.Duration {
	if vl, ok := GetDuration(keys); ok {
		return vl
	}
	return def
}

func MustDuration(keys string) time.Duration {
	if vl, ok := GetDuration(keys); ok {
		return vl
	}
	panic("missing duration config: " + keys)
}

func GetSlice(keys string) ([]interface{}, bool) {
	if vl, ok := Get(keys); ok {
		return ToSlice(vl), true
	}
	return nil, false
}

func OptiSlice(keys string, def []interface{}) []interface{} {
	if vl, ok := GetSlice(keys); ok {
		return vl
	}
	return def
}

func MustSlice(keys string) []interface{} {
	if vl, ok := GetSlice(keys); ok && len(vl) > 0 {
		return vl
	}
	panic("missing slice config: " + keys)
}

func GetStringSlice(keys string) ([]string, bool) {
	if vl, ok := Get(keys); ok {
		return ToStringSlice(vl), true
	}
	return nil, false
}

func OptiStringSlice(keys string, def []string) []string {
	if vl, ok := GetStringSlice(keys); ok {
		return vl
	}
	return def
}

func MustStringSlice(keys string) []string {
	if vl, ok := GetStringSlice(keys); ok && len(vl) > 0 {
		return vl
	}
	panic("missing string slice config: " + keys)
}

func Scan(keys string, ret interface{}) bool {
	if vl, ok := Get(keys); ok {
		if bs, err := yaml.Marshal(vl); err == nil {
			if err = yaml.Unmarshal(bs, ret); err == nil {
				return true
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	return false
}

func Scanf(keys string, f ScanFunc) (interface{}, bool) {
	if vl, ok := Get(keys); ok {
		return f(vl), true
	}
	return nil, false
}

/*
1. 参数-conf <xxx>或-conf=<xxx>
2. 当前路径./conf.yml
3. 工作路径 $CWD/conf.yml
*/
func init() {
	var path string

	path = os.Getenv(CONF_YAML_ENV)
	if path == "" {
		loc, _ := exec.LookPath(os.Args[0])
		path = filepath.Join(filepath.Dir(loc), CONF_YAML_FILE)
		if fi, err := os.Stat(path); fi == nil || os.IsNotExist(err) {
			dir, _ := os.Getwd()
			path = filepath.Join(dir, CONF_YAML_FILE)
			if fi, err := os.Stat(path); fi == nil || os.IsNotExist(err) {
				return
			}
		}
	}

	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Load conf failed: "+path)
		panic(err)
	}
	defer file.Close()

	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bs, &Values)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, "Load conf success: "+path)
}
