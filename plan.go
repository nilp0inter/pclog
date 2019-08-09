package main

type Plan struct {
	Name      string
	Exercises []Exercise
}

type Exercise struct {
	Name         string
	Progressions []Progression
}

type Progression struct {
	Name  string
	Goals []Goal
}

type Goal struct {
	Name string
	Sets int
	Reps int
}
