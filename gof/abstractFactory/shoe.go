package main

type shoe struct {
	name string
	size int
}

func (s *shoe) setName(name string) {
	s.name = name
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getName() string {
	return s.name
}

func (s *shoe) getSize() int {
	return s.size
}
