package sorter

type Person struct {
	FirstName string
	LastName  string
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].FirstName > p[j].FirstName
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p [j], p[i]
}
