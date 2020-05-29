package conf

import (
	"bytes"
	"encoding/json"
	"sync"
)

// 重用提高性能
var bytesBufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func BorrowBuffer() *bytes.Buffer {
	return bytesBufferPool.Get().(*bytes.Buffer)
}

func ReturnBuffer(buf *bytes.Buffer) {
	bytesBufferPool.Put(buf)
}


