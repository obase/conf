package v2

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println(Get(""))
	fmt.Println(Get("pvp"))
	fmt.Println(Get("pvp.kafkaNotifyLimit"))
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
