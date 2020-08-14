package conf

import (
	"bytes"
	"os"
	"strings"
)

/*
替换环境变量等等内容
*/
func Escape(val string) string {
	start := strings.IndexByte(val, '$')
	if start == -1 {
		return val
	}

	buf := new(bytes.Buffer)
	mark := 0
	end := 0
	plen := len(val)
	for {
		if start == -1 {
			buf.WriteString(val[mark:])
			break
		} else {
			buf.WriteString(val[mark:start])
		}
		mark = start + 1
		if val[mark] == '{' {
			mark++
			end = nextByte(&val, '}', mark, plen)
			if end == -1 {
				buf.WriteString(val[start:])
				break
			} else {
				buf.WriteString(os.Getenv(val[mark:end]))
			}
			mark = end + 1
		} else {
			end = nextNotIdenByte(&val, mark, plen)
			if end == -1 {
				buf.WriteString(val[start:])
				break
			} else {
				buf.WriteString(os.Getenv(val[mark:end]))
			}
			mark = end
		}
		start = nextByte(&val, '$', mark, plen)
	}

	return buf.String()
}

func nextByte(v *string, c byte, start int, end int) int {
	for i := start; i < end; i++ {
		if (*v)[i] == c {
			return i
		}
	}
	return -1
}

func nextNotIdenByte(v *string, start int, end int) int {
	for i := start; i < end; i++ {
		if ch := (*v)[i]; !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_') {
			return i
		}
	}
	return -1
}
