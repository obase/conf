package conf

import (
	"fmt"
	"testing"
	"time"
)

func TestElem(t *testing.T) {
	test, ok := Get("test")
	if !ok {
		fmt.Println("not existed")
	}
	fmt.Println(ElemBool(test, "bool"))
	fmt.Println(ElemInt(test, "long"))
	fmt.Println(ElemFloat64(test, "float"))
	fmt.Println(ElemTime(test, "time"))
	fmt.Println(ElemDuration(test, "duration"))
	fmt.Println(ElemStringSlice(test, "slice"))
	fmt.Println(ElemStringSlice(test, "slice2"))
	fmt.Println(ElemStringMap(test, "map"))
}

func TestGet(t *testing.T) {
	fmt.Println(GetBool("test.bool"))
	fmt.Println(GetInt("test.long"))
	fmt.Println(GetFloat64("test.float"))
	fmt.Println(GetTime("test.time"))
	fmt.Println(GetDuration("test.duration"))
	fmt.Println(GetStringSlice("test.slice"))
	fmt.Println(GetStringSlice("test.slice2"))
	fmt.Println(GetStringMap("test.map"))
}

func TestScan(t *testing.T) {
	var m map[string]interface{}
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		Scan("test.map", &m)
	}
	end := time.Now()
	fmt.Printf("used time: %v\n", end.Sub(start).Nanoseconds()/1000000)
	fmt.Println(m)
}
