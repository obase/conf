package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strconv"
	"strings"
	"time"
)

func ToString(val interface{}) string {
	switch val := val.(type) {
	case nil:
		return ""
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(uint64(val), 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return fmt.Sprint(val)
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
	panic(fmt.Sprintf("invalid value to bool: %+v", val))
}

func ToInt(val interface{}) int {
	switch val := val.(type) {
	case nil:
		return 0
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	case uint:
		return int(val)
	case uint8:
		return int(val)
	case uint16:
		return int(val)
	case uint32:
		return int(val)
	case uint64:
		return int(val)
	case float32:
		return int(val)
	case float64:
		return int(val)
	case string:
		if v, e := strconv.Atoi(val); e == nil {
			return v
		}
	}
	panic(fmt.Sprintf("invalid value to int: %+v", val))
}

func ToInt64(val interface{}) int64 {
	switch val := val.(type) {
	case nil:
		return 0
	case int64:
		return val
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case string:
		if v, e := strconv.ParseInt(val, 10, 64); e == nil {
			return v
		}
	}
	panic(fmt.Sprintf("invalid value to int64: %+v", val))
}

func ToFloat64(val interface{}) float64 {
	switch val := val.(type) {
	case nil:
		return 0
	case float64:
		return val
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case string:
		if v, e := strconv.ParseFloat(val, 64); e == nil {
			return v
		}
	}
	panic(fmt.Sprintf("invalid value to float64: %v", val))
}

const (
	DATETIME_LAYOUT string = "2006-01-02 15:04:05"
	DATETIME_LENGTH int    = len(DATETIME_LAYOUT)
)

func ToTime(val interface{}) time.Time {
	switch val := val.(type) {
	case nil:
		return time.Unix(0, 0)
	case int:
		return time.Unix(int64(val), 0)
	case int8:
		return time.Unix(int64(val), 0)
	case int16:
		return time.Unix(int64(val), 0)
	case int32:
		return time.Unix(int64(val), 0)
	case int64:
		return time.Unix(val, 0)
	case uint:
		return time.Unix(int64(val), 0)
	case uint8:
		return time.Unix(int64(val), 0)
	case uint16:
		return time.Unix(int64(val), 0)
	case uint32:
		return time.Unix(int64(val), 0)
	case uint64:
		return time.Unix(int64(val), 0)
	case string:
		vln := len(val)
		if vln == DATETIME_LENGTH {
			if ret, err := time.ParseInLocation(DATETIME_LAYOUT, val, time.Local); err == nil {
				return ret
			}
		} else if vln < DATETIME_LENGTH {
			if ret, err := time.ParseInLocation(DATETIME_LAYOUT[:vln], val, time.Local); err == nil {
				return ret
			}
		} else if vln > DATETIME_LENGTH {
			if ret, err := time.ParseInLocation(DATETIME_LAYOUT, val[:DATETIME_LENGTH], time.Local); err == nil {
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
		return time.Duration(int64(val))
	case int8:
		return time.Duration(int64(val))
	case int16:
		return time.Duration(int64(val))
	case int32:
		return time.Duration(int64(val))
	case int64:
		return time.Duration(val)
	case uint:
		return time.Duration(int64(val))
	case uint8:
		return time.Duration(int64(val))
	case uint16:
		return time.Duration(int64(val))
	case uint32:
		return time.Duration(int64(val))
	case uint64:
		return time.Duration(int64(val))
	case string:
		if len(val) == 0 {
			return 0
		} else if ret, err := time.ParseDuration(val); err == nil {
			return ret
		}
	}
	panic(fmt.Sprintf("invalid value to duration: %v", val))
}

func ToStringSlice(val interface{}) []string {
	switch val := val.(type) {
	case nil:
		return nil
	case []string:
		return val
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
	case map[string]interface{}:
		return val
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
	case map[string]string:
		return val
	case map[string]interface{}:
		ret := make(map[string]string, len(val))
		for k, v := range val {
			ret[k] = ToString(v)
		}
		return ret
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
	return
}

func ElemTime(val interface{}, key string) (ret time.Time, ok bool) {
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
	return
}

func ElemDuration(val interface{}, key string) (ret time.Duration, ok bool) {
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
	return
}

func ElemSlice(val interface{}, key string) (ret []interface{}, ok bool) {
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
	return
}

func ElemStringSlice(val interface{}, key string) (ret []string, ok bool) {
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
	return
}

func ElemMap(val interface{}, key string) (ret map[string]interface{}, ok bool) {
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
	return
}

func ElemStringMap(val interface{}, key string) (ret map[string]string, ok bool) {
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
	return
}

func Convert(src interface{}, dst interface{}) (err error) {
	if src == nil {
		return
	}
	bs, err := yaml.Marshal(src)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(bs, dst)
	return
}
