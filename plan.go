package main

type Selectable interface {
	Items() []string
	GetOption(int) Selectable
}

type Plan struct {
	Name      string
	Exercises []Exercise
}

func (p Plan) Items() []string {
	exercises := make([]string, len(p.Exercises))
	for i, e := range p.Exercises {
		exercises[i] = e.Name
	}
	return exercises
}

func (p Plan) GetOption(idx int) Selectable {
	return p.Exercises[idx]
}

type Exercise struct {
	Name         string
	Progressions []Progression
}

func (e Exercise) Items() []string {
	progressions := make([]string, len(e.Progressions))
	for i, p := range e.Progressions {
		progressions[i] = p.Name
	}
	return progressions
}

func (e Exercise) GetOption(idx int) Selectable {
	return e.Progressions[idx]
}

type Progression struct {
	Name  string
	Goals []Goal
}

func (p Progression) Items() []string {
	goals := make([]string, len(p.Goals))
	for i, g := range p.Goals {
		goals[i] = g.Name
	}
	return goals
}

func (p Progression) GetOption(idx int) Selectable {
	return p.Goals[idx]
}

type Goal struct {
	Name string
	Sets int
	Reps int
}

func (g Goal) Items() []string {
	return []string{}
}

func (g Goal) GetOption(_ int) Selectable {
	return nil
}
