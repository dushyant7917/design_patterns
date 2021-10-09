package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
	x int
}

var singleInstance *single

func GetSingleInstance(idx int) *single {
	if singleInstance != nil {
		fmt.Println("`singeInstance` is already initialized, idx:", idx)
		return singleInstance
	}

	lock.Lock()
	defer lock.Unlock()
	if singleInstance != nil {
		fmt.Println("`singeInstance` got initialized when locking was in progress, idx:", idx)
		return singleInstance
	}

	singleInstance = &single{}
	fmt.Println("Initialized `singeInstance`, idx:", idx)
	return singleInstance
}

/*
To test above code for multiple threads use below code
for idx := 1; idx <= 50; idx++ {
	go GetSingleInstance(idx)
}
*/
