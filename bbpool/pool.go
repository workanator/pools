package bbpool

import (
	"bytes"
	"sync"
)

var bbPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func Get() (bb *bytes.Buffer) {
	bb = bbPool.Get().(*bytes.Buffer)
	bb.Reset()
	return
}

func Put(bb *bytes.Buffer) {
	bbPool.Put(bb)
}
