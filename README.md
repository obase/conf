# package conf
统一加载与管理conf.yml, 并提供快捷API解析与提取特定的节点值
.

加载顺序:
1. 环境变量CONF_YAML指定位置的conf.yml
2. 安装位置的conf.yml
3. 运行位置的conf.yml

例如
```
[/data/pvpbroker/]$ /usr/local/xxx 
则xxx加载conf.yml的顺序是:
1. 如果使用env CONF_YAML=conf.yml则只加载conf.yml, 否则
2. 加载/usr/local/conf.yml
3. 加载/data/pvpbroker/conf.yml
```

# Installation
- go get 
```
go get -u github.com/obase/conf
```
- go mod
```
go mod edit -require=github.com/obase/conf@latest
```

# Index
- Constants
```
const (
	CONF_YAML_FILE string = "conf.yml"
	PATH_STEP_SEP  byte   = '.'
	CONF_YAML_ENV  string = "CONF_YAML"
)
// 默认文件名称, 默认环境变量
const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
	TIME_LENGTH = len(TIME_LAYOUT)
)
// 统一时间格式
```
    

- Variables 
```
var Values map[interface{}]interface{} = make(map[interface{}]interface{})
```
全部值缓存

- type ScanFunc
```
type ScanFunc func(val interface{}) interface{}
```
扫描操作函数, 与Scanf()函数配合使用.

- func ToString
```
func ToString(val interface{}) string 
```
string转换函数
- func ToBool
```
func ToBool(val interface{}) bool 
```
bool转换函数

- func ToInt
```
func ToInt(val interface{}) int
```
int转换函数

- func ToInt64
```
func ToInt64(val interface{}) int64
```
int64转换函数

- func ToFloat64
```
func ToFloat64(val interface{}) float64 
```
float64转换函数
    
- func ToTime
```
func ToTime(val interface{}) time.Time
```
time.Time转换函数
    
- func ToDuration
```
func ToDuration(val interface{}) time.Duration
```
time.Duration转换函数

- func ToStringSlice
```
func ToStringSlice(val interface{}) []string
```
[]string转换函数, 支持逗号(",")切割. 例如"a,b,c"转换为[]string{"a","b","c"}

- func ToSlice
```
func ToSlice(val interface{}) []interface{}
```
[]interface{}泛型数组转换函数

- func ToMap
```
func ToMap(val interface{}) map[string]interface{}
```
map[string]interface{}泛型哈希表转换函数

- func ToStringMap
```
func ToStringMap(val interface{}) map[string]string
```
map[string]string转换函数

- func Elem
```
func Elem(val interface{}, key string) (ret interface{}, ok bool)
```
节点属性获取函数

- func ElemString
```
func ElemString(val interface{}, key string) (ret string, ok bool) 
```
节点属性获取函数

- func ElemBool
```
func ElemBool(val interface{}, key string) (ret bool, ok bool)
```
节点属性获取函数

- func ElemInt
```
func ElemInt(val interface{}, key string) (ret int, ok bool) 
```
节点属性获取函数

- func ElemInt64
```
func ElemInt64(val interface{}, key string) (ret int64, ok bool)
```
节点属性获取函数

- func ElemFloat64
```
func ElemFloat64(val interface{}, key string) (ret float64, ok bool)
```
节点属性获取函数

- func ElemTime
```
func ElemTime(val interface{}, key string) (ret time.Time, ok bool)
```
节点属性获取函数

- func ElemDuration
```
func ElemDuration(val interface{}, key string) (ret time.Duration, ok bool) 
```
节点属性获取函数

- func ElemSlice
```
func ElemSlice(val interface{}, key string) (ret []interface{}, ok bool)
```
节点属性获取函数

- func ElemStringSlice
```
func ElemStringSlice(val interface{}, key string) (ret []string, ok bool)
```
节点属性获取函数

- func ElemMap
```
func ElemMap(val interface{}, key string) (ret map[string]interface{}, ok bool) 
```
节点属性获取函数

- func ElemStringMap
```
func ElemStringMap(val interface{}, key string) (ret map[string]string, ok bool)
```
节点属性获取函数

- func Get
```
func Get(keys string) (val interface{}, ok bool) 
```
节点值获取函数, 如果keys为"",则返回全部值.

- func GetMap
```
func GetMap(keys string) (map[string]interface{}, bool)
```
节点值获取函数, 如果keys为"",则返回全部值.

- func GetStringMap
```
func GetStringMap(keys string) (map[string]string, bool)
```
节点值获取函数, 如果keys为"",则返回全部值.

- func GetBool
```
func GetBool(keys string) (bool, bool)
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiBool
```
func OptiBool(keys string, def bool) bool
```
节点值获取函数, 如果为空则返回默认值

- func MustBool
```
func MustBool(keys string) bool 
```
节点值获取函数, 如果为空则panic

- func GetString
```
func GetString(keys string) (string, bool)
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiString
```
func OptiString(keys string, def string) string
```
节点值获取函数, 如果为空则返回默认值

- func MustString
```
func MustString(keys string) string 
```
节点值获取函数, 如果为空则panic

- func GetInt
```
func GetInt(keys string) (int, bool) 
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiInt
```
func OptiInt(keys string, def int) int 
```
节点值获取函数, 如果为空则返回默认值

- func MustInt
```
func MustInt(keys string) int 
```
节点值获取函数, 如果为空则panic

- func GetInt64
```
func GetInt64(keys string) (int64, bool) 
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiInt64
```
func OptiInt64(keys string, def int64) int64
```
节点值获取函数, 如果为空则返回默认值

- func MustInt64
```
func MustInt64(keys string) int64 
```
节点值获取函数, 如果为空则panic

- func GetFloat64
```
func GetFloat64(keys string) (float64, bool)
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiFloat64
```
func OptiFloat64(keys string, def float64) float64
```
节点值获取函数, 如果为空则返回默认值

- func MustFloat64
```
func MustFloat64(keys string) float64
```
节点值获取函数, 如果为空则panic

- func GetTime
```
func GetTime(keys string) (time.Time, bool) 
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiTime
```
func OptiTime(keys string, def time.Time) time.Time
```
节点值获取函数, 如果为空则返回默认值

- func MustTime
```
func MustTime(keys string) time.Time 
```
节点值获取函数, 如果为空则panic

- func GetDuration
```
func GetDuration(keys string) (time.Duration, bool) 
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiDuration
```
func OptiDuration(keys string, def time.Duration) time.Duration
```
节点值获取函数, 如果为空则返回默认值

- func MustDuration
```
func MustDuration(keys string) time.Duration
```
节点值获取函数, 如果为空则panic

- func GetSlice
```
func GetSlice(keys string) ([]interface{}, bool)
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiSlice
```
func OptiSlice(keys string, def []interface{}) []interface{}
```
节点值获取函数, 如果为空则返回默认值

- func MustSlice
```
func MustSlice(keys string) []interface{}
```
节点值获取函数, 如果为空则panic

- func GetStringSlice
```
func GetStringSlice(keys string) ([]string, bool)
```
节点值获取函数, 如果keys为"",则返回全部值.

- func OptiStringSlice
```
func OptiStringSlice(keys string, def []string) []string 
```
节点值获取函数, 如果为空则返回默认值

- func MustStringSlice
```
func MustStringSlice(keys string) []string
```
节点值获取函数, 如果为空则panic

- func Scan
```
func Scan(keys string, ret interface{}) bool 
```
扫描节点值,并绑定到ret变量, 其中ret必须是个指针! 类似json.Unmarshal()

- func Scanf
```
func Scanf(keys string, f ScanFunc) (interface{}, bool) 
```
扫描节点值,并调用ScanFunc函数进行处理


- func Convert
```
func Convert(dst interface{}, src interface{}) (err error) 
```
转换函数, dst必须是指针

# Examples

conf.yml
```
# 业务逻辑配置
pvp:
  resourcePrefix: "x2_pvp_match_result_"
  # kafka主题,限制
  kafkaNotifyTopic: "x2_pvp_match_result_"
  kafkaNotifyLimit: 1024
  # redis队列,限制,延迟,休眠
  redisNotifyQueue: "FAILED_MATCH"
  redisNotifyLimit: 8
  redisNotityDelay: "5s"
  redisNotifySleep: "1s"
  # 赛事安排
  schedules:
    - category: "3m"
      matchCode: "CJ"
      season: "02"
      beginDate: "2019-07-01"
      endDate: "2019-07-02"
    - category: "3s"
      matchCode: "CJ"
      season: "02"
      beginDate: "2019-07-01"
      endDate: "2019-07-02"
```
代码
```
type Schedule struct {
	Category  string `yaml:"category"`
	MatchCode string `yaml:"matchCode"`
	Season    string `yaml:"season"`
}

func TestScan2(t *testing.T) {
	vl, _ := Get("pvp.schedules")
	fmt.Println(vl)

	var ss []*Schedule
	Scan("pvp.schedules", &ss)
	for _, s := range ss {
		fmt.Printf("category=%v,matchCode=%v,season=%v\n", s.Category, s.MatchCode, s.Season)
	}
}

```
输出
```
[map[beginDate:2019-07-01 category:3m endDate:2019-07-02 matchCode:CJ season:02] map[beginDate:2019-07-01 category:3s endDate:2019-07-02 matchCode:CJ season:02]]
- beginDate: "2019-07-01"
  category: 3m
  endDate: "2019-07-02"
  matchCode: CJ
  season: "02"
- beginDate: "2019-07-01"
  category: 3s
  endDate: "2019-07-02"
  matchCode: CJ
  season: "02"

category=3m,matchCode=CJ,season=02
category=3s,matchCode=CJ,season=02
--- PASS: TestScan2 (6.25s)
```
