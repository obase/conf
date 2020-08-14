package conf

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println(Get(""))
	fmt.Println(Get("pvp"))
}

func TestEscape(t *testing.T) {

	os.Setenv("TEST", "这是一个测试")
	os.Setenv("STDOUT", "标准输出")

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

	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	// 添加环境变量替换支持



	str := Escape(string(bs))
	fmt.Println(str)

	var data map[string]interface{}

	err = yaml.Unmarshal([]byte(str), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据:", data)
}


func TestGetBool(t *testing.T) {
	fmt.Println(GetBool("pvp.testBool"))
}

func TestGetFloat64(t *testing.T) {
	fmt.Println(GetFloat64("pvp.kafkaNotifyLimit"))
}

func TestGetDuration(t *testing.T) {
	fmt.Println(GetDuration("pvp.redisNotityDelay2"))
}

func TestGetTime(t *testing.T) {
	fmt.Println(GetTime("pvp.beginDate"))
}

type Schedule struct {
	Category  string `yaml:"category"`
	MatchCode string `yaml:"matchCode"`
	Season    string `yaml:"season"`
}

func TestScan2(t *testing.T) {
	vl, _ := Get("pvp.schedules")
	fmt.Println(vl)

	var ss []*Schedule
	Bind("pvp.schedules", &ss)
	for _, s := range ss {
		fmt.Printf("category=%v,matchCode=%v,season=%v\n", s.Category, s.MatchCode, s.Season)
	}
}
