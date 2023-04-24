package main

type director struct {
	builder iBuilder
}

func newDirector(builder iBuilder) *director {
	return &director{
		builder: builder,
	}
}

func (d *director) setBuilder(builder iBuilder) {
	d.builder = builder
}

func (d *director) build() house {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setFloor()
	return d.builder.getHouse()
}
