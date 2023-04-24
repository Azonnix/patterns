package main

type maverik struct {
	gun
}

func newMaverik() *maverik {
	return &maverik{
		gun: gun{
			Name:  "maverik gun",
			Power: 5,
		},
	}
}
