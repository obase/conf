package conf

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

const (
	CONF_YAML_FILE string = "conf.yml"
	PATH_STEP_SEP  byte   = '.'
	CONF_YAML_ENV  string = "CONF_YAML"
)

type BindFunc func(val interface{}) interface{}

//注意: 只支持读,不支持写. 保证性能的情况下才不会触发并发问题
var (
	rwmutx sync.RWMutex
	values = make(map[string]interface{})
)

func Setup(vs map[string]interface{}) {
	rwmutx.Lock()
	for k, v := range vs {
		if v == nil {
			delete(values, k)
		} else {
			values[k] = v
		}
	}
	rwmutx.Unlock()
}

func Get(keys string) (interface{}, bool) {

	if keys == "" {
		var val = make(map[string]interface{}, len(values))
		rwmutx.RLock()
		for k, v := range values {
			val[k] = v
		}
		rwmutx.RUnlock()
		return val, true
	} else {
		var val interface{} = values
		var ok bool
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
				rwmutx.RLock()
				val, ok = Elem(val, keys[mk:ps])
				rwmutx.RUnlock()
				if !ok {
					// 非空需要重置,避免返回误解
					return nil, false
				}
			}
			mk = ps + 1
		}
		return val, ok
	}
}

func GetMap(keys string) (map[string]interface{}, bool) {
	if vl, ok := Get(keys); ok {
		return ToMap(vl), true
	}
	return nil, false
}

func OptiMap(keys string, def map[string]interface{}) map[string]interface{} {
	if vl, ok := GetMap(keys); ok {
		return vl
	}
	return def
}

func MustMap(keys string) map[string]interface{} {
	if vl, ok := GetMap(keys); ok {
		return vl
	}
	panic("missing map config: " + keys)
}

func GetStringMap(keys string) (map[string]string, bool) {
	if vl, ok := Get(keys); ok {
		return ToStringMap(vl), true
	}
	return nil, false
}

func OptiStringMap(keys string, def map[string]string) map[string]string {
	if vl, ok := GetStringMap(keys); ok {
		return vl
	}
	return def
}

func MustStringMap(keys string) map[string]string {
	if vl, ok := GetStringMap(keys); ok {
		return vl
	}
	panic("missing string map config: " + keys)
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

func Bind(keys string, ret interface{}) bool {
	if vl, ok := Get(keys); ok {
		if bs, err := yaml.Marshal(vl); err == nil {
			if err = yaml.Unmarshal(bs, ret); err == nil {
				return true
			}
		}
	}
	return false
}

func MustBind(keys string, ret interface{}) {
	if vl, ok := Get(keys); ok {
		if bs, err := yaml.Marshal(vl); err == nil {
			if err = yaml.Unmarshal(bs, ret); err == nil {
				return
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

func BindWith(keys string, f BindFunc) (interface{}, bool) {
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
		fmt.Fprintln(os.Stderr, "\nLoad conf failed: "+path)
		panic(err)
	}
	defer file.Close()

	var vs map[string]interface{}
	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(Escape(string(bs))), &vs)
	if err != nil {
		panic(err)
	}
	Setup(vs)
	fmt.Fprintln(os.Stdout, "\nLoad conf success: "+path)
}
