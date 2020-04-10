package bufpool

import "sync"

type Buf struct {
	Bytes []byte
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return &Buf{}
	},
}

func Get(size int) (buf *Buf) {
	buf = bufPool.Get().(*Buf)
	if size >= 0 {
		if buf.Bytes == nil || cap(buf.Bytes) < size {
			buf.Bytes = make([]byte, size)
		} else {
			buf.Bytes = buf.Bytes[:size]
		}
	}
	return buf
}

func Put(buf *Buf) {
	bufPool.Put(buf)
}
