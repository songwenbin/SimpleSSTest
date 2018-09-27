package loginserver

import (
	"fmt"
	"sync"
)

type AccountManager struct {
	sync.Mutex
	mapper map[string]*AccountInfo //string == publickey
}

type AccountInfo struct {
	ip   string
	port string
	key  string
}

func NewAccountManager() *AccountManager {
	return &AccountManager{
		mapper: map[string]*AccountInfo{},
	}
}

func (am *AccountManager) Add(publickey string) {
	am.Lock()
	am.mapper[publickey] = &AccountInfo{
		ip:   "localhost",
		port: RandomPort(),
		key:  RandomPassword(),
	}
	am.Unlock()
}

func (am *AccountManager) Get(publickey string) *AccountInfo {
	am.Lock()
	value, exists := am.mapper[publickey]
	if exists == false {
		fmt.Println("ok")
		am.Unlock()
		return nil
	} else {
		am.Unlock()
		return value
	}
}

func RandomPort() string {
	return "8888"
}

func RandomPassword() string {
	return "foobar"
}
