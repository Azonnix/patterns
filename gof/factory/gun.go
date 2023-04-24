package main

type gun struct {
	Name  string
	Power int
}

func (g *gun) setName(name string) {
	g.Name = name
}

func (g *gun) setPower(power int) {
	g.Power = power
}

func (g *gun) getName() string {
	return g.Name
}

func (g *gun) getPower() int {
	return g.Power
}
