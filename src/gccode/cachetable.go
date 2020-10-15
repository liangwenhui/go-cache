package gccode

import (
	"fmt"
	"sync"
	"time"
)

type CacheTable struct {
	sync.RWMutex
	name string
	data map[string]*CacheObject

	//logger *log.Logger

	loadNilDataCallback func(key string,arg ...interface{}) *CacheObject
	addedCallback []func(obj *CacheObject)
	deleteCallBack []func(obj *CacheObject)

	clearupTimer *time.Timer
	clearupInterval time.Duration

}

func (t *CacheTable)Name()string{
	return t.name
}
func (t *CacheTable)Count()int  {
	t.Lock()
	defer t.Unlock()
	return len(t.data)
}
func (t *CacheTable)Foreach(f func(key string,item *CacheObject)){
	t.RLock()
	defer t.RUnlock()
	for k,obj:=range t.data{
		f(k,obj)
	}
}
func (t *CacheTable)Keys(){
	t.RLock()
	defer t.RUnlock()
	f := func(k string,item *CacheObject) {
		fmt.Println(k)
	}
	t.Foreach(f)
}
func (t *CacheTable)SetLoadNilCallback(f func(key string,arg ...interface{}) *CacheObject){
	t.Lock()
	defer t.Unlock()
	t.loadNilDataCallback = f
}
func (t *CacheTable)AddAddedCallback(f func(obj *CacheObject)){
	t.Lock()
	defer t.Unlock()
	callback := t.addedCallback
	if callback==nil{
		callback = make([]func(obj *CacheObject),1)
	}
	t.addedCallback = append(callback,f)
}
func (t *CacheTable)AddDeleteCallBack (f func(obj *CacheObject)){
	t.Lock()
	defer t.Unlock()
	callback := t.deleteCallBack
	if callback==nil{
		callback = make([]func(obj *CacheObject),1)
	}
	t.deleteCallBack = append(callback,f)
}
//func (t *CacheTable)SetLogger(logger *log.Logger){
//	t.logger = logger
//}
func (t *CacheTable)remove(key string) *CacheObject{
	if len(key)==0{
		return nil
	}
	t.Lock()
	object,ok := t.data[key]
	if !ok{
		return nil
	}
	delete(t.data,key)
	//fmt.Println("remove table ",t.name,",key ",key)
	t.Unlock()
	callBacks := t.deleteCallBack
	if callBacks!=nil{
		for _,callBack :=range callBacks{
			callBack(object)
		}
	}
	return object
}

//失效检查
func (t *CacheTable)ExpCheck(){
	//t.Lock()

	if t.clearupTimer !=nil{
		t.clearupTimer.Stop()
	}
	//if t.clearupInterval>0{
	//	fmt.Println("exp check after",t.clearupInterval,"for table",t.name)
	//}else {
	//	fmt.Println("exp check after for table",t.name)
	//}
	now :=time.Now()
	for k,v :=range t.data {
		v.RLock()
		exp :=v.expiryTime
		accessTime := v.lastAccessTime
		v.RUnlock()
		if exp==0 {
			continue
		}
		if now.Sub(accessTime) >= exp{
			t.remove(k)
		}
	}
	//t.Unlock()
	t.clearupTimer =time.AfterFunc(t.clearupInterval,func(){
		go t.ExpCheck()
	})


}

func (t *CacheTable)Add(obj *CacheObject){

	fmt.Println("add [",obj.key,",",obj.value,"]")
	t.Lock()
	t.data[obj.key] = obj
	t.Unlock()
	callbacks := t.addedCallback
	if callbacks!=nil{
		for _,callback := range callbacks{
			callback(obj)
		}
	}
}

func (t *CacheTable) Get(key string) *CacheObject{
	t.RLock()
	defer t.RUnlock()
	if len(key)==0 {
		return nil
	}
	return t.data[key]
}

