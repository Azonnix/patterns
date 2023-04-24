package main

import "fmt"

type file struct {
	name string
}

func (f *file) print(data string) {
	fmt.Println(data + f.name)
}

func (f *file) copy() file {
	return file{name: f.name + "_copy"}
}
