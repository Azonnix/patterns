package main

import "fmt"

func main() {
	gun, err := gunFactory("ak47")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(gun.getName(), gun.getPower())

	gun, err = gunFactory("maverik")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(gun.getName(), gun.getPower())
}
