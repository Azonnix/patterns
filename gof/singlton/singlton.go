package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstence *single

func getSingleInstance() *single {
	if singleInstence == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstence == nil {
			singleInstence = &single{}
			fmt.Println("created")
		} else {
			fmt.Println("already created")
		}
	} else {
		fmt.Println("already created")
	}
	return singleInstence
}
