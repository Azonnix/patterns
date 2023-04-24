package main

type adidas struct {
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			name:  "adidas short",
			color: "red",
		},
	}
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			name: "adidas shoe",
			size: 10,
		},
	}
}
