package dht

import "sync"

var pool *sync.Pool

func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 8192)
		},
	}
}

// GetBuff 获取缓存
func GetBuff() []byte {
	buff := getBuff()
	return buff.([]byte)
}

// PutBuff 归还缓存
func PutBuff(buff interface{}) {
	pool.Put(buff)
}

func getBuff() interface{} {
	return pool.Get()
}
