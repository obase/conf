package conf

import (
	"bytes"
	"os"
	"strings"
)

/*
替换环境变量等等内容. 具体格式为: ${NAME:default}, 如果没有指定":"则保留${NAME}的字样
*/
func Escape(val string) string {

	var (
		start int
		end   int
	)

	start = strings.Index(val, "${")
	if start == -1 {
		return val
	}

	end = strings.IndexByte(val[start:], '}')
	if end == -1 {
		return val
	}

	buf := new(bytes.Buffer)
	for {
		end += start + 1
		buf.WriteString(val[:start])
		buf.WriteString(getenv(val[start:end]))
		val = val[end:]
		start = strings.Index(val, "${")
		if start == -1 {
			buf.WriteString(val)
			break
		}
		end = strings.IndexByte(val[start:], '}')
		if end == -1 {
			buf.WriteString(val)
			break
		}
	}

	return buf.String()
}

/*
格式: ${NAME:default}
- ${NAME}: 如果os.Getenv(NAME)不为空则返回结果, 为空则返回${NAME}. 即不做替换
- ${NAME:default}: 如果os.Getenv(NAME)不为空则返回结果, 为空则返回default
*/
func getenv(segment string) string {

	var (
		poc int
		val string
	)

	poc = strings.IndexByte(segment, ':') // 取首而非尾
	if poc == -1 {
		val = os.Getenv(segment[2 : len(segment)-1])
		if val != "" {
			return val
		}
		return segment
	} else {
		val = os.Getenv(segment[2:poc])
		if val != "" {
			return val
		}
		return segment[poc+1 : len(segment)-1]
	}
}
