package bbpool

import (
	"strings"
	"sync"
)

var sbPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func Get() (sb *strings.Builder) {
	sb = sbPool.Get().(*strings.Builder)
	sb.Reset()
	return
}

func Put(sb *strings.Builder) {
	sbPool.Put(sb)
}
