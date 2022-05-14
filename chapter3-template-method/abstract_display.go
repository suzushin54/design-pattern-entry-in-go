package main

type AbstractDisplay interface {
	Open()
	Print()
	Close()
}

type abstractDisplay struct {
	AbstractDisplay
}

func (a *abstractDisplay) Display() {
	a.Open()
	for i := 0; i < 5; i++ {
		a.Print()
	}
	a.Close()
}
