package gccode

import (
	"sync"
	"time"
)

/*
缓存基础结构类
 */
type CacheObject struct {
	sync.RWMutex
	//k-v
	key string
	value interface{}
	//持续时间
	expiryTime time.Duration

	createTime time.Time
	lastAccessTime time.Time
	accessCount int64

}
//--------------------CacheObject new-------------------------
func  NewCacheObject(key string,value interface{},exp time.Duration) *CacheObject  {
	now := time.Now()
	co := CacheObject{
		key:			key,
		value:			value,
		expiryTime:		exp,
		createTime:		now,
		lastAccessTime:	now,
		accessCount:	0,
	}
	return &co
}

//--------------------CacheObject func-------------------
func (obj *CacheObject)ExpTime()time.Duration{
	return obj.expiryTime
}
func (obj *CacheObject)Key()string{
	return obj.key
}
func (obj *CacheObject)Value()interface{}{
	obj.RLock()
	defer obj.RUnlock()
	return obj.value
}
func (obj *CacheObject)AccessCount()int64{
	obj.RLock()
	defer obj.RUnlock()
	return obj.accessCount
}
