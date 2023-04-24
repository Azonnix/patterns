package main

type iSportsFactory interface {
	makeShort() iShort
	makeShoe() iShoe
}

func getSportsFactory(typeFactory string) iSportsFactory {
	if typeFactory == "adidas" {
		return &adidas{}
	}

	if typeFactory == "puma" {
		return &puma{}
	}

	return nil
}
