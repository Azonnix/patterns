package main

import "fmt"

func gunFactory(gun string) (iGun, error) {
	if gun == "ak47" {
		return newAk47(), nil
	}
	if gun == "maverik" {
		return newMaverik(), nil
	}

	return nil, fmt.Errorf("No gun instraction")
}
