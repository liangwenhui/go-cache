package gccode

import (
	"sync"
	"time"
)

var(
	lock sync.RWMutex
	cache = make(map[string]*CacheTable)
)

func Cache(tbname string) *CacheTable{
	lock.RLock()
	table,ok := cache[tbname]
	lock.RUnlock()
	if !ok {
		table = &CacheTable{
			name: tbname,
			data: make(map[string]*CacheObject,16),
			clearupTimer: time.NewTimer(time.Second),
			clearupInterval: time.Second,
		}

		//table.SetLogger()
		cache[tbname]= table
		go table.ExpCheck()
	}

	return table
}

func Tables()[]string{
	lock.RLock()
	defer lock.RUnlock()
	if cache ==nil || len(cache)==0{
		return nil
	}
	var tbns = make([]string,0)
	for k,_:=range cache{
		tbns = append(tbns,k)
	}
	return tbns
}

