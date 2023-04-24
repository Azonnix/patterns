package main

type puma struct {
}

func (p *puma) makeShort() iShort {
	return &pumaShort{
		short: short{
			name:  "pume short",
			color: "blue",
		},
	}
}

func (p *puma) makeShoe() iShoe {
	return &pumaShoe{
		shoe: shoe{
			name: "puma shoe",
			size: 12,
		},
	}
}
