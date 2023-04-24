package main

type short struct {
	name  string
	color string
}

func (s *short) setName(name string) {
	s.name = name
}

func (s *short) setColor(color string) {
	s.color = color
}

func (s *short) getName() string {
	return s.name
}

func (s *short) getColor() string {
	return s.color
}
