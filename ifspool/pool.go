package ifspool

import "sync"

type Slice struct {
	Items []interface{}
}

var slicePool = sync.Pool{
	New: func() interface{} {
		return &Slice{}
	},
}

func Get(size int) (slice *Slice) {
	slice = slicePool.Get().(*Slice)
	if size >= 0 {
		if slice.Items == nil || cap(slice.Items) < size {
			slice.Items = make([]interface{}, size)
		} else {
			slice.Items = slice.Items[:size]
		}
	}
	return slice
}

func GetNil(size int) (slice *Slice) {
	slice = Get(size)
	for i := 0; i < len(slice.Items); i++ {
		slice.Items[i] = nil
	}
	return
}

func Put(slice *Slice) {
	slicePool.Put(slice)
}

func PutNil(slice *Slice) {
	for i := 0; i < len(slice.Items); i++ {
		slice.Items[i] = nil
	}
	Put(slice)
}
